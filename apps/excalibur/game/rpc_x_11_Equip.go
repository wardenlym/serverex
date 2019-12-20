package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) Equip(req *msg.Request, arg *msg.EquipRequest) *msg.EquipResponse {

	err := try_do(func() {
		actor := g.
			use_actor(req.Uid).
			character(arg.CharacterId).
			equip(arg.FromCellIndex, arg.ToCellIndex)
		actor.save()
	})

	if err != msg.Success {
		return &msg.EquipResponse{Err: err}
	}
	return &msg.EquipResponse{}
}
