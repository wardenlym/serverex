package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) JobUpgrade(req *msg.Request, arg *msg.JobUpgradeRequest) *msg.JobUpgradeResponse {
	user := g.get_user(req.Uid)
	chara := user.Characters[arg.CharacterId]

	if uint(chara.BattleAttribute.KillValue) < arg.CostKillValue {
		return &msg.JobUpgradeResponse{Err: msg.ErrAvalibleCountNotEnough}
	}
	killValue := chara.BattleAttribute.KillValue - arg.CostKillValue

	is_upstar, index, star := func() (bool, int, uint) {
		for i, v := range chara.BattleAttribute.JobTree.JobStatus {
			if v.JobId == arg.JobStatus.JobId {
				return true, i, v.JobStar + 1
			}
		}
		return false, -1, 0
	}()

	if is_upstar {
		chara.BattleAttribute.JobTree.JobStatus[index].JobStar = star
	} else {
		chara.BattleAttribute.JobTree.JobStatus = append(chara.BattleAttribute.JobTree.JobStatus, arg.JobStatus)
	}

	chara.BattleAttribute.KillValue = killValue

	user.Characters[arg.CharacterId] = chara

	g.save_user(user)

	return &msg.JobUpgradeResponse{
		JobTree:   chara.BattleAttribute.JobTree,
		KillValue: killValue,
	}
}
