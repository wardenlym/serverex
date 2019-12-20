package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) GM_ResetJob(req *msg.Request, arg *msg.GM_ResetJobRequest) *msg.GM_ResetJobResponse {
	user := g.get_user(req.Uid)
	chara := user.Characters[arg.CharacterId]

	chara.BattleAttribute.KillValue = arg.KillValue
	chara.BattleAttribute.JobTree = odm.NewJobTree(chara.HeroId)

	user.Characters[arg.CharacterId] = chara
	g.save_user(user)

	return &msg.GM_ResetJobResponse{
		JobTree:   chara.BattleAttribute.JobTree,
		KillValue: uint(chara.BattleAttribute.KillValue),
	}
}
