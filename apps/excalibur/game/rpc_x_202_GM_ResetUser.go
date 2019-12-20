package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) GM_ResetUser(req *msg.Request, arg *msg.GM_ResetUserRequest) *msg.GM_ResetUserResponse {
	g.db.DeleteUser(req.Uid)
	g.db.CreateNewUserIfNotExist(req.Uid)
	return &msg.GM_ResetUserResponse{}
}
