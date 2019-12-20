package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) GetPersonalRankingInfo(req *msg.Request, arg *msg.GetPersonalRankingInfoRequest) *msg.GetPersonalRankingInfoResponse {

	var info odm.PersonalRankingInfo
	err := try_do(func() {
		actor := g.use_actor(req.Uid)
		info = actor.make_current_user_personal_ranking_info()
		g.db.Ranking().InsertPersonalRanking(info)

		list := g.db.Ranking().GetPersonalRankingList(arg.RankingType, 0, 200)

		for _, v := range list {
			if v.Uid == info.Uid {
				info = v
				break
			}
		}
		if info.Rank == 0 {
			info.Rank = 999
		}
	})

	if err != msg.Success {
		return &msg.GetPersonalRankingInfoResponse{Err: err}
	}

	return &msg.GetPersonalRankingInfoResponse{
		Info: info,
	}
}
