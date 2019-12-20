package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) GM_ActivateUser(req *msg.Request, arg *msg.GM_ActivateUserRequest) *msg.GM_ActivateUserResponse {
	g.db.CreateNewUserIfNotExist(req.Uid)
	return &msg.GM_ActivateUserResponse{}
}
