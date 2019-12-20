package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) ExitStage(req *msg.Request, arg *msg.ExitStageRequest) *msg.ExitStageResponse {

	var best_stage_id int
	var chara_state string

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)
		actor.update_dungeon_stat(arg.AttributeSnapshot, arg.StageStat)
		actor.exit_stage(arg.StageClear)
		best_stage_id = actor.chara_info().RankingInfo.BestStageId

		if arg.IsChapterEnding {
			if rec, is_new_record := actor.try_make_new_dungeon_ranking_record(arg.StageId); is_new_record {
				g.db.Ranking().InsertDungeonRanking(rec)
			}
		}
		actor.save()
	})

	if err != msg.Success {
		return &msg.ExitStageResponse{Err: err}
	}
	return &msg.ExitStageResponse{
		State:       chara_state,
		BestStageId: best_stage_id,
	}
}
