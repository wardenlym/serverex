package game

import (
	"math"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func calc_expand_bag_cost(n int) uint {
	if n >= 75 {
		raise_error(msg.ErrOutOfRangeLimit)
	}

	cost := uint(math.Ceil(math.Pow(2, float64(n/15-1)) * 120))
	return cost
}

func (g *GameService) ExpandBagCell(req *msg.Request, arg *msg.ExpandBagCellRequest) *msg.ExpandBagCellResponse {

	var diamond uint
	var start int
	var offset int

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)

		start = len(actor.chara.Bag.Cells)
		offset = consts.BagExpandCount
		cost := calc_expand_bag_cost(len(actor.chara.Bag.Cells))
		actor.cost_diamond(cost)

		for i := 0; i < offset; i++ {
			actor.chara.Bag.Cells = append(actor.chara.Bag.Cells, odm.Cell{})
		}

		diamond = actor.user.Diamond
		actor.save()
	})

	if err != msg.Success {
		return &msg.ExpandBagCellResponse{Err: err}
	}

	return &msg.ExpandBagCellResponse{
		Diamond: diamond,
		Start:   start,
		Offset:  offset,
	}
}

func (g *GameService) ExpandBagCell2(req *msg.Request, arg *msg.ExpandBagCellRequest) *msg.ExpandBagCellResponse {

	user := g.get_user(req.Uid)
	chara := user.Characters[arg.CharacterId]

	diamond := user.Diamond

	start := len(chara.Bag.Cells)
	offset := consts.BagExpandCount

	for i := 0; i < offset; i++ {
		chara.Bag.Cells = append(chara.Bag.Cells, odm.Cell{})
	}

	user.Diamond = diamond
	user.Characters[arg.CharacterId] = chara
	g.save_user(user)
	return &msg.ExpandBagCellResponse{
		Diamond: user.Diamond,
		Start:   start,
		Offset:  offset,
	}
}
