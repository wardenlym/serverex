package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/util"
)

func (g *GameService) RuneEquip(req *msg.Request, arg *msg.RuneEquipRequest) *msg.RuneEquipResponse {
	user := g.get_user(req.Uid)
	chara := user.Characters[arg.CharacterId]

	// 如果背包里没有这个符文

	if arg.FromCellIndex < 0 || arg.FromCellIndex >= len(chara.Bag.Cells) ||
		chara.Bag.Cells[arg.FromCellIndex].Item == nil || chara.Bag.Cells[arg.FromCellIndex].Num < 1 ||
		!is_rune_type(chara.Bag.Cells[arg.FromCellIndex].Item.TypeId) {
		return &msg.RuneEquipResponse{Err: msg.ErrInvalidCellIndex}
	}

	// 如果符文插槽错误
	if arg.ToCellIndex < 0 || arg.ToCellIndex >= len(chara.RuneTree.Cells) ||
		chara.RuneTree.Cells[arg.ToCellIndex].Item != nil {
		return &msg.RuneEquipResponse{Err: msg.ErrInvalidRuneTypeId}
	}

	// 符文插槽类型错误
	bagCell, runeCell := chara.Bag.Cells[arg.FromCellIndex], chara.RuneTree.Cells[arg.ToCellIndex]
	if runeCell.Enable != true ||
		runeCell.SlotColor != csv.Table_RuneData[bagCell.Item.TypeId].Slot {
		return &msg.RuneEquipResponse{Err: msg.ErrInvalidRuneTypeId}
	}

	old_cells := make([]odm.Cell, len(chara.Bag.Cells))
	copy(old_cells, chara.Bag.Cells)

	item := *chara.Bag.Cells[arg.FromCellIndex].Item
	// swap
	if chara.Bag.Cells[arg.FromCellIndex].Num <= 1 {
		chara.Bag.Cells[arg.FromCellIndex].Item = nil
		chara.Bag.Cells[arg.FromCellIndex].Num = 0
	} else {
		chara.Bag.Cells[arg.FromCellIndex].Num--
	}

	chara.RuneTree.Cells[arg.ToCellIndex].Item = &odm.Item{
		Guid:            util.NewSnowflakeID(),
		TypeId:          item.TypeId,
		StarLevel:       item.StarLevel,
		StarUpgradeInfo: []int{},
	}

	user.Characters[arg.CharacterId] = chara

	bag_diff := cell_diff(old_cells, chara.Bag.Cells)

	g.save_user(user)

	return &msg.RuneEquipResponse{
		RuneTree: chara.RuneTree,
		BagDiff:  bag_diff,
	}
}
