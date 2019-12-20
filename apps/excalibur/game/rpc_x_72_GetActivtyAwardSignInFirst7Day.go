package game

import (
	"time"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/log"
)

const (
	ActivitySignInFirst7Day_Process_End = 7 // 7 全部已经领取

	ActivitySignInFlushTime = 6
)

func is_flushed_time_valid_everyday(last_timestamp int64) bool {
	if last_timestamp == 0 {
		return true
	}

	last_time := time.Unix(last_timestamp, 0)
	delta_time := time.Since(last_time)

	// last_time ~ now > 24hour
	if delta_time > time.Hour*24 {
		return true
	} else {
		now := time.Now()
		today_flush_point_at := time.Date(now.Year(), now.Month(), now.Day(),
			ActivitySignInFlushTime, // 6:00AM+8:00
			0, 0, 0, now.Location())

		// last_time < today_flush_point < now  -> ok
		if now.After(last_time) &&
			now.After(today_flush_point_at) &&
			last_time.Before(today_flush_point_at) {
			return true
		}
	}
	return false
}

func (g *GameService) GetActivtyAwardSignInFirst7Day(req *msg.Request, arg *msg.GetActivtyAwardSignInFirst7DayRequest) *msg.GetActivtyAwardSignInFirst7DayResponse {

	var activity_id int
	var process_info odm.ActivityProcessInfo
	bag_diff := []odm.LocatedCell{}
	var diamond uint
	var gold uint

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)
		if actor.user_info().OperationActivity.SignInFirst7Day.Process >= ActivitySignInFirst7Day_Process_End {
			raise_error(msg.ErrInvalidActivityProcess)
		}

		if !is_flushed_time_valid_everyday(actor.user_info().OperationActivity.SignInFirst7Day.LastGetTime) {
			raise_error(msg.ErrInvalidActivityTime)
		}

		actor.user.OperationActivity.SignInFirst7Day.Process += 1
		id, reward, diamond_add, gold_add := activity_get_sigin_first_week(actor.user.OperationActivity.SignInFirst7Day.Process)
		activity_id = id

		log.Warn(id, diamond_add, gold_add, reward)
		award := cell_plains_to_cells(reward)
		log.Warn(award)
		actor.put_to_bag(award)
		actor.incr_diamond(diamond_add)
		actor.incr_gold(gold_add)

		bag_diff = actor.diff_bag()
		process_info = actor.user_info().OperationActivity.SignInFirst7Day
		diamond = actor.user_info().Diamond
		gold = actor.chara_info().Gold
		actor.user.OperationActivity.SignInFirst7Day.LastGetTime = time.Now().Unix()
		actor.save()
	})

	if err != msg.Success {
		return &msg.GetActivtyAwardSignInFirst7DayResponse{Err: err}
	}

	return &msg.GetActivtyAwardSignInFirst7DayResponse{
		ActivityId:  activity_id,
		ProcessInfo: process_info,
		BagDiff:     bag_diff,
		Diamond:     diamond,
		Gold:        gold,
	}
}
