package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) DestroyMultipleItem(req *msg.Request, arg *msg.DestroyMultipleItemRequest) *msg.DestroyMultipleItemResponse {

	var bag_diff []odm.LocatedCell

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)
		actor.destory_in_bag(arg.CellIndices)
		bag_diff = actor.diff_bag()
		actor.save()
	})

	if err != msg.Success {
		return &msg.DestroyMultipleItemResponse{Err: err}
	}
	return &msg.DestroyMultipleItemResponse{
		BagDiff: bag_diff,
	}
}
