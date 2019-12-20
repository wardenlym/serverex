package game

import (
	"time"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/util"
)

func (a *Actor) chara_state() string {
	return a.chara.DungeonStatus.State
}

func (a *Actor) enter_dungeon() *Actor {

	if a.chara.DungeonStatus.State != odm.S_InTown {
		raise_error(msg.ErrInvalidDungeonState)
	}
	a.chara.DungeonStatus = odm.NewDungeonStatus()
	a.chara.DungeonStatus.EnterAt = time.Now().Unix()
	a.chara.BattleAttribute = odm.NewAttribute(a.chara.HeroId)
	a.chara.DungeonStatus.State = odm.S_InDungeon
	return a
}

func (a *Actor) exit_dungeon(is_safe_exit bool) *Actor {

	a.chara.Bag.Cells = cell_remove_cannot_take_back(a.chara.Bag.Cells)
	a.chara.Bag.Equipments = cell_remove_cannot_take_back(a.chara.Bag.Equipments)

	if !is_safe_exit || a.chara.DungeonStatus.State == odm.S_Dead {

		a.chara.Bag.Cells = cell_remove_all(a.chara.Bag.Cells)

		if !a.is_vip_godbless() {
			a.chara.Bag.Equipments = cell_remove_random_1(a.chara.Bag.Equipments)
		}
	}

	a.user.Stash.Gold += a.chara.Gold
	a.chara.Gold = 0
	a.chara.DungeonStatus = odm.NewDungeonStatus()
	a.chara.BattleAttribute = odm.NewAttribute(a.chara.HeroId)
	a.chara.DungeonStatus.State = odm.S_InTown
	return a
}

func (a *Actor) current_dungeon_stat() ([]odm.Npc, []odm.Cell, uint, uint) {

	npc_list := []odm.Npc{}
	item_list := make([]odm.Cell, len(a.chara.DungeonStatus.LootTotal))
	gold_total := a.chara.DungeonStatus.GoldTotal
	gold_ext_total := a.chara.DungeonStatus.GoldExtTotal

	for _, v := range a.chara.DungeonStatus.KilledNpcList {
		npc_list = append(npc_list, v)
	}

	copy(item_list, a.chara.DungeonStatus.LootTotal)

	return npc_list, item_list, gold_total, gold_ext_total
}

func (a *Actor) enter_stage(stage_info odm.StageInfo) *Actor {
	if a.chara.DungeonStatus.State != odm.S_InDungeon {
		raise_error(msg.ErrInvalidDungeonState)
	}
	a.chara.DungeonStatus.State = odm.S_InStage
	a.chara.DungeonStatus.CurrentStageInfo = stage_info
	return a
}

func (a *Actor) exit_stage(clear bool) *Actor {

	if a.chara.DungeonStatus.State == odm.S_Dead {
		a.chara.DungeonStatus.State = odm.S_Dead
	} else {
		a.chara.DungeonStatus.State = odm.S_InDungeon
	}

	if clear {
		a.chara.DungeonStatus.StageProgress =
			append(a.chara.DungeonStatus.StageProgress, a.chara.DungeonStatus.CurrentStageInfo)

		if a.chara.RankingInfo.BestStageId < a.chara.DungeonStatus.CurrentStageInfo.StageId {
			a.chara.RankingInfo.BestStageId = a.chara.DungeonStatus.CurrentStageInfo.StageId
		}
		if a.user.RankingInfo.BestStageId < a.chara.DungeonStatus.CurrentStageInfo.StageId {
			a.user.RankingInfo.BestStageId = a.chara.DungeonStatus.CurrentStageInfo.StageId
		}
	}
	a.chara.DungeonStatus.CurrentBattleNpc = odm.Npc{}
	a.chara.DungeonStatus.CurrentStageInfo = odm.StageInfo{}
	return a
}

