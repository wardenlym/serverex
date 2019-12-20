package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) Reborn(req *msg.Request, arg *msg.RebornRequest) *msg.RebornResponse {

	var chara_state string
	var diamond uint

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)
		actor.reborn(arg.Diamond)
		chara_state = actor.chara_state()
		diamond = actor.user_info().Diamond
		actor.save()
	})

	if err != msg.Success {
		return &msg.RebornResponse{Err: err}
	}
	return &msg.RebornResponse{
		State:   chara_state,
		Diamond: diamond,
	}
}
