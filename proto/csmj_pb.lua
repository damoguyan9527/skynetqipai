return {["buf"]="\
�\9\
\
csmj.proto\18\4csmj\";\
\
RoomOption\18\11\
\3key\24\1 \1(\9\18\14\
\6nValue\24\2 \1(\5\18\16\
\8strValue\24\3 \1(\9\"-\
\8RoomInfo\18!\
\7options\24\1 \3(\0112\16.csmj.RoomOption\"�\1\
\12UserGameData\18\14\
\6userid\24\1 \2(\4\18\16\
\8nickname\24\2 \2(\9\18\11\
\3sex\24\3 \2(\5\18\18\
\
headimgurl\24\4 \2(\9\18\13\
\5score\24\5 \2(\5\18\
\
\2ip\24\6 \2(\9\18\12\
\4seat\24\7 \2(\5\18\13\
\5ready\24\8 \2(\8\18\16\
\8latitude\24\13 \2(\9\18\17\
\9longitude\24\14 \2(\9\18\12\
\4adds\24\15 \2(\9\18\14\
\6online\24\16 \2(\8\";\
\19GAME_PlayerEnterAck\18$\
\8userData\24\1 \2(\0112\18.csmj.UserGameData\"$\
\19GAME_PlayerLeaveAck\18\13\
\5nSeat\24\1 \2(\5\"c\
\17GAME_EnterGameAck\18\11\
\3err\24\1 \2(\5\18#\
\7players\24\2 \3(\0112\18.csmj.UserGameData\18\28\
\4room\24\3 \2(\0112\14.csmj.RoomInfo\"/\
\
VoteResult\18\13\
\5nSeat\24\1 \2(\5\18\18\
\
nVoteState\24\2 \2(\5\"�\1\
\17GAME_GameSceneAck\18\11\
\3err\24\1 \2(\5\18\28\
\4room\24\2 \2(\0112\14.csmj.RoomInfo\18\14\
\6status\24\3 \2(\9\18#\
\7players\24\4 \3(\0112\18.csmj.UserGameData\18\17\
\9bank_seat\24\5 \2(\5\18\20\
\12nDissoveSeat\24\7 \2(\5\18\20\
\12nDissoveTime\24\8 \2(\5\18\30\
\4vote\24\9 \3(\0112\16.csmj.VoteResult\"\31\
\13GAME_ReadyReq\18\14\
\6bAgree\24\1 \2(\8\"1\
\13GAME_ReadyAck\18\16\
\8wChairID\24\1 \2(\5\18\14\
\6bAgree\24\2 \2(\8\"\15\
\13GAME_BeginReq\"c\
\18GAME_GameEndPlayer\18\12\
\4seat\24\1 \2(\5\18\13\
\5score\24\2 \2(\5\18\19\
\11total_score\24\3 \2(\3\18\13\
\5cards\24\4 \3(\5\18\12\
\4type\24\5 \2(\5\"b\
\15GAME_GameEndAck\18'\
\5infos\24\1 \3(\0112\24.csmj.GAME_GameEndPlayer\18\18\
\
bank_score\24\2 \2(\5\18\18\
\
bank_count\24\3 \2(\5\"\25\
\23GAME_GameTotalEndPlayer\"D\
\20GAME_GameTotalEndAck\18,\
\5infos\24\1 \3(\0112\29.csmj.GAME_GameTotalEndPlayer",}