package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) KillCharacter(req *msg.Request, arg *msg.KillCharacterRequest) *msg.KillCharacterResponse {

	var chara_state string
	var bag_diff []odm.LocatedCell

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)
		actor.chara_info().DungeonStatus.State = odm.S_Dead
		chara_state = actor.chara_state()
		actor.save()
	})

	if err != msg.Success {
		return &msg.KillCharacterResponse{Err: err}
	}

	return &msg.KillCharacterResponse{
		State:   chara_state,
		BagDiff: bag_diff,
	}
}
