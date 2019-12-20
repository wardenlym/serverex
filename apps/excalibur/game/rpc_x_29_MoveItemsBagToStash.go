package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) MoveItemsBagToStash(req *msg.Request, arg *msg.MoveItemsBagToStashRequest) *msg.MoveItemsBagToStashResponse {
	user := g.get_user(req.Uid)
	chara := user.Characters[arg.CharacterId]

	to_move := cell_select_by_index(arg.Indices, chara.Bag.Cells)

	newstash, unpick := cell_move_from_to(to_move, user.Stash.Cells)
	if len(unpick) != 0 {
		return &msg.MoveItemsBagToStashResponse{Err: msg.ErrBagCellNotEnough}
	}

	newbag := cell_remove_by_index(arg.Indices, chara.Bag.Cells)

	stash_diff := cell_diff(user.Stash.Cells, newstash)
	bag_diff := cell_diff(chara.Bag.Cells, newbag)

	user.Stash.Cells = newstash
	chara.Bag.Cells = newbag

	user.Characters[arg.CharacterId] = chara
	g.save_user(user)
	return &msg.MoveItemsBagToStashResponse{
		StashDiff: stash_diff,
		BagDiff:   bag_diff,
	}
}
