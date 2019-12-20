package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) GetAchievementAward(req *msg.Request, arg *msg.GetAchievementAwardRequest) *msg.GetAchievementAwardResponse {

	var achievement_score uint
	var bag_diff []odm.LocatedCell
	var diamond uint
	var gold uint

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)

		gold_add, diamond_add, plain_award_add, a_score_add := func(id int) (uint, uint, []odm.PlainCell, uint) {

			info, exist := csv.Table_AchievementData[id]
			if !exist {
				raise_error(msg.ErrInvalidAchievementId)
			}

			award := []odm.PlainCell{}
			for _, typeid := range info.Goods_rewards {
				award = append(award, odm.PlainCell{
					TypeId: typeid,
					Num:    1,
				})
			}

			return uint(info.Gold_rewards), uint(info.Diamond_rewards), award, uint(info.Point)

		}(arg.AchievementId)

		award_add := cell_plains_to_cells(plain_award_add)

		actor.put_to_bag(award_add)
		actor.incr_diamond(diamond_add)
		actor.incr_gold(gold_add)
		actor.user.StatInfo.AchievementScore += a_score_add

		achievement_score = actor.user_info().StatInfo.AchievementScore
		diamond = actor.user_info().Diamond
		gold = actor.chara_info().Gold
		bag_diff = actor.diff_bag()
		actor.save()
	})

	if err != msg.Success {
		return &msg.GetAchievementAwardResponse{Err: err}
	}

	return &msg.GetAchievementAwardResponse{
		AchievementScore: achievement_score,
		BagDiff:          bag_diff,
		Diamond:          diamond,
		Gold:             gold,
	}
}
