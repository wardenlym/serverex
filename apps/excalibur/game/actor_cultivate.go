package game

import (
	"time"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/util"
)

func (a *Actor) cost_diamond(count uint) *Actor {

	if a.user.Diamond < count {
		raise_error(msg.ErrDiamondNotEnough)
	}
	a.user.Diamond -= count

	return a
}

func (a *Actor) cost_gold(count uint) *Actor {

	if a.chara.DungeonStatus.State == odm.S_InDungeon {
		if a.chara.Gold < count {
			raise_error(msg.ErrGoldNotEnough)
		} else {
			a.chara.Gold -= count
		}
		return a
	}

	if a.user.Stash.Gold+a.chara.Gold < count {
		raise_error(msg.ErrGoldNotEnough)
	}

	if a.chara.Gold > count {
		a.chara.Gold -= count
	} else {
		d := count - a.chara.Gold
		a.chara.Gold = 0
		a.user.Stash.Gold -= d
	}

	return a
}

func (a *Actor) incr_diamond(count uint) *Actor {
	a.user.Diamond += count
	return a
}

func (a *Actor) incr_gold(count uint) *Actor {
	a.chara.Gold += count
	return a
}

func (a *Actor) diff_diamond() *Actor {
	return a
}

func (a *Actor) diff_gold() *Actor {
	return a
}

func (a *Actor) make_current_user_personal_ranking_info() odm.PersonalRankingInfo {
	var gold uint

	for _, v := range a.user.Characters {
		if v.Gold > gold {
			gold = v.Gold
		}
	}
	gold += a.user.Stash.Gold

	return odm.PersonalRankingInfo{
		Uid:         a.user.Uid,
		Nickname:    a.user.Nickname,
		Date:        time.Now().Unix(),
		Gold:        gold,
		Kill:        a.user.StatInfo.KillMonsterCount,
		Achievement: a.user.StatInfo.AchievementScore,
		ThreadId:    util.NewSnowflakeID(),
	}
}
