package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) DestroyItem(req *msg.Request, arg *msg.DestroyItemRequest) *msg.DestroyItemResponse {

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)
		actor.destory_in_bag([]int{arg.CellIndex})
		actor.save()
	})

	if err != msg.Success {
		return &msg.DestroyItemResponse{Err: err}
	}
	return &msg.DestroyItemResponse{}
}
