package game

import (
	"fmt"
	"testing"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/util"
)

func Test_a(t *testing.T) {
	for i := 0; i < 100; i++ {
		println(util.RandomIn(0, 1))
	}
}

func Test_put(t *testing.T) {

	award := []odm.Cell{
		cell_new_with_typeid_n(10015, 10000),
		cell_new_with_typeid_n(20000, 1),
		cell_new_with_typeid_n(11000, 60),
		cell_new_with_typeid_n(11001, 60),
	}
	cells := make([]odm.Cell, 15)

	newcells, unpick := cell_move_from_to(award, cells)

	fmt.Println(unpick)
	fmt.Println(newcells)

	newcells2, unpick := cell_move_from_to(award, newcells)

	fmt.Println(unpick)
	fmt.Println(newcells2)

}

func Test_put_a(t *testing.T) {

	award := []odm.Cell{
		cell_new_with_typeid_n(13016, 2),
		cell_new_with_typeid_n(13017, 1),
	}
	cells := make([]odm.Cell, 15)

	newcells, unpick := cell_move_from_to(award, cells)

	fmt.Println(unpick)
	fmt.Println(newcells)
}

func Test_lot(t *testing.T) {
	countMap := map[int]int{}

	check := 500000
	for i := 0; i < check; i++ {
		v := lot.Lots(up_prob...)
		countMap[v]++
	}

	for i, count := range countMap {
		result := float64(count) / float64(check) * 100
		prob := float64(up_prob[i].Prob())

		name := fmt.Sprint(up_prob[i].(upgrade_prob).star)

		// 0.1 check
		if (prob-1) <= result && result < (prob+1) {
			fmt.Printf("ok %3.5f%%(%7d) : star%s\n", result, count, name)
		} else {
			t.Errorf("error %3.5f%%(%7d) : star%s\n", result, count, name)
		}
	}
}

func Test_market(t *testing.T) {

}

func Test_market_2(t *testing.T) {
	a := makeOnSaleInfoFromTable(1001)
	fmt.Println(a)
}

func Test_cell_plains_to_cells(t *testing.T) {
	plains := []odm.PlainCell{
		odm.PlainCell{10015, 10000},
		odm.PlainCell{20000, 3},
		odm.PlainCell{11000, 100},
		odm.PlainCell{11001, 101},
	}
	cells := cell_plains_to_cells(plains)
	fmt.Println(cells)
}

func Test_cell_plains_fold(t *testing.T) {
	plains := []odm.PlainCell{
		odm.PlainCell{10015, 1},
		odm.PlainCell{10015, 1},
		odm.PlainCell{10015, 1},
		odm.PlainCell{10015, 1},
	}
	cells := cell_plains_fold(plains)
	fmt.Println(cells)
}
