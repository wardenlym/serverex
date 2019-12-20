package game

import (
	"time"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) VipYueKaMaterialGetAward(req *msg.Request, arg *msg.VipYueKaMaterialGetAwardRequest) *msg.VipYueKaMaterialGetAwardResponse {

	var diamond uint
	var bag_diff []odm.LocatedCell

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)

		if actor.user.VipInfo.Vip_YueKa_Material_ExpireTime < time.Now().Unix() {
			raise_error(msg.ErrVipYueKaGetAwardExpire)
		}

		if !is_flushed_time_valid_everyday(actor.user.VipInfo.Vip_YueKa_Material_LastGetTime) {
			raise_error(msg.ErrVipYueKaGetAwardFailure)
		}

		actor.incr_diamond(consts.VipYueKaMaterialDiamond)
		actor.user.VipInfo.Vip_YueKa_Material_LastGetTime = time.Now().Unix()
		diamond = actor.user.Diamond

		actor.put_to_bag([]odm.Cell{cell_new_with_typeid_1(30013)})
		bag_diff = actor.diff_bag()
		actor.save()
	})

	if err != msg.Success {
		return &msg.VipYueKaMaterialGetAwardResponse{Err: err}
	}

	return &msg.VipYueKaMaterialGetAwardResponse{
		Diamond: diamond,
		BagDiff: bag_diff,
	}
}
