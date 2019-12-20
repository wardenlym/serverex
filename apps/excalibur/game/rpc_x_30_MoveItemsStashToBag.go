package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) MoveItemsStashToBag(req *msg.Request, arg *msg.MoveItemsStashToBagRequest) *msg.MoveItemsStashToBagResponse {
	user := g.get_user(req.Uid)
	chara := user.Characters[arg.CharacterId]

	to_move := cell_select_by_index(arg.Indices, user.Stash.Cells)

	newbag, unpick := cell_move_from_to(to_move, chara.Bag.Cells)
	if len(unpick) != 0 {
		return &msg.MoveItemsStashToBagResponse{Err: msg.ErrBagCellNotEnough}
	}

	newstash := cell_remove_by_index(arg.Indices, user.Stash.Cells)

	stash_diff := cell_diff(user.Stash.Cells, newstash)
	bag_diff := cell_diff(chara.Bag.Cells, newbag)

	user.Stash.Cells = newstash
	chara.Bag.Cells = newbag

	user.Characters[arg.CharacterId] = chara
	g.save_user(user)
	return &msg.MoveItemsStashToBagResponse{
		StashDiff: stash_diff,
		BagDiff:   bag_diff,
	}
}
