package game

import (
	"time"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) ReduceExploreTime(req *msg.Request, arg *msg.ReduceExploreTimeRequest) *msg.ReduceExploreTimeResponse {

	var explore_info odm.ExploreInfo
	var diamond uint

	err := try_do(func() {
		actor := g.use_actor(req.Uid)

		actor.reduce_explore_time(
			arg.CharacterId,
			arg.ExploreStageId,
			arg.ReduceTime)

		explore_info = actor.user_info().ExploreInfo
		diamond = actor.user_info().Diamond

		actor.save()
	})

	if err != msg.Success {
		return &msg.ReduceExploreTimeResponse{Err: err}
	}
	return &msg.ReduceExploreTimeResponse{
		ServerTime:  time.Now().Unix(),
		ExploreInfo: explore_info,
		Diamond:     diamond,
	}
}
