package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) SortItemsInBag(req *msg.Request, arg *msg.SortItemsInBagRequest) *msg.SortItemsInBagResponse {

	var bag_diff []odm.LocatedCell

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)
		actor.sort_bag()
		bag_diff = actor.diff_bag()
		actor.save()
	})

	if err != msg.Success {
		return &msg.SortItemsInBagResponse{Err: err}
	}
	return &msg.SortItemsInBagResponse{
		BagDiff: bag_diff,
	}
}
