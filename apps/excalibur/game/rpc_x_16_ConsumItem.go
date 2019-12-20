package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) ConsumItem(req *msg.Request, arg *msg.ConsumItemRequest) *msg.ConsumItemResponse {

	var cell odm.Cell

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)
		cell = actor.consum_item_n(arg.CellIndex, arg.Num)
		actor.save()
	})

	if err != msg.Success {
		return &msg.ConsumItemResponse{Err: err}
	}
	return &msg.ConsumItemResponse{CellNewStatus: cell}
}
