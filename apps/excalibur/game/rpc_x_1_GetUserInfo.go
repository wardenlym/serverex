package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) GetUserInfo(req *msg.Request, arg *msg.GetUserInfoRequest) *msg.GetUserInfoResponse {

	var user_info *odm.User

	err := try_do(func() {
		user_info = g.use_actor(req.Uid).user_info()
	})

	if err != msg.Success {
		return &msg.GetUserInfoResponse{Err: err}
	}
	return &msg.GetUserInfoResponse{User: user_info}
}
