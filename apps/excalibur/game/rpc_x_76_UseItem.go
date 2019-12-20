package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/util"
)

func (a *Actor) consum_items(typeid int, n uint) {

	for i, v := range a.chara.Bag.Cells {
		if v.Item != nil && v.Num != 0 && v.Item.TypeId == typeid {
			if v.Num > n {
				a.chara.Bag.Cells[i].Num = v.Num - n
				return
			} else if v.Num == n {
				a.chara.Bag.Cells[i] = cell_new()
				return
			} else {
				n -= a.chara.Bag.Cells[i].Num
				a.chara.Bag.Cells[i] = cell_new()
			}
		}
	}

	if n > 0 {
		raise_error(msg.ErrAvalibleCountNotEnough)
	}

}

func gen_plain_award_by_typeid_n(typeid int, n uint) []odm.PlainCell {

	award := []odm.PlainCell{}
	info, exist := csv.Table_GiftPackage[typeid]
	if exist {
		for i, id := range info.ItemId {
			award = append(award, odm.PlainCell{
				TypeId: id,
				Num:    uint(util.RandomIn(info.NumMin[i], info.NumMax[i]+1)),
			})
		}
	}

	return award
}

func (a *Actor) diamond_gold_bagdiff() (uint, uint, uint, []odm.LocatedCell) {
	diamond := a.user.Diamond
	gold := a.chara.Gold
	stash_gold := a.user.Stash.Gold
	bag_diff := a.diff_bag()
	return diamond, gold, stash_gold, bag_diff
}

func (g *GameService) UseItem(req *msg.Request, arg *msg.UseItemRequest) *msg.UseItemResponse {

	var diamond uint
	var gold uint
	var stash_gold uint
	bag_diff := []odm.LocatedCell{}

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)

		actor.consum_items(arg.TypeId, arg.Num)

		award := cell_plains_to_cells(
			gen_plain_award_by_typeid_n(arg.TypeId, arg.Num))

		actor.put_to_bag(award)

		diamond, gold, stash_gold, bag_diff = actor.diamond_gold_bagdiff()

		actor.save()
	})

	if err != msg.Success {
		return &msg.UseItemResponse{Err: err}
	}

	return &msg.UseItemResponse{
		Diamond:   diamond,
		Gold:      gold,
		StashGold: stash_gold,
		BagDiff:   bag_diff,
	}
}
