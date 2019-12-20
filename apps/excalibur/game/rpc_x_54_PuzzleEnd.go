package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) PuzzleEnd(req *msg.Request, arg *msg.PuzzleEndRequest) *msg.PuzzleEndResponse {
	err := try_do(func() {
		g.use_actor(req.Uid).
			character(arg.CharacterId).
			puzzle_end(arg.Selected).
			save()
	})

	if err != msg.Success {
		return &msg.PuzzleEndResponse{Err: err}
	}
	return &msg.PuzzleEndResponse{}
}
