package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) ListNoticeInfo(req *msg.Request, arg *msg.ListNoticeInfoRequest) *msg.ListNoticeInfoResponse {
	list := g.db.Notice().GetNoticeList(req.Uid)
	return &msg.ListNoticeInfoResponse{
		NoticeInfos: list,
	}
}
