package game

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_loot_table_index(t *testing.T) {
	Table_LootGroupIndex = make_loot_table_index()
	fmt.Println(Table_LootGroupIndex)
}
func Test_loot_table_data(t *testing.T) {
	Table_LootGroupData = make_loot_table_data()
	fmt.Println(Table_LootGroupData)
}

func test_loot1_report(stageid, typeid, lucky int) string {
	thing, group_key, sub_group_ids := loot1(stageid, typeid, lucky)
	report := fmt.Sprintf("%d,%d,%d,%s,%v,%v\n", stageid, typeid, lucky, group_key, sub_group_ids, thing)
	return report
}

func test_loot2_report(stageid, monster_type, lucky int) string {
	thing, group_key, sub_group_ids := loot2(stageid, monster_type, lucky)
	report := fmt.Sprintf("%d,%d,%d,%s,%v,%v\n", stageid, monster_type, lucky, group_key, sub_group_ids, thing)
	return report
}

func Test_loot1(t *testing.T) {
	thing, group_key, sub_group_ids := loot1(1120, 10000218, 0)
	fmt.Println(thing, group_key, sub_group_ids)
}

func Test_loot2_report(t *testing.T) {
	a := fmt.Sprintf("%s,%s,%s,%s,%s,%s\n", "stageid", "typeid", "lucky", "group_key", "sub_group_ids", "thing")

	check := 10

	for i := 0; i < check; i++ {
		a += test_loot2_report(1050, 1, 0)
	}
	for i := 0; i < check; i++ {
		a += test_loot2_report(1060, 1, 0)
	}
	for i := 0; i < check; i++ {
		a += test_loot2_report(1120, 1, 0)
	}

	for i := 0; i < check; i++ {
		a += test_loot2_report(1015, 5, 0)
	}
	for i := 0; i < check; i++ {
		a += test_loot2_report(1100, 5, 0)
	}
	for i := 0; i < check; i++ {
		a += test_loot2_report(1120, 5, 0)
	}

	for i := 0; i < check; i++ {
		a += test_loot2_report(1050, 7, 0)
	}
	for i := 0; i < check; i++ {
		a += test_loot2_report(1120, 7, 0)
	}

	ioutil.WriteFile("helper_func_loot_test_掉落测试1_物品.csv", []byte(a), 0644)
}

func Test_loot_prob(t *testing.T) {

	check := 10000

	test_prob := func(group_id int) {
		m := map[int]int{}
		fmt.Printf("掉落组id %d \n", group_id)
		for i := 0; i < check; i++ {
			v := lot_from_subgroup(group_id)
			m[v.TypeId]++
		}

		for k, v := range m {
			name := k
			count := v
			result := float64(count) / float64(check) * 100
			fmt.Printf("ok %3.5f%%(%7d) : typeid %d \n", result, count, name)
		}
	}

	for id, _ := range Table_LootGroupData {
		test_prob(id)
	}

}

var suit1 = [][]int{
	{1001, 1, 0, 10},
	{1001, 5, 0, 10},
	{1001, 6, 0, 10},
	{1001, 7, 0, 10},
	{1001, 8, 0, 10},
	{2010, 1, 20, 10},
	{2010, 5, 20, 10},
	{2010, 6, 20, 10},
	{2010, 7, 20, 10},
	{2010, 8, 20, 10},
	{3024, 1, 0, 10},
	{3024, 5, 0, 10},
	{3024, 6, 0, 10},
	{3024, 7, 0, 10},
	{3024, 8, 0, 10},
	{3024, 1, 100, 10},
	{3024, 5, 100, 10},
	{3024, 6, 100, 10},
	{3024, 7, 100, 10},
	{3024, 8, 100, 10}}

func Test_loot4_report(t *testing.T) {
	a := fmt.Sprintf("%s,%s,%s,%s,%s,%s\n", "stageid", "typeid", "lucky", "group_key", "sub_group_ids", "thing")

	check := 10
	v := []int{2010, 1, 20, 10}
	stageid, monster_type, lucky := v[0], v[1], v[2]
	for i := 0; i < check; i++ {
		a += test_loot2_report(stageid, monster_type, lucky)
	}

	ioutil.WriteFile("helper_func_loot_test_掉落测试1_物品.csv", []byte(a), 0644)
}

func Test_loot3_report(t *testing.T) {
	a := fmt.Sprintf("%s,%s,%s,%s,%s,%s\n", "stageid", "typeid", "lucky", "group_key", "sub_group_ids", "thing")

	check := 10

	for _, v := range suit1 {
		stageid, monster_type, lucky := v[0], v[1], v[2]
		for i := 0; i < check; i++ {
			a += test_loot2_report(stageid, monster_type, lucky)
		}
	}

	ioutil.WriteFile("helper_func_loot_test_掉落测试1_物品.csv", []byte(a), 0644)
}

func Test_loot5_report(t *testing.T) {
	a := fmt.Sprintf("%s,%s,%s,%s,%s,%s\n", "stageid", "typeid", "lucky", "group_key", "sub_group_ids", "thing")

	check := 10
	v := []int{1001, 7, 0, 10}
	stageid, _, lucky := v[0], v[1], v[2]
	for i := 0; i < check; i++ {
		a += test_loot1_report(stageid, 10000617, lucky)
	}

	ioutil.WriteFile("helper_func_loot_test_掉落测试1_物品.csv", []byte(a), 0644)
}

var suit2 = [][]int{
	{1001, 1, 0, 100},
	{2001, 1, 0, 100},
	{3001, 1, 0, 100},
}

func Test_loot6_report(t *testing.T) {
	a := fmt.Sprintf("%s,%s,%s,%s,%s,%s\n", "stageid", "typeid", "lucky", "group_key", "sub_group_ids", "thing")

	check := 10

	for _, v := range suit2 {
		stageid, monster_type, lucky := v[0], v[1], v[2]
		for i := 0; i < check; i++ {
			a += test_loot2_report(stageid, monster_type, lucky)
		}
	}

	ioutil.WriteFile("helper_func_loot_test_掉落测试1_物品.csv", []byte(a), 0644)
}
