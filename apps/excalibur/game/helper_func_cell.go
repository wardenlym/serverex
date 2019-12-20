package game

import (
	"reflect"
	"sort"

	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/log"
	"gitlab.com/megatech/serverex/lib/util"
)

func cell_new() odm.Cell {
	return odm.Cell{}
}

func cell_is_vaild_index(i int, cells []odm.Cell) bool {
	if i >= 0 && i < len(cells) {
		return true
	}
	return false
}

func cell_is_empty(c odm.Cell) bool {
	if c.Item == nil || c.Num == 0 {
		return true
	}
	return false
}

func cell_remove_all(cells []odm.Cell) []odm.Cell {
	for i, _ := range cells {
		cells[i].Item = nil
		cells[i].Num = 0
	}
	return cells
}

func cell_remove_item(to_destory []int, cells []odm.Cell) (bool, []odm.Cell) {
	for _, i := range to_destory {
		if cell_is_vaild_index(i, cells) {
			cells[i] = cell_new()
		} else {
			return false, cells
		}
	}
	return true, cells
}

func cell_is_stackable(typeid int) bool {
	if entry, exist := csv.Table_GoodsItem[typeid]; exist && entry.Sp_number_value > 1 {
		return true
	}
	return false
}

func cell_is_can_carry_back(typeid int) bool {
	if entry, exist := csv.Table_GoodsItem[typeid]; exist && entry.Carry_back == 1 {
		return true
	}
	return false
}

func cell_is_equipment_type(typeid int) bool {
	if entry, exist := csv.Table_GoodsItem[typeid]; exist && (entry.GoodsEnum == 4 || entry.GoodsEnum == 5) {
		return true
	}
	return false
}

func cell_stackable_max_count(typeid int) uint {
	if entry, exist := csv.Table_GoodsItem[typeid]; exist && entry.Sp_number_value > 1 {
		return uint(entry.Sp_number_value)
	}
	return 1
}

// 这个现在我自己也看不懂了，but it works。没有bug，但需要重构
// (可以先做一个查找表处理第一趟)
func cell_move_from_to(from []odm.Cell, to_cells []odm.Cell) (newcells []odm.Cell, unpick []odm.Cell) {

	unpick = make([]odm.Cell, 0)
	newcells = make([]odm.Cell, len(to_cells))
	copy(newcells, to_cells)

	wait_for_second_check := []odm.Cell{}

	// 第一趟会处理所有不堆叠物品，并尽量放置可堆叠物品
	for _, v := range from {
		toput := v
		func() {
			for i, cell := range newcells {
				if (cell.Item == nil || cell.Num == 0) && !cell_is_stackable(toput.Item.TypeId) {
					// 不可堆叠，放在空位
					newcells[i] = toput
					return // 处理下一个待放置物品
				} else {
					if cell.Item != nil && cell.Item.TypeId == toput.Item.TypeId && cell_is_stackable(toput.Item.TypeId) {
						max := cell_stackable_max_count(toput.Item.TypeId)
						if cell.Num < max {
							put_avalible_num := max - cell.Num
							if toput.Num <= put_avalible_num {
								cell.Num += toput.Num
								newcells[i] = cell
								return // 可堆叠物品堆叠完了
							} else {
								cell.Num += put_avalible_num
								toput.Num -= put_avalible_num
								newcells[i] = cell
							}
						}
					}
				}
			}
			// 检查完背包，无可用格子，留待二次确认
			wait_for_second_check = append(wait_for_second_check, toput)
		}()
	}

	// 第二趟检查可堆叠材料
	for _, v := range wait_for_second_check {
		func() {
			for i, cell := range newcells {
				if cell.Item == nil {
					newcells[i] = v
					return
				}
			}
			unpick = append(unpick, v)
		}()
	}

	return newcells, unpick
}

func cell_select_by_index(indices []int, cells []odm.Cell) (selected []odm.Cell) {
	for _, index := range indices {
		selected = append(selected, cells[index])
	}
	return selected
}

func cell_remove_by_index(indices []int, cells []odm.Cell) []odm.Cell {
	for _, index := range indices {
		cells[index] = odm.Cell{}
	}
	return cells
}

func cell_diff(a []odm.Cell, b []odm.Cell) (diff []odm.LocatedCell) {
	for i, v := range a {
		x := b[i]
		if !reflect.DeepEqual(v, x) {
			diff = append(diff, odm.LocatedCell{
				Index: i,
				Cell:  x,
			})
		}
	}

	if len(a) < len(b) {
		for i := len(b) - len(a) - 1; i < len(b); i++ {
			x := b[i]
			diff = append(diff, odm.LocatedCell{
				Index: i,
				Cell:  x,
			})
		}
	}
	return diff
}

func cell_move_all_from_to(from []odm.Cell, to []odm.Cell) (rest []odm.Cell, new_from []odm.Cell, new_to []odm.Cell) {
	indices := make([]int, len(from))
	for i, _ := range indices {
		indices[i] = i
	}
	return cell_move_from_to_2_todotodo(indices, from, to)
}

// 希望用来替代丑陋的cell_move_from_to，尚未实现
func cell_move_from_to_2_todotodo(indices []int, from []odm.Cell, to []odm.Cell) (rest []odm.Cell, new_from []odm.Cell, new_to []odm.Cell) {
	new_from = make([]odm.Cell, len(from))
	copy(new_from, from)
	new_to = make([]odm.Cell, len(to))
	copy(new_to, to)

	return rest, new_from, new_to
}

