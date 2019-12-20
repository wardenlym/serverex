package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (a *Actor) upgrade_item(index int) odm.Cell {

	if index < 0 || index >= len(a.chara.Bag.Cells) ||
		a.chara.Bag.Cells[index].Item == nil || a.chara.Bag.Cells[index].Num == 0 {
		raise_error(msg.ErrInvalidCellIndex)
	}

	// [1,4]
	upgrade := upgradeRandom()
	//

	old := a.chara.Bag.Cells[index]
	new := old
	new.Item.StarLevel = new.Item.StarLevel + 1
	new.Item.StarUpgradeInfo = append(new.Item.StarUpgradeInfo, upgrade)
	a.chara.Bag.Cells[index] = new

	return new
}

func (g *GameService) Enhance(req *msg.Request, arg *msg.EnhanceRequest) *msg.EnhanceResponse {

	var new odm.Cell
	var gold uint
	var stash_gold uint
	var diamond uint
	var bag_diff []odm.LocatedCell

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)
		actor.cost_gold(arg.CostGold)
		actor.cost_diamond(arg.CostDiamond)

		actor.cost_requires_in_bag(arg.Requires)

		new = actor.upgrade_item(arg.ToEnhanceIndex)
		gold = actor.chara.Gold
		stash_gold = actor.user.Stash.Gold
		diamond = actor.user.Diamond
		bag_diff = actor.diff_bag()

		actor.save()
	})

	if err != msg.Success {
		return &msg.EnhanceResponse{Err: err}
	}

	return &msg.EnhanceResponse{
		EnhancedItem: new,
		Gold:         gold,
		StashGold:    stash_gold,
		Diamond:      diamond,
		BagDiff:      bag_diff,
	}
}

func (g *GameService) Enhance2(req *msg.Request, arg *msg.EnhanceRequest) *msg.EnhanceResponse {
	user := g.get_user(req.Uid)
	chara := user.Characters[arg.CharacterId]

	// 暂时没有检测消耗金币的正确数字
	if user.Diamond < arg.CostDiamond {
		return &msg.EnhanceResponse{Err: msg.ErrDiamondNotEnough}
	}
	// if user.Gold < arg.CostGold {
	// 	return &msg.EnhanceResponse{Err: msg.ErrGoldNotEnough}
	// }

	// gold := user.Gold - arg.CostGold
	diamond := user.Diamond - arg.CostDiamond

	old_cells := make([]odm.Cell, len(chara.Bag.Cells))
	copy(old_cells, chara.Bag.Cells)

	ok, cells := cell_remove_materials_if_enough(arg.Requires, chara.Bag.Cells)

	if !ok {
		return &msg.EnhanceResponse{Err: msg.ErrAvalibleCountNotEnough}
	}

	ok, index := func(i int, cells []odm.Cell) (bool, int) {
		if i < 0 || i >= len(cells) {
			return false, -1
		}
		if cells[i].Item == nil || cells[i].Num == 0 {
			return false, -1
		}
		return true, i
	}(arg.ToEnhanceIndex, cells)

	if !ok {
		return &msg.EnhanceResponse{Err: msg.ErrInvalidCellIndex}
	}
	//

	// [1,4]
	upgrade := upgradeRandom()
	//

	old := cells[index]
	new := old
	new.Item.StarLevel = new.Item.StarLevel + 1
	new.Item.StarUpgradeInfo = append(new.Item.StarUpgradeInfo, upgrade)
	cells[index] = new

	chara.Bag.Cells = cells
	user.Characters[arg.CharacterId] = chara
	// user.Gold = gold
	user.Diamond = diamond

	bag_diff := cell_diff(old_cells, chara.Bag.Cells)

	g.save_user(user)

	return &msg.EnhanceResponse{
		EnhancedItem: new,
		// Gold:         gold,
		Diamond: diamond,
		BagDiff: bag_diff,
	}
}
