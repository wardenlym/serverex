package game

import (
	"math/rand"
	"time"

	"github.com/kyokomi/lottery"
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func code_to_string(code msg.Code) string {
	if name, exist := msg.Code_To_String[code]; exist {
		return name
	} else {
		return "UnknowCode"
	}

}

func errcode_to_string(code msg.ErrCode) string {
	if name, exist := msg.ErrCode_To_String[code]; exist {
		return name
	} else {
		return "ErrUnknown"
	}
}

func checkCellEnough(cells []odm.Cell) (bool, int) {
	for i, c := range cells {
		if c.Item == nil || c.Num == 0 {
			return true, i
		}
	}
	return false, -1
}

func checkPutToCells(from []odm.Cell, to_cells []odm.Cell) (new_cells []odm.Cell, unpick []odm.Cell) {

	unpick = make([]odm.Cell, 0)
	new_cells = make([]odm.Cell, len(to_cells))
	copy(new_cells, to_cells)

	k := 0
	for i, v := range new_cells {
		if v.Item == nil || v.Num == 0 {
			if k >= len(from) {
				break
			}
			new_cells[i] = from[k]
			k++
		}
	}

	for i := k; i < len(from); i++ {
		unpick = append(unpick, from[i])
	}
	return new_cells, unpick
}

// func check_materials_enough_in_cells(materials []odm.PlainCell, cells []odm.Cell) (bool, []odm.Cell) {
// 	requiredCnt := map[int]uint{}
// 	haveCnt := map[int]uint{}
// 	new_cells := make([]odm.Cell, len(cells))

// 	for _, material := range materials {
// 		requiredCnt[material.TypeId] = material.Num
// 		haveCnt[material.TypeId] = 0
// 	}

// 	for i, cell := range cells {
// 		if cell.Item == nil || cell.Num == 0 {
// 			new_cells[i] = cell
// 			continue
// 		}
// 		id := cell.Item.TypeId
// 		if _, yes := requiredCnt[id]; yes {
// 			if haveCnt[id] < requiredCnt[id] {
// 				// 尚未计算完消耗材料个数

// 				to_use := requiredCnt[id] - haveCnt[id]
// 				if to_use < cell.Num {
// 					// 当前格子消耗后已经够了
// 					haveCnt[id] = haveCnt[id] + to_use
// 					cell.Num = cell.Num - to_use
// 					new_cells[i] = cell
// 				} else if to_use == cell.Num {
// 					haveCnt[id] = haveCnt[id] + cell.Num
// 					cell.Num = 0
// 					cell.Item = nil
// 					new_cells[i] = cell
// 				} else {
// 					haveCnt[id] = haveCnt[id] + cell.Num
// 					cell.Num = 0
// 					cell.Item = nil
// 					new_cells[i] = cell
// 				}
// 			} else {
// 				new_cells[i] = cell
// 			}
// 		} else {
// 			new_cells[i] = cell
// 		}
// 	}

// 	for id, cnt := range requiredCnt {
// 		if haveCnt[id] < cnt {
// 			log.Warn("###################", id)
// 			return false, nil
// 		}
// 	}

// 	return true, new_cells
// }

type upgrade_prob struct {
	star int
	prob int
}

func (d upgrade_prob) Prob() int {
	return d.prob
}

var lot lottery.Lottery
var up_prob []lottery.Interface = []lottery.Interface{
	upgrade_prob{1, 50},
	upgrade_prob{2, 30},
	upgrade_prob{3, 15},
	upgrade_prob{4, 5},
}

func init() {
	lot = lottery.New(rand.New(rand.NewSource(time.Now().UnixNano())))
}

func upgradeRandom() int {
	return up_prob[lot.Lots(up_prob...)].(upgrade_prob).star
}

func is_rune_type(id int) bool {
	if _, ok := csv.Table_RuneData[id]; ok {
		return true
	}
	return false
}
