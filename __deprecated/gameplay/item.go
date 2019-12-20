package gameplay

import (
	"fmt"

	"github.com/json-iterator/go"
	"gitlab.com/megatech/serverex/apps/excalibur/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/odm"
	"gitlab.com/megatech/serverex/lib/log"
)

func (g *game) destroyItem(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.DestroyItemRequest
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

	if o.CellIndex < 0 || o.CellIndex >= info.Bag.CellsCapacity {
		res.Data = &msg.DestroyItemResponse{IsSuccess: false, ErrMsg: "无法摧毁"}
		return res
	}

	cell := odm.Cell{Index: o.CellIndex}
	err = g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".bag.cells."+fmt.Sprint(o.CellIndex), cell)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	res.Data = &msg.DestroyItemResponse{IsSuccess: true}
	return res
}

func (g *game) consumItem(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.ConsumItemRequest
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

	bag := info.Bag
	if o.CellIndex < 0 || o.CellIndex >= bag.CellsCapacity || bag.Cells[o.CellIndex].Num < o.Num {
		res.Data = &msg.ConsumItemResponse{IsSuccess: false, ErrMsg: "使用物品错误"}
		log.Error(o.CellIndex, bag.Cells[o.CellIndex].Num, o.Num)
		return res
	}

	newCell := odm.Cell{}
	if bag.Cells[o.CellIndex].Num > o.Num {
		newCell = bag.Cells[o.CellIndex]
		newCell.Num = newCell.Num - o.Num
	}

	err = g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".bag.cells."+fmt.Sprint(o.CellIndex), newCell)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	res.Data = &msg.ConsumItemResponse{IsSuccess: true, CellNewStatus: newCell}
	return res
}

func (g *game) equip(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.EquipRequest
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

	// 如果背包里没有这个item
	if o.FromCellIndex < 0 || o.FromCellIndex >= info.Bag.CellsCapacity ||
		info.Bag.Cells[o.FromCellIndex].Item == nil || info.Bag.Cells[o.FromCellIndex].Num == 0 ||
		// 如果装备栏不是空的
		o.ToCellIndex < 0 || o.ToCellIndex >= len(info.Bag.Equipments) {
		res.Data = &msg.EquipResponse{IsSuccess: false, ErrMsg: "装备物品错误"}
		log.Error(o, info.Bag)
		return res
	}

	bag := info.Bag

	// swap
	origin := info.Bag.Equipments[o.ToCellIndex]

	toEquip := bag.Cells[o.FromCellIndex].Item
	num := bag.Cells[o.FromCellIndex].Num

	bag.Cells[o.FromCellIndex].Item = origin.Item
	bag.Cells[o.FromCellIndex].Num = origin.Num

	bag.Equipments[o.ToCellIndex].Item = toEquip
	bag.Equipments[o.ToCellIndex].Num = num

	err = g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".bag", bag)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	res.Data = &msg.EquipResponse{IsSuccess: true}
	return res
}

func (g *game) unequip(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.UnequipRequest
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

	// 如果装备栏里没有这个item
	if o.FromCellIndex < 0 || o.FromCellIndex >= len(info.Bag.Equipments) ||
		info.Bag.Equipments[o.FromCellIndex].Item == nil || info.Bag.Equipments[o.FromCellIndex].Num == 0 ||
		// 如果背包不是空的
		o.ToCellIndex < 0 || o.ToCellIndex >= info.Bag.CellsCapacity ||
		info.Bag.Cells[o.ToCellIndex].Num != 0 || info.Bag.Cells[o.ToCellIndex].Item != nil {
		res.Data = &msg.UnequipResponse{IsSuccess: false, ErrMsg: "卸下装备错误"}
		log.Error(o)
		return res
	}

	bag := info.Bag
	toUnequip := bag.Equipments[o.FromCellIndex].Item
	num := bag.Equipments[o.FromCellIndex].Num

	bag.Equipments[o.FromCellIndex].Item = nil
	bag.Equipments[o.FromCellIndex].Num = 0

	bag.Cells[o.ToCellIndex].Item = toUnequip
	bag.Cells[o.ToCellIndex].Num = num

	err = g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".bag", bag)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	res.Data = &msg.UnequipResponse{IsSuccess: true}
	return res
}
