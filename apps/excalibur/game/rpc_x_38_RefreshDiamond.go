package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) RefreshDiamond(req *msg.Request, arg *msg.RefreshDiamondRequest) *msg.RefreshDiamondResponse {

	var diamond uint

	err := try_do(func() {
		actor := g.use_actor(req.Uid)
		diamond = actor.user.Diamond
	})

	if err != msg.Success {
		return &msg.RefreshDiamondResponse{Err: err}
	}
	return &msg.RefreshDiamondResponse{
		Diamond: diamond,
	}
}
