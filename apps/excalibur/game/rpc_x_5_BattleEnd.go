package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) BattleEnd(req *msg.Request, arg *msg.BattleEndRequest) *msg.BattleEndResponse {

	var chara_state string
	var kill_value_increased uint
	var kill_value uint
	var gold_increased uint
	var gold_ext_increased uint
	var gold uint
	var unpickup_items []odm.Cell
	var all_award_items []odm.Cell
	var bag_cell []odm.Cell
	var bag_diff []odm.LocatedCell

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)

		npc := actor.current_battle_with()

		kill_value_increased = g.world().calc_kill_value(npc.NpcTypeId)

		stage_id, npc_type_id, lucky :=
			actor.chara_info().DungeonStatus.CurrentStageInfo.StageId,
			npc.NpcTypeId,
			actor.chara_info().BattleAttribute.LuckyValue

		all_award_items, gold_increased =
			g.world().gen_loot_award(stage_id, npc_type_id, int(lucky))

		if actor.is_vip_godbless() {
			gold_ext_increased = gold_increased / 2
		}

		unpickup_items = actor.pickup_to_bag(all_award_items)
		bag_cell = actor.chara_info().Bag.Cells // TODO: 删除
		bag_diff = actor.diff_bag()

		actor.battle_end(
			arg.BattleResult,
			kill_value_increased,
			gold_increased,
			gold_ext_increased)

		chara_state = actor.chara_state()
		gold = actor.chara_info().Gold
		kill_value = actor.chara_info().BattleAttribute.KillValue

		for _, v := range all_award_items {
			actor.stat_if_chengzhuang(v.Item.TypeId, v.Num)
		}

		actor.save()
	})

	if err != msg.Success {
		return &msg.BattleEndResponse{Err: err}
	}

	return &msg.BattleEndResponse{
		State:              chara_state,
		KillValueIncreased: kill_value_increased,
		KillValue:          kill_value,
		GoldIncreased:      gold_increased,
		GoldExtIncreased:   gold_ext_increased,
		Gold:               gold,
		UnpickupItems:      unpickup_items,
		AllAwardItems:      all_award_items,
		BagDiff:            bag_diff,
	}
}
