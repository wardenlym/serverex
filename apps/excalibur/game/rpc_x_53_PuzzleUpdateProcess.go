package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) PuzzleUpdateProcess(req *msg.Request, arg *msg.PuzzleUpdateProcessRequest) *msg.PuzzleUpdateProcessResponse {

	err := try_do(func() {
		g.use_actor(req.Uid).
			character(arg.CharacterId).
			puzzle_update(arg.Process).
			save()
	})

	if err != msg.Success {
		return &msg.PuzzleUpdateProcessResponse{Err: err}
	}
	return &msg.PuzzleUpdateProcessResponse{}
}
