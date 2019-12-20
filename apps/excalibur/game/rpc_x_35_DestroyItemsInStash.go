package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) DestroyItemsInStash(req *msg.Request, arg *msg.DestroyItemsInStashRequest) *msg.DestroyItemsInStashResponse {
	user := g.get_user(req.Uid)

	ok, new_stash := cell_remove_item(arg.Indices, user.Stash.Cells)

	if !ok {
		return &msg.DestroyItemsInStashResponse{Err: msg.ErrInvalidCellIndex}
	}

	stash_diff := cell_diff(user.Stash.Cells, new_stash)

	user.Stash.Cells = new_stash
	g.save_user(user)
	return &msg.DestroyItemsInStashResponse{
		StashDiff: stash_diff,
	}
}
