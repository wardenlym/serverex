package game

import (
	"math/rand"
	"time"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/lib/util"

	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

type World struct {
}

func (g *GameService) world() *World {
	return &World{}
}

func (w *World) gen_stage_info(stage_id int) odm.StageInfo {
	info, exist := csv.Table_StageData[stage_id]
	if !exist {
		raise_error(msg.ErrInvalidStageId)
	}

	mapid, npclist := make_npc_list(info)

	boxlist := []odm.Box{}
	for _, v := range info.Box {
		boxlist = append(boxlist, odm.Box{
			Id:     len(boxlist),
			TypeId: v,
		})
	}

	return odm.StageInfo{
		StageId:      stage_id,
		MapId:        mapid,
		EntranceType: info.EntranceType,
		Seed:         int(rand.Int31()),
		NpcList:      npclist,
		BoxList:      boxlist,
	}
}

func make_npc_list(info csv.StageData) (mapid int, npclist []odm.Npc) {
	npclist = []odm.Npc{}

	index, mapid := func() (int, int) {
		if len(info.MapId) == 1 {
			return 0, info.MapId[0]
		} else {
			index := util.RandomIn(0, len(info.MapId))
			return index, info.MapId[index]
		}
	}()

	max := info.Cave_number[index] + 1

	count := util.RandomIn(info.Random_monster_min, max)
	index_list := util.RandomIndexList(len(info.NormalMonsterList), count)

	for _, i := range index_list {
		npclist = append(npclist, odm.Npc{
			NpcId:     len(npclist),
			NpcTypeId: info.NormalMonsterList[i],
		})
	}

	return mapid, npclist
}

func (w *World) gen_onsale_info() {
}

func (w *World) calc_kill_value(npc_type_id int) uint {
	if v, exist := csv.Table_MonsterData[npc_type_id]; exist {
		return uint(v.Killsvalue)
	}
	return 0
}

func (w *World) gen_loot_award(stage_id, npc_type_id, lucky int) ([]odm.Cell, uint) {
	if v, exist := csv.Table_MonsterData[npc_type_id]; exist {

		N := stageid_to_chapterid(stage_id) / 1000
		gold := N*15 + (v.Monster_type * v.Monster_type * (((stage_id - N*1000) / 5) + 1) * 5)

		award := gen_loot_item(stage_id, npc_type_id, lucky)
		return award, uint(gold)
	}
	return []odm.Cell{}, 0
}

func (w *World) calc_score_dungeon(status odm.DungeonStatus) odm.DungeonScore {

	last_stage_id := consts.FirstStageId
	if len(status.StageProgress) > 0 {
		last_stage_id = status.StageProgress[len(status.StageProgress)-1].StageId
	}

	return odm.DungeonScore{
		BestStageId:      last_stage_id,
		OpenBoxScore:     calc_score_openbox(status.OpenBoxList),
		KillMonsterScore: calc_score_openbox(status.KilledNpcList),
		KillBossScore:    calc_score_openbox(status.KilledNpcList),
		HealthScore:      int(status.AttributeSnapshot.Hp),
		GearScore:        status.AttributeSnapshot.GearScore,
		AwardItemScore:   calc_score_award_item(cell_cells_to_plains(status.LootTotal)),
		TimeCost:         time.Now().Unix() - status.EnterAt,
	}
}

func monster_type(npc_type_id int) int {
	monster_type, err := lookup_monster_type(npc_type_id)
	if err != nil {
		return 0
	}
	return monster_type
}

const (
	PuTongGuai      = 1
	JingYingGuai    = 2
	XiaoLingZhuGuai = 3
	DaLingZhuGuai   = 4
	QingTongXiang   = 5
	MiYinXiang      = 6
	HuangJinXiang   = 7
	CangBaoGuan     = 8
)

func calc_score_openbox(npc []odm.Npc) int {
	var score int
	for _, v := range npc {
		switch monster_type(v.NpcTypeId) {
		case QingTongXiang:
			score += 100
		case MiYinXiang:
			score += 200
		case HuangJinXiang:
			score += 500
		case CangBaoGuan:
			score += 100
		}
	}
	return score
}

func calc_score_kill_monster(npc []odm.Npc) int {
	var score int
	for _, v := range npc {
		switch monster_type(v.NpcTypeId) {
		case PuTongGuai:
			score += 100
		case JingYingGuai:
			score += 250
		case XiaoLingZhuGuai:
			score += 400
		}
	}
	return score
}

func calc_score_kill_boss(npc []odm.Npc) int {
	var score int
	for _, v := range npc {
		switch monster_type(v.NpcTypeId) {
		case DaLingZhuGuai:
			score += 10000
		}
	}
	return score
}

func calc_score_award_item(items []odm.PlainCell) int {

	var score int
	for _, v := range items {
		if cell_is_can_carry_back(v.TypeId) {
			score += int(v.Num) * 500
		}
	}
	return score
}
