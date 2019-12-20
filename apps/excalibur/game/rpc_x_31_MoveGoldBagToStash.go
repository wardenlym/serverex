package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) MoveGoldBagToStash(req *msg.Request, arg *msg.MoveGoldBagToStashRequest) *msg.MoveGoldBagToStashResponse {
	user := g.get_user(req.Uid)
	chara := user.Characters[arg.CharacterId]

	user.Characters[arg.CharacterId] = chara
	g.save_user(user)

	if chara.Gold < arg.Gold {
		return &msg.MoveGoldBagToStashResponse{Err: msg.ErrGoldNotEnough}
	}

	user.Stash.Gold += arg.Gold
	chara.Gold -= arg.Gold

	user.Characters[arg.CharacterId] = chara
	g.save_user(user)
	return &msg.MoveGoldBagToStashResponse{
		StashGold: user.Stash.Gold,
		BagGold:   chara.Gold,
	}
}
