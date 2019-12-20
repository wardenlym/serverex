package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) RuneUnequip(req *msg.Request, arg *msg.RuneUnequipRequest) *msg.RuneUnequipResponse {

	user := g.get_user(req.Uid)
	chara := user.Characters[arg.CharacterId]

	// 如果符文插槽内不存在
	if arg.FromCellIndex < 0 || arg.FromCellIndex >= len(chara.RuneTree.Cells) ||
		chara.RuneTree.Cells[arg.FromCellIndex].Item == nil {
		return &msg.RuneUnequipResponse{Err: msg.ErrInvalidRuneTypeId}
	}

	old_cells := make([]odm.Cell, len(chara.Bag.Cells))
	copy(old_cells, chara.Bag.Cells)

	// 找存储空间放置
	ok, cells, index := findPosition(chara.RuneTree.Cells[arg.FromCellIndex], chara.Bag.Cells)
	chara.RuneTree.Cells[arg.FromCellIndex].Item = nil
	if !ok {
		return &msg.RuneUnequipResponse{Err: msg.ErrBagCellNotEnough}
	}

	chara.Bag.Cells = cells

	user.Characters[arg.CharacterId] = chara
	bag_diff := cell_diff(old_cells, chara.Bag.Cells)
	g.save_user(user)

	return &msg.RuneUnequipResponse{
		RuneTree:  chara.RuneTree,
		CellIndex: index,
		BagDiff:   bag_diff,
	}
}

func findPosition(r odm.RuneCell, cells []odm.Cell) (bool, []odm.Cell, int) {
	for i, v := range cells {
		if v.Item != nil && v.Num != 0 {
			if v.Item.TypeId == r.Item.TypeId && v.Num < 30 {
				cells[i].Num++
				return true, cells, i
			}
		} else {
			cells[i].Item = r.Item
			cells[i].Num = 1
			return true, cells, i
		}
	}
	return false, nil, -1
}
