package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) DeleteNoticeInfo(req *msg.Request, arg *msg.DeleteNoticeInfoRequest) *msg.DeleteNoticeInfoResponse {
	g.db.Notice().DeleteNotice(req.Uid, arg.NoticeInfoId)
	return &msg.DeleteNoticeInfoResponse{}
}
