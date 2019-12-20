package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) SortItemsInStash(req *msg.Request, arg *msg.SortItemsInStashRequest) *msg.SortItemsInStashResponse {

	var stash_diff []odm.LocatedCell

	err := try_do(func() {
		actor := g.use_actor(req.Uid)
		actor.sort_stash()
		stash_diff = actor.diff_stash()
		actor.save()
	})

	if err != msg.Success {
		return &msg.SortItemsInStashResponse{Err: err}
	}
	return &msg.SortItemsInStashResponse{
		StashDiff: stash_diff,
	}

}
