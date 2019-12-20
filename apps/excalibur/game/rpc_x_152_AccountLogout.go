package game

import (
	"fmt"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) AccountLogout(req *msg.Request, arg *msg.AccountLogoutRequest) *msg.AccountLogoutResponse {
	return &msg.AccountLogoutResponse{Message: fmt.Sprint(req.Uid) + " logout ok."}
}
