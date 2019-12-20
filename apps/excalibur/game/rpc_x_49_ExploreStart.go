package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) ExploreStart(req *msg.Request, arg *msg.ExploreStartRequest) *msg.ExploreStartResponse {

	var explore_info odm.ExploreInfo
	err := try_do(func() {
		actor := g.use_actor(req.Uid).
			explore_start(arg.CharacterId, arg.ExploreStageId)
		actor.save()
		explore_info = actor.user_info().ExploreInfo
	})

	if err != msg.Success {
		return &msg.ExploreStartResponse{Err: err}
	}
	return &msg.ExploreStartResponse{
		ExploreInfo: explore_info,
	}
}
