package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func get_workshop_info(typeid int) (uint, uint, []odm.PlainCell, []odm.PlainCell) {
	info, exist := csv.Table_Workshop[typeid]
	if !exist {
		raise_error(msg.ErrInvalidGoodsId)
	}

	require_plain := []odm.PlainCell{}
	decompoose_plain := []odm.PlainCell{}

	for i, _ := range info.Reslove_materials {
		require_plain = append(require_plain, odm.PlainCell{
			TypeId: info.Need_materials[i],
			Num:    uint(info.Need_materials_num[i]),
		})

		decompoose_plain = append(decompoose_plain, odm.PlainCell{
			TypeId: info.Reslove_materials[i],
			Num:    uint(info.Reslove_materials_num[i]),
		})
	}

	return uint(info.Cost_diamond), uint(info.Cost_gold), require_plain, decompoose_plain
}

func (a *Actor) stat_if_chengzhuang(typeid int, num uint) {
	if v, exist := csv.Table_GoodsItem[typeid]; exist && v.Quality == 5 {
		a.user.StatInfo.HuoDeChengZhuangShu += num
	}
}

func (g *GameService) Craft(req *msg.Request, arg *msg.CraftRequest) *msg.CraftResponse {

	var crafed_item odm.Cell
	var diamond uint
	var gold uint
	var stash_gold uint
	bag_diff := []odm.LocatedCell{}

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)

		diamond_decr, gold_decr, requires, _ := get_workshop_info(arg.ToCraftTypeId)

		actor.cost_diamond(diamond_decr)
		actor.cost_gold(gold_decr)
		actor.cost_requires_in_bag(requires)

		crafed_item = cell_new_with_typeid_1(arg.ToCraftTypeId)

		actor.put_to_bag([]odm.Cell{crafed_item})

		diamond = actor.user.Diamond
		gold = actor.chara.Gold
		stash_gold = actor.user.Stash.Gold
		bag_diff = actor.diff_bag()

		actor.stat_if_chengzhuang(crafed_item.Item.TypeId, crafed_item.Num)

		actor.save()
	})

	if err != msg.Success {
		return &msg.CraftResponse{Err: err}
	}

	return &msg.CraftResponse{
		CraftedItems: crafed_item,
		Gold:         gold,
		StashGold:    stash_gold,
		Diamond:      diamond,
		BagDiff:      bag_diff,
	}
}
