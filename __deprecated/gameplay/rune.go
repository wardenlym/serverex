package gameplay

import (
	jsoniter "github.com/json-iterator/go"
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/odm"
	"gitlab.com/megatech/serverex/lib/util"
)

func is_rune_type(id int) bool {
	if _, ok := csv.Table_RuneData[id]; ok {
		return true
	}
	return false
}

func (g *game) runeEquip(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.RuneEquipRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}

	info, err := g.db.GetCharacter(req.Uid, o.CharacterId)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	runeCells := info.RuneTree.Cells
	bagCells := info.Bag.Cells

	// 如果背包里没有这个符文
	if o.FromCellIndex < 0 || o.FromCellIndex >= info.Bag.CellsCapacity ||
		bagCells[o.FromCellIndex].Item == nil || bagCells[o.FromCellIndex].Num < 1 ||
		!is_rune_type(bagCells[o.FromCellIndex].Item.TypeId) {
		res.Data = &msg.RuneEquipResponse{IsSuccess: false, ErrMsg: "背包内符文错误"}
		return res
	}

	// 如果符文插槽错误
	if o.ToCellIndex < 0 || o.ToCellIndex >= len(runeCells) ||
		runeCells[o.ToCellIndex].Item != nil {
		res.Data = &msg.RuneEquipResponse{IsSuccess: false, ErrMsg: "符文插槽错误"}
		return res
	}

	// 符文插槽类型错误
	bagCell := bagCells[o.FromCellIndex]
	runeCell := runeCells[o.ToCellIndex]
	if runeCell.Enable != true ||
		runeCell.SlotColor != csv.Table_RuneData[bagCell.Item.TypeId].Slot {
		res.Data = &msg.RuneEquipResponse{IsSuccess: false, ErrMsg: "符文插槽类型错误"}
		return res
	}

	bag := info.Bag
	tree := info.RuneTree
	item := *bag.Cells[o.FromCellIndex].Item
	// swap
	if bag.Cells[o.FromCellIndex].Num <= 1 {
		bag.Cells[o.FromCellIndex].Item = nil
		bag.Cells[o.FromCellIndex].Num = 0
	} else {
		bag.Cells[o.FromCellIndex].Num--
	}
	tree.Cells[o.ToCellIndex].Item = &odm.Item{
		Guid:            util.NewSnowflakeID(),
		TypeId:          item.TypeId,
		StarLevel:       item.StarLevel,
		StarUpgradeInfo: []int{},
	}

	g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".bag", bag)
	g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".runeTree", tree)

	res.Data = &msg.RuneEquipResponse{
		IsSuccess: true,
		BagStatus: bag,
		RuneTree:  tree,
	}
	return res
}

func (g *game) runeUnequip(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.RuneUnequipRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}

	info, err := g.db.GetCharacter(req.Uid, o.CharacterId)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	tree := info.RuneTree

	// 如果符文插槽内不存在
	if o.FromCellIndex < 0 || o.FromCellIndex >= len(tree.Cells) ||
		tree.Cells[o.FromCellIndex].Item == nil {
		res.Data = &msg.RuneUnequipResponse{IsSuccess: false, ErrMsg: "符文插槽错误"}
		return res
	}

	// 找存储空间放置
	ok, errmsg, cells, index := findPosition(tree.Cells[o.FromCellIndex], info.Bag.Cells)
	tree.Cells[o.FromCellIndex].Item = nil
	if !ok {
		res.Data = &msg.RuneUnequipResponse{IsSuccess: false, ErrMsg: errmsg}
		return res
	}

	bag := info.Bag
	bag.Cells = cells
	g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".bag", bag)
	g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".runeTree", tree)

	res.Data = &msg.RuneUnequipResponse{
		IsSuccess: true,
		BagStatus: bag,
		RuneTree:  tree,
		CellIndex: index,
	}
	return res
}

func findPosition(r odm.RuneCell, cells []odm.Cell) (bool, string, []odm.Cell, int) {
	for i, v := range cells {
		if v.Item != nil && v.Num != 0 {
			if v.Item.TypeId == r.Item.TypeId && v.Num < 30 {
				cells[i].Num++
				return true, "", cells, i
			}
		} else {
			cells[i].Item = r.Item
			cells[i].Num = 1
			return true, "", cells, i
		}
	}
	return false, "背包空间不足", nil, -1
}
