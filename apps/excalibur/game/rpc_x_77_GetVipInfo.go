package game

import (
	"time"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) GetVipInfo(req *msg.Request, arg *msg.GetVipInfoRequest) *msg.GetVipInfoResponse {

	var vip_info odm.VipInfo

	err := try_do(func() {
		actor := g.use_actor(req.Uid)
		vip_info = actor.user.VipInfo
		actor.save()
	})

	if err != msg.Success {
		return &msg.GetVipInfoResponse{Err: err}
	}

	return &msg.GetVipInfoResponse{
		ServerTime: time.Now().Unix(),
		VipInfo:    vip_info,
	}
}
