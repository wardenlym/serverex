package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) ExploreEnd(req *msg.Request, arg *msg.ExploreEndRequest) *msg.ExploreEndResponse {

	var explore_info odm.ExploreInfo
	var award_items []odm.Cell
	var bag_diff []odm.LocatedCell

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)

		hero_id, hero_level, use_time := actor.explore_end(arg.ExploreCharacterId, arg.ExploreStageId)
		explore_info = actor.user_info().ExploreInfo

		plains := calc_explore_award(arg.ExploreStageId, hero_id, hero_level, use_time)
		award_items = cell_plains_to_cells(plains)
		unpicked := actor.pickup_to_bag(award_items)
		if len(unpicked) > 0 {
			raise_error(msg.ErrBagCellNotEnough)
		}

		bag_diff = actor.diff_bag()
		actor.save()
	})

	if err != msg.Success {
		return &msg.ExploreEndResponse{Err: err}
	}
	return &msg.ExploreEndResponse{
		ExploreInfo: explore_info,
		AwardItems:  award_items,
		BagDiff:     bag_diff,
	}
}
