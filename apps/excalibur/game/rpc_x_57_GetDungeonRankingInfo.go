package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/util"
)

func (g *GameService) GetDungeonRankingInfo(req *msg.Request, arg *msg.GetDungeonRankingInfoRequest) *msg.GetDungeonRankingInfoResponse {
	var info odm.DungeonRankingInfo

	err := try_do(func() {
		actor := g.use_actor(req.Uid)
		for _, v := range actor.user.BestDungeonRanking {
			if v.ChapterId == arg.ChapterId {
				info = v
				break
			}
		}

		list := g.db.Ranking().GetDungeonRankingList(arg.ChapterId, 0, 200)

		for _, v := range list {
			if v.Uid == info.Uid {
				info = v
				break
			}
		}
		if info.Rank == 0 {
			info.Rank = 999
			info.ThreadId = util.NewSnowflakeID()
		}
	})

	if err != msg.Success {
		return &msg.GetDungeonRankingInfoResponse{Err: err}
	}

	return &msg.GetDungeonRankingInfoResponse{Info: info}
}
