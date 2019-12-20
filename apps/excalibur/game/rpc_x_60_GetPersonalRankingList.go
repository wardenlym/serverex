package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) GetPersonalRankingList(req *msg.Request, arg *msg.GetPersonalRankingListRequest) *msg.GetPersonalRankingListResponse {
	var info []odm.PersonalRankingInfo
	err := try_do(func() {
		actor := g.use_actor(req.Uid)
		me := actor.make_current_user_personal_ranking_info()
		g.db.Ranking().InsertPersonalRanking(me)

		info = g.db.Ranking().GetPersonalRankingList(arg.RankingType, arg.IndexFrom, arg.Count)
	})

	if err != msg.Success {
		return &msg.GetPersonalRankingListResponse{Err: err}
	}

	return &msg.GetPersonalRankingListResponse{
		Info: info,
	}
}
