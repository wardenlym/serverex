package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) RefreshOnSale(req *msg.Request, arg *msg.RefreshOnSaleRequest) *msg.RefreshOnSaleResponse {

	var diamond uint
	err := try_do(func() {
		actor := g.use_actor(req.Uid)
		if actor.user.RankingInfo.BestStageId == 0 {
			actor.user.RankingInfo.BestStageId = consts.FirstStageId
		}

		actor.cost_diamond(100)
		actor.user.OnSaleInfos = makeOnSaleInfoFromTable(actor.user.RankingInfo.BestStageId)
		diamond = actor.user.Diamond
		actor.save()
	})

	if err != msg.Success {
		return &msg.RefreshOnSaleResponse{Err: err}
	}
	return &msg.RefreshOnSaleResponse{
		Diamond: diamond,
	}
}
