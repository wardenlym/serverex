package game

import (
	"time"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

const (
	ActivitySignIn30Day_Process_End = 30
)

func (g *GameService) GetActivtyAwardSignIn30Day(req *msg.Request, arg *msg.GetActivtyAwardSignIn30DayRequest) *msg.GetActivtyAwardSignIn30DayResponse {

	var activity_id int
	var process_info odm.ActivityProcessInfo
	bag_diff := []odm.LocatedCell{}
	var diamond uint
	var gold uint

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)
		if actor.user_info().OperationActivity.SignIn30Day.Process >= ActivitySignIn30Day_Process_End {
			actor.user_info().OperationActivity.SignIn30Day.Process = 0
		}

		if !is_flushed_time_valid_everyday(actor.user_info().OperationActivity.SignIn30Day.LastGetTime) {
			raise_error(msg.ErrInvalidActivityTime)
		}

		actor.user.OperationActivity.SignIn30Day.Process += 1
		id, reward, diamond_add, gold_add := activity_get_sigin_month(actor.user.OperationActivity.SignIn30Day.Process)
		activity_id = id

		award := cell_plains_to_cells(reward)
		actor.put_to_bag(award)
		actor.incr_diamond(diamond_add)
		actor.incr_gold(gold_add)

		bag_diff = actor.diff_bag()
		process_info = actor.user_info().OperationActivity.SignIn30Day
		diamond = actor.user_info().Diamond
		gold = actor.chara_info().Gold
		actor.user.OperationActivity.SignIn30Day.LastGetTime = time.Now().Unix()
		actor.save()
	})

	if err != msg.Success {
		return &msg.GetActivtyAwardSignIn30DayResponse{Err: err}
	}

	return &msg.GetActivtyAwardSignIn30DayResponse{
		ActivityId:  activity_id,
		ProcessInfo: process_info,
		BagDiff:     bag_diff,
		Diamond:     diamond,
		Gold:        gold,
	}
}
