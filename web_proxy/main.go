package main

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"
	"time"
)

type Client struct {
	time  int64
	count int
}

type handle struct {
	Lhost        string            `json:"local_host"`
	Rhost        string            `json:"remote_host"`
	Length       int               `json:"max_length"`
	Frequncy     int               `json:"frequncy"`
	Urlmap       map[string]string `json:"url_map"`
	HeartbeatUrl string            `json:"heartbeat_url"`
	Heartbeat    string            `json:"heartbeat"`

	black_list map[string]bool
	blmutex    *sync.RWMutex
	count_map  map[string]*Client
	cmmutex    *sync.Mutex
}

var h handle = handle{
	Lhost:      "",
	Rhost:      "",
	Length:     1024,
	Frequncy:   180,
	Urlmap:     map[string]string{},
	black_list: map[string]bool{},
	blmutex:    new(sync.RWMutex),
	count_map:  map[string]*Client{},
	cmmutex:    new(sync.Mutex),
}

func readCfg(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic("fail to read config file: " + path)
		return
	}

	fi, _ := file.Stat()

	buff := make([]byte, fi.Size())
	_, err = file.Read(buff)
	buff = []byte(os.ExpandEnv(string(buff)))

	err = json.Unmarshal(buff, &h)
	if err != nil {
		log.Print(err)
		panic("failed to unmarshal file")
		return
	}
	log.Print(h)
}

func (this *handle) CheckIP(ip string) bool {
	this.blmutex.RLock()
	_, err := h.black_list[ip]
	this.blmutex.RUnlock()
	if err {
		return false
	}

	this.cmmutex.Lock()
	defer this.cmmutex.Unlock()
	now := time.Now().Unix()
	c, err1 := h.count_map[ip]
	if !err1 {
		h.count_map[ip] = &Client{time: now, count: 1}
	} else {
		if now/60 == c.time/60 {
			c.count++
			if c.count > h.Frequncy {
				this.blmutex.Lock()
				h.black_list[ip] = true
				this.blmutex.Unlock()
				log.Printf("ip:%s add to black list", ip)
				return false
			}
		} else {
			c.time = now
			c.count = 1
		}
	}
	return true
}

func checkRequest(query string) bool {
	// 检查crc32
	data := []byte(query)

	sum := string(data[1:9])
	ieee := crc32.NewIEEE()
	io.WriteString(ieee, string(data[10:]))
	s := ieee.Sum32()
	ssum := fmt.Sprintf("%08x", s)
	fmt.Print(query, "  ", sum, "  ", string(data[8:]), "  ", ssum, "  ")
	if sum != ssum {
		fmt.Println("http的crc32校验失败")
		return false
	}

	return true
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle request")
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("ip切割非法")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !h.CheckIP(ip) {
		fmt.Println("ip检查非法")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//   _, err1 := h.Urlmap[r.URL.Path]
	//   if !err1 {
	//       fmt.Println("url检查非法")
	//       http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	//       return
	//   }

	remote, err2 := url.Parse(h.Rhost)
	if err2 != nil {
		fmt.Println("客户端地址获取非法")
		h.black_list[ip] = true
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fmt.Println("代理开始")
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func startServer() {
	//    err := http.ListenAndServe(h.Lhost, &h)
	//    if err != nil {
	//        fmt.Printf("ListenAndServe: ", err)
	//    }

	s := &http.Server{
		Addr:           h.Lhost,
		Handler:        &h,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1024,
	}
	fmt.Println(s.ListenAndServe())
}

func main() {
	if len(os.Args) != 2 {
		panic("please identify a config file")
		return
	}

	readCfg(os.Args[1])

	go heartbeat()
	startServer()

}
