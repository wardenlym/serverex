package game

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/kyokomi/lottery"
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/log"
	"gitlab.com/megatech/serverex/lib/util"
)

func cell_new_with_typeid_1(typeid int) odm.Cell {
	return odm.Cell{Num: 1, Item: &odm.Item{
		Guid:            util.NewSnowflakeID(),
		TypeId:          typeid,
		StarUpgradeInfo: []int{},
	}}
}

func cell_new_with_typeid_n(typeid int, num uint) odm.Cell {
	return odm.Cell{Num: num, Item: &odm.Item{
		Guid:            util.NewSnowflakeID(),
		TypeId:          typeid,
		StarUpgradeInfo: []int{},
	}}
}

// func gen_fake_loot_item() []odm.Cell {
// 	i := util.RandomIn(0, len(csv.Table_FakelootData))
// 	loot1 := csv.Table_FakelootData[i]
// 	award := []odm.Cell{}

// 	for i, typeid := range loot1.Loot {
// 		award = append(award, cell_new_with_typeid_n(typeid, uint(loot1.Drop_number[i])))
// 	}
// 	return award
// }

var loot_dicer lottery.Lottery

func init() {
	loot_dicer = lottery.New(rand.New(rand.NewSource(time.Now().UnixNano())))
}

const DummyTypeId = 0

type LootDice struct {
	TypeId    int
	ProbValue int
	NumRange  [2]int
}

func (d LootDice) Prob() int {
	return d.ProbValue
}

type LootResult struct {
	TypeId int
	Num    int
}

type LootGropIndex struct {
	SubId1    []int
	LuckyThan []int
	SubId2    []int
}

var Table_LootGroupIndex map[string]LootGropIndex
var Table_LootGroupData map[int][]lottery.Interface

func lookup_monster_type(typeid int) (int, error) {
	if v, exist := csv.Table_MonsterData[typeid]; exist {
		return v.Monster_type, nil
	}
	return -1, errors.New("monster type id not exist")
}

func loot_index_key(monster_type, stage_id int) string {
	return fmt.Sprint(monster_type) + "_" + fmt.Sprint(stage_id)
}

func select_subgroup_id(key string, lucky int) []int {
	ids := []int{}
	index1, exist := Table_LootGroupIndex[key]
	if !exist {
		return ids
	}

	for i := 0; i < len(index1.SubId1); i++ {
		if index1.LuckyThan[i] == 0 {
			ids = append(ids, index1.SubId1[i])
		} else {
			if lucky > index1.LuckyThan[i] {
				ids = append(ids, index1.SubId2[i])
			} else {
				ids = append(ids, index1.SubId1[i])
			}
		}
	}

	return ids
}

func lot_from_subgroup(id int) LootResult {
	group, exist := Table_LootGroupData[id]
	if !exist {
		return LootResult{}
	}

	i := loot_dicer.Lots(group...)
	a := group[i].(LootDice)

	return LootResult{
		TypeId: a.TypeId,
		Num:    util.RandomIn(a.NumRange[0], a.NumRange[1]),
	}
}

func loot2(stageid, monster_type, lucky int) ([]LootResult, string, []int) {
	thing := []LootResult{}

	group_key := loot_index_key(monster_type, stageid)

	subgroup_ids := select_subgroup_id(group_key, lucky)

	for _, id := range subgroup_ids {
		item := lot_from_subgroup(id)
		if item.TypeId != DummyTypeId {
			thing = append(thing, item)
		}
	}

	return thing, group_key, subgroup_ids
}

func loot1(stageid, typeid, lucky int) ([]LootResult, string, []int) {
	thing := []LootResult{}

	monster_type, err := lookup_monster_type(typeid)
	if err != nil {
		log.Error(err)
		return thing, "", nil
	}

	group_key := loot_index_key(monster_type, stageid)

	subgroup_ids := select_subgroup_id(group_key, lucky)

	for _, id := range subgroup_ids {
		item := lot_from_subgroup(id)
		if item.TypeId != DummyTypeId {
			thing = append(thing, item)
		}
	}

	return thing, group_key, subgroup_ids
}

func gen_loot_item(stage_id int, npc_type_id int, lucky int) []odm.Cell {

	items := []odm.Cell{}
	thing, _, _ := loot1(stage_id, npc_type_id, lucky)

	for _, v := range thing {
		if v.TypeId != DummyTypeId {
			items = append(items, cell_new_with_typeid_n(v.TypeId, uint(v.Num)))
		}
	}

	return items
}

func make_loot_table_index() map[string]LootGropIndex {
	ret := map[string]LootGropIndex{}

	for _, v := range csv.Table_LootoneData {

		for stageid := v.Stageid[0]; stageid <= v.Stageid[1]; stageid++ {
			key := loot_index_key(v.Monster_type, stageid)
			if _, exist := ret[key]; !exist {
				ret[key] = LootGropIndex{
					SubId1:    []int{},
					LuckyThan: []int{},
					SubId2:    []int{},
				}
			}
			a := LootGropIndex{
				SubId1:    append(ret[key].SubId1, v.Subgroupid),
				LuckyThan: append(ret[key].LuckyThan, v.Lucky_value),
				SubId2:    append(ret[key].SubId2, v.Advanced_group),
			}
			ret[key] = a
		}
	}

	return ret
}

func make_loot_table_data() map[int][]lottery.Interface {

	ret := map[int][]lottery.Interface{}
	for _, v := range csv.Table_LoottwoData {

		if _, exist := ret[v.Subgroupid]; !exist {
			ret[v.Subgroupid] = []lottery.Interface{}
		}

		list := ret[v.Subgroupid]

		a := LootDice{
			TypeId:    v.Id,
			ProbValue: v.Weight,
			NumRange:  [2]int{v.Number_min, v.Number_max + 1},
		}
		list = append(list, a)
		ret[v.Subgroupid] = list
	}

	for k, v := range ret {
		all := 100000
		for _, i := range v {
			a := i.(LootDice)
			all = all - a.ProbValue
		}
		if all < 0 {
			panic("有数据不ok,概率数和>10000 " + fmt.Sprint(k) + " " + fmt.Sprint(v))
		}
		if all == 0 {

		}
		if all > 0 {
			dummy_prob := all
			v = append(v, LootDice{
				TypeId:    DummyTypeId,
				ProbValue: dummy_prob,
				NumRange:  [2]int{0, 1},
			})
			ret[k] = v
		}
	}

	return ret
}

func init() {
	Table_LootGroupIndex = make_loot_table_index()
	Table_LootGroupData = make_loot_table_data()
}
