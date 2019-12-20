package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func calc_expand_stash_cost(n int) uint {
	if n >= 160 {
		return 1800
	}

	cost := uint((n-15)/5*60 + 60)
	return cost
}

func (g *GameService) ExpandStashCell(req *msg.Request, arg *msg.ExpandStashCellRequest) *msg.ExpandStashCellResponse {

	var diamond uint
	var start int
	var offset int

	err := try_do(func() {
		actor := g.use_actor(req.Uid)

		start = len(actor.user.Stash.Cells)
		offset = consts.StashExpandCount
		cost := calc_expand_stash_cost(len(actor.user.Stash.Cells))
		actor.cost_diamond(cost)

		for i := 0; i < offset; i++ {
			actor.user.Stash.Cells = append(actor.user.Stash.Cells, odm.Cell{})
		}

		diamond = actor.user.Diamond
		actor.save()
	})

	if err != msg.Success {
		return &msg.ExpandStashCellResponse{Err: err}
	}

	return &msg.ExpandStashCellResponse{
		Diamond: diamond,
		Start:   start,
		Offset:  offset,
	}
}
