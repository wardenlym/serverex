package game

import (
	"fmt"
	"testing"

	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func Test_diff_cell(t *testing.T) {

	a := cell_diff(nil, nil)
	fmt.Println(a)
}

func Test_diff_cell_2(t *testing.T) {

	fmt.Println(cell_move_from_to_2_todotodo(nil, nil, nil))
}

func Test_diff_cell_index_error(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			if fmt.Sprint(err) == "runtime error: index out of range" {
				println(1)
			} else {
				println(2)
			}
		}
	}()

	a := []odm.Cell{}
	fmt.Println(a[1].Num)
}

func Test_npc_list(t *testing.T) {
	info := csv.Table_StageData[1001]

	mapid, npclist := make_npc_list(info)

	fmt.Println(mapid)
	fmt.Println(npclist)
}

func Test_calc_explore_award(t *testing.T) {
	a := calc_explore_award(1, 10000, 1, 60*60)
	fmt.Println(a)
	award_items := cell_plains_to_cells(a)
	fmt.Println(award_items)
	newcells, unpicked := cell_move_from_to(award_items, make([]odm.Cell, 10))
	fmt.Println(newcells, unpicked)
}
