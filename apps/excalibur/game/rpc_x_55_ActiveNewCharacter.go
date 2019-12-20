package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) ActiveNewCharacter(req *msg.Request, arg *msg.ActiveNewCharacterRequest) *msg.ActiveNewCharacterResponse {

	var chara_id string
	var diamond uint
	var gold uint
	var stash_gold uint
	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)
		chara_id = actor.next_character_id()
		actor.new_character(chara_id, arg.HeroId).
			save()
		diamond = actor.user_info().Diamond
		stash_gold = actor.user_info().Stash.Gold
		gold = actor.chara_info().Gold
	})
	if err != msg.Success {
		return &msg.ActiveNewCharacterResponse{Err: err}
	}
	return &msg.ActiveNewCharacterResponse{
		CharacterId: chara_id,
		Diamond:     diamond,
		Gold:        gold,
		StashGold:   stash_gold,
	}
}
