package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

const (
	ActivityFirstPay_Process_1 = 1 // 0 未充值，1已充值未领取，2已经领取
	ActivityFirstPay_Process_2 = 2
)

func (g *GameService) GetActivtyAwardFisrtPay(req *msg.Request, arg *msg.GetActivtyAwardFisrtPayRequest) *msg.GetActivtyAwardFisrtPayResponse {
	var activity_id int
	var process_info odm.ActivityProcessInfo
	bag_diff := []odm.LocatedCell{}
	var diamond uint
	var gold uint

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)
		if actor.user_info().OperationActivity.FisrtPay.Process != ActivityFirstPay_Process_1 {
			raise_error(msg.ErrInvalidActivityProcess)
		}

		id, reward, diamond_add, gold_add := activity_get_firstpay()
		activity_id = id

		award := cell_plains_to_cells(reward)
		actor.put_to_bag(award)
		actor.incr_diamond(diamond_add)
		actor.incr_gold(gold_add)

		bag_diff = actor.diff_bag()
		process_info = actor.user_info().OperationActivity.FisrtPay
		diamond = actor.user_info().Diamond
		gold = actor.chara_info().Gold
		actor.user.OperationActivity.FisrtPay.Process = ActivityFirstPay_Process_2
		actor.save()
	})

	if err != msg.Success {
		return &msg.GetActivtyAwardFisrtPayResponse{Err: err}
	}

	return &msg.GetActivtyAwardFisrtPayResponse{
		ActivityId:  activity_id,
		ProcessInfo: process_info,
		BagDiff:     bag_diff,
		Diamond:     diamond,
		Gold:        gold,
	}
}
