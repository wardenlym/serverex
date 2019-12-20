package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) GetUserStatInfo(req *msg.Request, arg *msg.GetUserStatInfoRequest) *msg.GetUserStatInfoResponse {

	var stat_info odm.StatInfo

	err := try_do(func() {
		actor := g.use_actor(req.Uid)
		stat_info = actor.user.StatInfo
	})

	if err != msg.Success {
		return &msg.GetUserStatInfoResponse{Err: err}
	}
	return &msg.GetUserStatInfoResponse{
		StatInfo: stat_info,
	}
}
