package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) GM_SetVipInfo(req *msg.Request, arg *msg.GM_SetVipInfoRequest) *msg.GM_SetVipInfoResponse {
	user := g.get_user(req.Uid)
	user.VipInfo = arg.VipInfo
	g.save_user(user)

	return &msg.GM_SetVipInfoResponse{}
}
