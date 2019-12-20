package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) ExitDungeon(req *msg.Request, arg *msg.ExitDungeonRequest) *msg.ExitDungeonResponse {

	var chara_state string
	var npc_list []odm.Npc
	var item_list []odm.Cell
	var item_plain_list []odm.PlainCell
	var equipments_diff []odm.LocatedCell
	var bag_diff []odm.LocatedCell
	var gold_total uint
	var gold_ext_total uint
	var dungeon_score odm.DungeonScore
	var stash_gold uint
	var gold uint

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)

		dungeon_score = g.world().calc_score_dungeon(actor.chara_info().DungeonStatus)

		npc_list, item_list, gold_total, gold_ext_total = actor.current_dungeon_stat()

		item_plain_list = cell_cells_to_plains(item_list)

		actor.exit_dungeon(arg.IsSafeExit)

		equipments_diff = actor.diff_equipments()
		bag_diff = actor.diff_bag()
		stash_gold = actor.user.Stash.Gold

		chara_state = actor.chara_state()
		actor.save()
	})

	if err != msg.Success {
		return &msg.ExitDungeonResponse{Err: err}
	}

	return &msg.ExitDungeonResponse{
		State:          chara_state,
		Gold:           gold,
		StashGold:      stash_gold,
		DungeonGold:    gold_total,
		DungeonExtGold: gold_ext_total,
		NpcList:        npc_list,
		ItemList:       item_list,
		ItemPlainList:  item_plain_list,
		EquipmentsDiff: equipments_diff,
		BagDiff:        bag_diff,
		DungeonScore:   dungeon_score,
	}
}
