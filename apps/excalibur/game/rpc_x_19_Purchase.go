package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) Purchase(req *msg.Request, arg *msg.PurchaseRequest) *msg.PurchaseResponse {

	var diamond uint
	var gold uint
	var stash_gold uint
	bag_diff := []odm.LocatedCell{}

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)

		onsales := actor.user_info().OnSaleInfos
		index := func() int {
			for i, goods := range onsales {
				if arg.Guid == goods.Guid && arg.GoodsId == goods.GoodsId && arg.TypeId == goods.TypeId {
					if goods.NumAvailable >= arg.Num {
						return i
					} else {
						raise_error(msg.ErrAvalibleCountNotEnough)
					}
				}
			}
			raise_error(msg.ErrInvalidGoodsId)
			return -1
		}()

		sq_num := onsales[index].NumSQ
		onsales[index].NumAvailable = onsales[index].NumAvailable - arg.Num

		actor.cost_diamond(onsales[index].CostDiamond)
		actor.cost_gold(onsales[index].CostGold)

		thing := cell_new_with_typeid_n(onsales[index].TypeId, sq_num)
		actor.put_to_bag([]odm.Cell{thing})

		gold = actor.chara.Gold
		stash_gold = actor.user.Stash.Gold
		diamond = actor.user.Diamond
		bag_diff = actor.diff_bag()

		actor.stat_if_chengzhuang(thing.Item.TypeId, thing.Num)
		actor.save()
	})

	if err != msg.Success {
		return &msg.PurchaseResponse{Err: err}
	}

	return &msg.PurchaseResponse{
		Gold:      gold,
		StashGold: stash_gold,
		Diamond:   diamond,
		BagDiff:   bag_diff,
	}
}