type sortable_cell []odm.Cell

func (c sortable_cell) Len() int {
	return len(c)
}

func (c sortable_cell) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c sortable_cell) Less(i, j int) bool {
	switch {
	case c[i].Item == nil && c[j].Item == nil:
		return true
	case c[i].Item != nil && c[j].Item == nil:
		return true
	case c[i].Item == nil && c[j].Item != nil:
		return false
	case c[i].Item == nil && c[j].Item != nil:
		return false
	case c[i].Item != nil && c[j].Item != nil && c[i].Item.TypeId == c[j].Item.TypeId:
		return c[i].Num > c[j].Num
	case c[i].Item != nil && c[j].Item != nil && c[i].Item.TypeId != c[j].Item.TypeId:
		return c[i].Item.TypeId < c[j].Item.TypeId
	default:
		// wont be here
		panic(msg.ErrImpossiableLogic)
	}
}

func cell_sort(src []odm.Cell) []odm.Cell {
	dst := make([]odm.Cell, len(src))

	plains := cell_cells_to_plains(src)
	cells := cell_plains_to_cells(plains)

	tosort, _ := cell_move_from_to(cells, dst)

	sort.Sort(sortable_cell(tosort))
	return tosort
}

func cell_plains_to_cells(plain []odm.PlainCell) []odm.Cell {
	ret := []odm.Cell{}
	for _, v := range plain {
		per := cell_stackable_max_count(v.TypeId)
		total := v.Num
		for total > 0 {
			num := func() uint {
				if total >= per {
					return per
				} else {
					return total
				}
			}()
			ret = append(ret, cell_new_with_typeid_n(v.TypeId, num))
			total -= num
		}
	}
	return ret
}

func cell_plains_fold(plain []odm.PlainCell) []odm.PlainCell {
	d := map[int]*odm.PlainCell{}

	for _, v := range plain {
		if _, exist := d[v.TypeId]; exist {
			d[v.TypeId].Num += v.Num
		} else {
			a := v
			d[v.TypeId] = &a
		}
	}

	ret := []odm.PlainCell{}
	for _, v := range d {
		ret = append(ret, *v)
	}
	return ret
}

func cell_cells_to_plains(cells []odm.Cell) []odm.PlainCell {
	ret := []odm.PlainCell{}
	dic := map[int]uint{}

	for _, v := range cells {
		if v.Item != nil {
			dic[v.Item.TypeId] += v.Num
		}
	}

	for id, num := range dic {
		ret = append(ret, odm.PlainCell{
			TypeId: id,
			Num:    num,
		})
	}

	return ret
}

func cell_remove_cannot_take_back(src []odm.Cell) []odm.Cell {
	dst := make([]odm.Cell, len(src))
	copy(dst, src)
	for i, v := range dst {
		if v.Item != nil {
			if !cell_is_can_carry_back(v.Item.TypeId) {
				dst[i].Item = nil
				dst[i].Num = 0
			}
		}
	}
	return dst
}

func cell_remove_random_1(src []odm.Cell) []odm.Cell {
	dst := make([]odm.Cell, len(src))
	copy(dst, src)
	if len(dst) == 0 {
		return dst
	}

	idx := []int{}
	for i, v := range dst {
		if v.Item != nil && v.Num != 0 {
			idx = append(idx, i)
		}
	}

	if len(idx) == 0 {
		return dst
	}

	i := util.RandomIn(0, len(idx))
	dst[idx[i]].Item = nil
	dst[idx[i]].Num = 0
	return dst
}

func cell_remove_materials_if_enough(materials []odm.PlainCell, cells []odm.Cell) (bool, []odm.Cell) {
	requiredCnt := map[int]uint{}
	haveCnt := map[int]uint{}
	new_cells := make([]odm.Cell, len(cells))

	for _, material := range materials {
		requiredCnt[material.TypeId] = material.Num
		haveCnt[material.TypeId] = 0
	}

	for i, cell := range cells {
		if cell.Item == nil || cell.Num == 0 {
			new_cells[i] = cell
			continue
		}
		id := cell.Item.TypeId
		if _, yes := requiredCnt[id]; yes {
			if haveCnt[id] < requiredCnt[id] {
				// 尚未计算完消耗材料个数

				to_use := requiredCnt[id] - haveCnt[id]
				if to_use < cell.Num {
					// 当前格子消耗后已经够了
					haveCnt[id] = haveCnt[id] + to_use
					cell.Num = cell.Num - to_use
					new_cells[i] = cell
				} else if to_use == cell.Num {
					haveCnt[id] = haveCnt[id] + cell.Num
					cell.Num = 0
					cell.Item = nil
					new_cells[i] = cell
				} else {
					haveCnt[id] = haveCnt[id] + cell.Num
					cell.Num = 0
					cell.Item = nil
					new_cells[i] = cell
				}
			} else {
				new_cells[i] = cell
			}
		} else {
			new_cells[i] = cell
		}
	}

	for id, cnt := range requiredCnt {
		if haveCnt[id] < cnt {
			log.Warn(" cell_remove_materials_if_enough ###################", id)
			return false, nil
		}
	}

	return true, new_cells
}
