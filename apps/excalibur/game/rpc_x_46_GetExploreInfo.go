package game

import (
	"time"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) GetExploreInfo(req *msg.Request, arg *msg.GetExploreInfoRequest) *msg.GetExploreInfoResponse {

	var explore_info odm.ExploreInfo

	err := try_do(func() {
		actor := g.use_actor(req.Uid)

		actor.if_explore_info_need_refresh()

		explore_info = actor.user_info().ExploreInfo
		actor.save()
	})

	if err != msg.Success {
		return &msg.GetExploreInfoResponse{Err: err}
	}
	return &msg.GetExploreInfoResponse{
		ServerTime:  time.Now().Unix(),
		ExploreInfo: explore_info,
	}
}
