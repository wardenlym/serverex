package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) GM_SetMyMoney(req *msg.Request, arg *msg.GM_SetMyMoneyRequest) *msg.GM_SetMyMoneyResponse {
	user := g.get_user(req.Uid)

	user.Stash.Gold = arg.Gold
	user.Diamond = arg.Diamond

	g.save_user(user)
	return &msg.GM_SetMyMoneyResponse{}
}
