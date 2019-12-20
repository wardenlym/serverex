package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (a *Actor) puzzle_start(puzzle_info odm.PuzzleInfo) *Actor {
	if a.chara.DungeonStatus.State != odm.S_InStage {
		raise_error(msg.ErrInvalidDungeonState)
	}
	a.chara.DungeonStatus.CurrentStageInfo.PuzzleInfo = puzzle_info
	return a
}

func (a *Actor) puzzle_update(process int) *Actor {
	if a.chara.DungeonStatus.State != odm.S_InStage {
		raise_error(msg.ErrInvalidDungeonState)
	}
	a.chara.DungeonStatus.CurrentStageInfo.PuzzleInfo.Process = process
	return a
}

func (a *Actor) puzzle_end(selected int) *Actor {
	if a.chara.DungeonStatus.State != odm.S_InStage {
		raise_error(msg.ErrInvalidDungeonState)
	}
	a.chara.DungeonStatus.CurrentStageInfo.PuzzleInfo.Process =
		a.chara.DungeonStatus.CurrentStageInfo.PuzzleInfo.Total
	a.chara.DungeonStatus.CurrentStageInfo.PuzzleInfo.Selected = selected
	return a
}
