package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) EnterDungeon(req *msg.Request, arg *msg.EnterDungeonRequest) *msg.EnterDungeonResponse {

	var chara_state string

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)
		actor.enter_dungeon()
		chara_state = actor.chara_state()
		actor.save()
	})

	if err != msg.Success {
		return &msg.EnterDungeonResponse{Err: err}
	}
	return &msg.EnterDungeonResponse{
		State: chara_state,
	}
}