func (a *Actor) update_dungeon_stat(snapshot odm.DungeonAttributeSnapshot, stat odm.DungeonStageStat) *Actor {
	a.chara.DungeonStatus.AttributeSnapshot = snapshot
	a.chara.DungeonStatus.StageStat.Damage += stat.Damage
	a.chara.DungeonStatus.StageStat.Kill += stat.Kill
	a.chara.DungeonStatus.StageStat.Healing += stat.Healing

	return a
}

func (a *Actor) try_make_new_dungeon_ranking_record(stage_id int) (odm.DungeonRankingInfo, bool) {

	info := odm.DungeonRankingInfo{
		Uid:               a.user.Uid,
		CharacterId:       a.chara_id,
		Nickname:          a.user.Nickname,
		Time:              time.Now().Unix() - a.chara.DungeonStatus.EnterAt,
		Date:              time.Now().Unix(),
		ChapterId:         stageid_to_chapterid(stage_id),
		AttributeSnapshot: a.chara.DungeonStatus.AttributeSnapshot,
		StageStat:         a.chara.DungeonStatus.StageStat,
		Equipments:        a.chara.Bag.Equipments,
		JobTree:           a.chara.BattleAttribute.JobTree,
		ThreadId:          util.NewSnowflakeID(),
	}

	if len(a.user.BestDungeonRanking) == 0 {
		a.user.BestDungeonRanking = append(a.user.BestDungeonRanking, info)
		return info, true
	} else {
		for i, v := range a.user.BestDungeonRanking {
			if v.ChapterId == info.ChapterId &&
				info.Time < v.Time {
				a.user.BestDungeonRanking[i] = info
				return info, true
			}
		}
	}

	return info, false
}

func (a *Actor) battle_start(npc odm.Npc, lucky uint) *Actor {

	if a.chara.DungeonStatus.State != odm.S_InStage {
		raise_error(msg.ErrInvalidDungeonState)
	}
	a.chara.DungeonStatus.State = odm.S_InBattle
	a.chara.DungeonStatus.CurrentBattleNpc = npc
	a.chara.BattleAttribute.LuckyValue = lucky
	return a
}

func (a *Actor) current_battle_with() odm.Npc {

	if a.chara.DungeonStatus.State == odm.S_InBattle {
		return a.chara.DungeonStatus.CurrentBattleNpc
	}
	raise_error(msg.ErrInvalidDungeonState)
	return odm.Npc{}
}

func (a *Actor) battle_end(result string, kill_value_increased uint, gold_increased, gold_ext_increased uint) *Actor {

	if a.chara.DungeonStatus.State != odm.S_InBattle {
		raise_error(msg.ErrInvalidDungeonState)
	}

	npc_id := a.chara.DungeonStatus.CurrentBattleNpc.NpcId
	a.chara.DungeonStatus.CurrentBattleNpc = odm.Npc{}

	switch result {
	case odm.Battle_Victory:
		a.chara.DungeonStatus.State = odm.S_InStage

		a.chara.DungeonStatus.CurrentStageInfo.NpcList[npc_id].IsDead = true
		a.chara.DungeonStatus.KilledNpcList =
			append(a.chara.DungeonStatus.KilledNpcList,
				a.chara.DungeonStatus.CurrentStageInfo.NpcList[npc_id])

		a.chara.BattleAttribute.KillValue += kill_value_increased
		a.chara.Gold += (gold_increased + gold_ext_increased)
		a.chara.DungeonStatus.GoldTotal += gold_increased
		a.chara.DungeonStatus.GoldTotal += gold_ext_increased

		a.user.StatInfo.KillMonsterCount++
	case odm.Battle_Escape:
		a.chara.DungeonStatus.State = odm.S_InStage

	case odm.Battle_Dead:
		a.chara.DungeonStatus.State = odm.S_Dead
	default:
		a.chara.DungeonStatus.State = odm.S_InStage
	}
	return a
}

func (a *Actor) reborn(diamond uint) *Actor {
	if a.chara.DungeonStatus.State != odm.S_Dead {
		raise_error(msg.ErrInvalidDungeonState)
	}
	a.cost_diamond(diamond)
	a.chara.DungeonStatus.State = odm.S_InStage
	return a
}
