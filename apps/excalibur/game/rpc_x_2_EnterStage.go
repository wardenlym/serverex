package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) EnterStage(req *msg.Request, arg *msg.EnterStageRequest) *msg.EnterStageResponse {

	var stage_info odm.StageInfo
	var chara_state string

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)
		stage_info = g.world().gen_stage_info(arg.StageId)
		actor.enter_stage(stage_info)
		chara_state = actor.chara_state()
		actor.save()
	})

	if err != msg.Success {
		return &msg.EnterStageResponse{Err: err}
	}
	return &msg.EnterStageResponse{
		State:     chara_state,
		StageInfo: stage_info,
	}
}
