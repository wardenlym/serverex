package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) Unequip(req *msg.Request, arg *msg.UnequipRequest) *msg.UnequipResponse {

	err := try_do(func() {
		actor := g.
			use_actor(req.Uid).
			character(arg.CharacterId).
			unequip(arg.FromCellIndex, arg.ToCellIndex)
		actor.save()
	})

	if err != msg.Success {
		return &msg.UnequipResponse{Err: err}
	}
	return &msg.UnequipResponse{}
}
