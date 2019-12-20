package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) GetDungeonRankingList(req *msg.Request, arg *msg.GetDungeonRankingListRequest) *msg.GetDungeonRankingListResponse {

	info := g.db.Ranking().GetDungeonRankingList(arg.ChapterId, arg.IndexFrom, arg.Count)
	return &msg.GetDungeonRankingListResponse{Info: info}
}
