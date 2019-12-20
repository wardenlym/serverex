package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) OpenBox(req *msg.Request, arg *msg.OpenBoxRequest) *msg.OpenBoxResponse {
	user := g.get_user(req.Uid)
	chara := user.Characters[arg.CharacterId]

	unpicked := []odm.Cell{}
	newcells := []odm.Cell{}
	award := []odm.Cell{}

	gold_added := uint(0)

	stage_id, npc_type_id, lucky :=
		chara.DungeonStatus.CurrentStageInfo.StageId, arg.TypeId, chara.BattleAttribute.LuckyValue
	award = gen_loot_item(stage_id, npc_type_id, int(lucky))

	newcells, unpicked = cell_move_from_to(award, chara.Bag.Cells)

	bag_diff := cell_diff(chara.Bag.Cells, newcells)

	chara.Bag.Cells = newcells
	chara.DungeonStatus.UnpickedLoot = unpicked
	chara.DungeonStatus.LootTotal = append(chara.DungeonStatus.LootTotal, award...)

	// TODO: openbox list error
	chara.DungeonStatus.OpenBoxList = append(chara.DungeonStatus.OpenBoxList, odm.Npc{NpcTypeId: arg.TypeId})

	chara.Gold += gold_added
	chara.DungeonStatus.GoldTotal += gold_added

	user.Characters[arg.CharacterId] = chara
	g.save_user(user)
	return &msg.OpenBoxResponse{
		UnpickupItems: unpicked,
		AllAwardItems: award,
		BagDiff:       bag_diff,
	}
}
