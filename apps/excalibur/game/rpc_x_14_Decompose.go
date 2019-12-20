package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) Decompose(req *msg.Request, arg *msg.DecomposeRequest) *msg.DecomposeResponse {

	var decompose_id int
	var decompose_num uint
	var diamond uint
	var gold uint
	var stash_gold uint
	bag_diff := []odm.LocatedCell{}

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)

		typeid := actor.chara.Bag.Cells[arg.DecomposedCellIndex].Item.TypeId
		_, _, _, decomposed := get_workshop_info(typeid)
		decompose_id = decomposed[0].TypeId
		decompose_num = decomposed[0].Num
		// 放置背包
		actor.chara.Bag.Cells[arg.DecomposedCellIndex].Item = nil
		actor.chara.Bag.Cells[arg.DecomposedCellIndex].Num = 0

		award := cell_plains_to_cells(decomposed)
		actor.put_to_bag(award)

		diamond = actor.user.Diamond
		gold = actor.chara.Gold
		stash_gold = actor.user.Stash.Gold
		bag_diff = actor.diff_bag()

		actor.save()
	})

	if err != msg.Success {
		return &msg.DecomposeResponse{Err: err}
	}

	return &msg.DecomposeResponse{
		TypeId:    decompose_id,
		Num:       decompose_num,
		Gold:      gold,
		StashGold: stash_gold,
		Diamond:   diamond,
		BagDiff:   bag_diff,
	}
}
