package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) PuzzleStart(req *msg.Request, arg *msg.PuzzleStartRequest) *msg.PuzzleStartResponse {

	err := try_do(func() {
		puzzle_info := odm.PuzzleInfo{
			RandomSeed: arg.RandomSeed,
			Total:      arg.Total,
		}
		g.use_actor(req.Uid).
			character(arg.CharacterId).
			puzzle_start(puzzle_info).
			save()
	})

	if err != msg.Success {
		return &msg.PuzzleStartResponse{Err: err}
	}
	return &msg.PuzzleStartResponse{}
}
