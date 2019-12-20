package game

import "time"

func (a *Actor) is_vip_godbless() bool {
	return a.user.VipInfo.IsVip_Godbless
}

func (a *Actor) enable_vip_godbless() {
	a.user.VipInfo.IsVip_Godbless = true
}

func (a *Actor) is_vip_yueka_material() bool {
	return a.user.VipInfo.Vip_YueKa_Material_ExpireTime > time.Now().Unix()
}

func (a *Actor) enable_vip_yueka_material(day int) {
	if a.user.VipInfo.Vip_YueKa_Material_ExpireTime < time.Now().Unix() {
		a.user.VipInfo.Vip_YueKa_Material_ExpireTime =
			time.Now().
				AddDate(0, 0, consts.VipYueKaMaterialExpireDays).
				Unix()
	} else {
		a.user.VipInfo.Vip_YueKa_Material_ExpireTime =
			time.
				Unix(a.user.VipInfo.Vip_YueKa_Material_ExpireTime, 0).
				AddDate(0, 0, consts.VipYueKaMaterialExpireDays).
				Unix()
	}
}
