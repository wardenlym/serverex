package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/log"
)

func (a *Actor) equip(from, to int) *Actor {

	if from < 0 || from >= len(a.chara.Bag.Cells) ||
		to < 0 || to >= len(a.chara.Bag.Equipments) {
		// 如果index非法
		log.Error(1)
		raise_error(msg.ErrInvalidCellIndex)
	}

	if a.chara.Bag.Cells[from].Item == nil || a.chara.Bag.Cells[from].Num == 0 {
		// 如果背包格子没有装备
		log.Error(2)
		raise_error(msg.ErrInvalidCellIndex)
	}

	if !cell_is_equipment_type(a.chara.Bag.Cells[from].Item.TypeId) {
		raise_error(msg.ErrInvalidItemId)
	}

	//允许替换装备，去掉空装备栏的检查

	// if a.chara.Bag.Equipments[to].Item != nil || a.chara.Bag.Equipments[to].Num != 0 {
	// 	// 如果装备栏不是空的
	// 	log.Error(3)
	// 	raise_error(msg.ErrInvalidCellIndex)
	// }

	// 交换
	a.chara.Bag.Cells[from], a.chara.Bag.Equipments[to] =
		a.chara.Bag.Equipments[to], a.chara.Bag.Cells[from]

	return a
}

func (a *Actor) unequip(from, to int) *Actor {

	if from < 0 || from >= len(a.chara.Bag.Equipments) ||
		to < 0 || to >= len(a.chara.Bag.Cells) {
		// 如果index非法
		log.Error(1)
		raise_error(msg.ErrInvalidCellIndex)
	}

	if a.chara.Bag.Equipments[from].Item == nil || a.chara.Bag.Equipments[from].Num == 0 {
		// 如果装备栏没有装备
		log.Error(2)
		raise_error(msg.ErrInvalidCellIndex)
	}

	if a.chara.Bag.Cells[to].Item != nil || a.chara.Bag.Cells[to].Num != 0 {
		// 如果背包格子不是空的
		log.Error(a.chara.Bag.Cells[to].Item, a.chara.Bag.Cells[to].Num)
		raise_error(msg.ErrInvalidCellIndex)
	}

	// 交换
	a.chara.Bag.Equipments[from], a.chara.Bag.Cells[to] =
		a.chara.Bag.Cells[to], a.chara.Bag.Equipments[from]

	return a
}

func (a *Actor) expand_bag() *Actor {
	return a
}

func (a *Actor) expand_stash() *Actor {
	return a
}

func (a *Actor) is_bag_contains() bool {
	return false
}

func (a *Actor) is_stash_and_bag_contains() bool {
	return false
}

func (a *Actor) remove_from_bag() *Actor {
	return a
}

func (a *Actor) put_to_bag(award []odm.Cell) *Actor {
	newcells, unpicked := cell_move_from_to(award, a.chara.Bag.Cells)

	if len(unpicked) > 0 {
		raise_error(msg.ErrBagCellNotEnough)
	}

	a.chara.Bag.Cells = newcells

	return a
}

func (a *Actor) pickup_to_bag(award []odm.Cell) []odm.Cell {
	newcells, unpicked := cell_move_from_to(award, a.chara.Bag.Cells)
	a.chara.Bag.Cells = newcells
	a.chara.DungeonStatus.UnpickedLoot = unpicked
	a.chara.DungeonStatus.LootTotal = append(a.chara.DungeonStatus.LootTotal, award...)
	return unpicked
}

func (a *Actor) destory_in_bag(indices []int) *Actor {
	ok, cells := cell_remove_item(indices, a.chara.Bag.Cells)
	if !ok {
		raise_error(msg.ErrInvalidCellIndex)
	}
	a.chara.Bag.Cells = cells
	return a
}

func (a *Actor) destory_in_stash(indices []int) *Actor {
	return a
}

func (a *Actor) move_from_stash_to_bag(indices []int) *Actor {
	return a
}

func (a *Actor) move_from_bag_to_stash(indices []int) *Actor {
	return a
}

func (a *Actor) sort_bag() {
	new_bag := cell_sort(a.chara.Bag.Cells)
	a.chara.Bag.Cells = new_bag
}

func (a *Actor) sort_stash() {
	new_stash := cell_sort(a.chara.Bag.Cells)
	a.user.Stash.Cells = new_stash
}

func (a *Actor) diff_stash() []odm.LocatedCell {
	return cell_diff(a.old_user.Stash.Cells, a.user.Stash.Cells)
}

func (a *Actor) diff_bag() []odm.LocatedCell {
	return cell_diff(a.old_user.Characters[a.chara_id].Bag.Cells, a.chara.Bag.Cells)
}

func (a *Actor) diff_equipments() []odm.LocatedCell {
	return cell_diff(a.old_user.Characters[a.chara_id].Bag.Equipments, a.chara.Bag.Equipments)
}

func (a *Actor) consum_item_n(index int, count uint) odm.Cell {
	if index < 0 || index >= len(a.chara.Bag.Cells) {
		raise_error(msg.ErrInvalidCellIndex)
	}

	if a.chara.Bag.Cells[index].Num < count {
		raise_error(msg.ErrAvalibleCountNotEnough)
	}

	newCell := odm.Cell{}
	if a.chara.Bag.Cells[index].Num > count {
		newCell = a.chara.Bag.Cells[index]
		newCell.Num = newCell.Num - count
	}

	if a.chara.Bag.Cells[index].Num == count {
		newCell = a.chara.Bag.Cells[index]
		newCell.Num = 0
		newCell.Item = nil
	}
	return newCell
}

func (a *Actor) cost_requires_in_bag(requires []odm.PlainCell) *Actor {
	ok, cells := cell_remove_materials_if_enough(requires, a.chara.Bag.Cells)

	if !ok {
		raise_error(msg.ErrAvalibleCountNotEnough)
	}
	a.chara.Bag.Cells = cells
	return a
}
