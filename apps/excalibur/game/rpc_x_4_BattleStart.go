package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) BattleStart(req *msg.Request, arg *msg.BattleStartRequest) *msg.BattleStartResponse {

	var chara_state string

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)
		battle_with := odm.Npc{NpcId: arg.NpcId, NpcTypeId: arg.NpcTypeId}
		actor.battle_start(battle_with, arg.LuckyValue)
		chara_state = actor.chara_state()
		actor.save()
	})

	if err != msg.Success {
		return &msg.BattleStartResponse{Err: err}
	}
	return &msg.BattleStartResponse{
		State: chara_state,
	}
}
