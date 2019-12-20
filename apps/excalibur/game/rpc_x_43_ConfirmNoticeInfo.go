package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) ConfirmNoticeInfo(req *msg.Request, arg *msg.ConfirmNoticeInfoRequest) *msg.ConfirmNoticeInfoResponse {
	user := g.get_user(req.Uid)
	chara := user.Characters[arg.CharacterId]

	info, err := g.db.Notice().GetNotice(user.Uid, arg.NoticeInfoId)
	if err != nil {
		return &msg.ConfirmNoticeInfoResponse{Err: msg.ErrInvalidNoticeId}
	}

	user.Diamond += info.Diamond
	chara.Gold += info.Gold

	cells, unpick := cell_move_from_to(info.Gift, chara.Bag.Cells)
	if len(unpick) != 0 {
		return &msg.ConfirmNoticeInfoResponse{Err: msg.ErrBagCellNotEnough}
	}

	g.db.Notice().ConfirmNotice(user.Uid, arg.NoticeInfoId)
	bag_diff := cell_diff(chara.Bag.Cells, cells)
	chara.Bag.Cells = cells

	user.Characters[arg.CharacterId] = chara
	g.save_user(user)
	return &msg.ConfirmNoticeInfoResponse{
		BagGold:   chara.Gold,
		StashGold: user.Stash.Gold,
		Diamond:   user.Diamond,
		BagDiff:   bag_diff,
	}
}
