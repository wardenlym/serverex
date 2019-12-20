package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) MoveGoldStashToBag(req *msg.Request, arg *msg.MoveGoldStashToBagRequest) *msg.MoveGoldStashToBagResponse {
	user := g.get_user(req.Uid)
	chara := user.Characters[arg.CharacterId]

	user.Characters[arg.CharacterId] = chara
	g.save_user(user)

	// if user.Gold < arg.Gold {
	// 	return &msg.MoveGoldStashToBagResponse{Err: msg.ErrGoldNotEnough}
	// }

	user.Stash.Gold -= arg.Gold
	chara.Gold += arg.Gold

	user.Characters[arg.CharacterId] = chara
	g.save_user(user)
	return &msg.MoveGoldStashToBagResponse{
		StashGold: user.Stash.Gold,
		BagGold:   chara.Gold,
	}
}
