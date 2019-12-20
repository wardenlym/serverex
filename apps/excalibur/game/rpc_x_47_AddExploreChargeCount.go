package game

import (
	"time"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) AddExploreChargeCount(req *msg.Request, arg *msg.AddExploreChargeCountRequest) *msg.AddExploreChargeCountResponse {

	var explore_info odm.ExploreInfo
	var diamond uint

	err := try_do(func() {
		actor := g.use_actor(req.Uid)

		actor.add_explore_charge_count(arg.Count,
			uint(arg.Count)*consts.ExploreChangeDiamand)

		explore_info = actor.user_info().ExploreInfo
		diamond = actor.user_info().Diamond

		actor.save()
	})

	if err != msg.Success {
		return &msg.AddExploreChargeCountResponse{Err: err}
	}
	return &msg.AddExploreChargeCountResponse{
		ServerTime:  time.Now().Unix(),
		ExploreInfo: explore_info,
		Diamond:     diamond,
	}
}
