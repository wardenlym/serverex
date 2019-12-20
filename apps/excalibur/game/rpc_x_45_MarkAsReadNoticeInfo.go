package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) MarkAsReadNoticeInfo(req *msg.Request, arg *msg.MarkAsReadNoticeInfoRequest) *msg.MarkAsReadNoticeInfoResponse {

	user := g.get_user(req.Uid)

	g.db.Notice().MarkAsReadNotice(user.Uid, arg.NoticeInfoId)
	g.save_user(user)

	return &msg.MarkAsReadNoticeInfoResponse{}
}
