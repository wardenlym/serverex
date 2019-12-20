package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/log"
)

func (g *GameService) LevelUp2(req *msg.Request, arg *msg.LevelUpRequest) *msg.LevelUpResponse {
	user := g.get_user(req.Uid)
	chara := user.Characters[arg.CharacterId]

	if chara.Level != arg.CurrentLevel {
		return &msg.LevelUpResponse{Err: msg.ErrLevelNotEnough}
	}

	level := chara.Level + 1

	ok, cells := cell_remove_materials_if_enough(arg.Requires, chara.Bag.Cells)
	if !ok {
		return &msg.LevelUpResponse{Err: msg.ErrAvalibleCountNotEnough}
	}

	chara.Level = level
	chara.Bag.Cells = cells

	user.Characters[arg.CharacterId] = chara

	g.save_user(user)

	return &msg.LevelUpResponse{
		CurrentLevel: level,
	}
}

func (a *Actor) open_rune_slot() {
	info, exist := csv.Table_RuneopenData[a.chara.Level]
	if !exist {
		log.Error("not exist")
		return
	}

	to_open := []int{}
	log.Error("::", a.chara.HeroId)
	switch a.chara.HeroId {
	case 10000:
		to_open = info.Hero_10000
	case 20000:
		to_open = info.Hero_20000
	case 30000:
		to_open = info.Hero_30000
	default:
		log.Error("default")
	}

	log.Error("::：", to_open)
	for _, v := range to_open {
		log.Error(v, a.chara.RuneTree.Cells[v])
		a.chara.RuneTree.Cells[v].Enable = true
	}
	log.Error("::：", a.chara.RuneTree.Cells)
}

func (g *GameService) LevelUp(req *msg.Request, arg *msg.LevelUpRequest) *msg.LevelUpResponse {

	var level int
	var gold uint
	var stash_gold uint
	var bag_diff []odm.LocatedCell
	var rune_tree odm.RuneTree

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)

		if actor.chara.Level != arg.CurrentLevel {
			raise_error(msg.ErrLevelNotEnough)
		}

		actor.cost_gold(arg.CostGold)

		actor.chara.Level += 1

		actor.open_rune_slot()

		actor.cost_requires_in_bag(arg.Requires)
		level = actor.chara.Level
		gold = actor.chara.Gold
		stash_gold = actor.user.Stash.Gold
		bag_diff = actor.diff_bag()
		rune_tree = actor.chara.RuneTree
		actor.save()
	})

	if err != msg.Success {
		return &msg.LevelUpResponse{Err: err}
	}

	return &msg.LevelUpResponse{
		CurrentLevel: level,
		Gold:         gold,
		StashGold:    stash_gold,
		BagDiff:      bag_diff,
		RuneTree:     rune_tree,
	}
}
