package game

import (
	"math"
	"time"

	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/util"
)

func (a *Actor) if_explore_info_need_refresh() *Actor {

	if is_flushed_time_valid_everyday(a.user_info().ExploreInfo.LastRefreshChargeTime) {
		a.user_info().ExploreInfo.ChargeAvailable = 5
		a.user_info().ExploreInfo.ChargeTotal = 5
		a.user_info().ExploreInfo.LastRefreshChargeTime = time.Now().Unix()
	}

	return a
}

func (a *Actor) add_explore_charge_count(count int, diamond uint) *Actor {
	if a.user_info().Diamond < diamond {
		raise_error(msg.ErrDiamondNotEnough)
	}
	a.user_info().ExploreInfo.ChargeAvailable += count
	return a
}

func find_explore_index(info []*odm.ExploreStatus, f func(*odm.ExploreStatus) bool) (int, bool) {

	for i, v := range info {
		if f(v) {
			return i, true
		}
	}
	return -1, false
}

func (a *Actor) reduce_explore_time(chara_id string, stage_id int, sec int64) *Actor {

	index, exist := find_explore_index(
		a.user_info().ExploreInfo.Status,
		func(v *odm.ExploreStatus) bool {
			if v.CharacterId == chara_id && v.ExploreStageId == stage_id {
				return true
			}
			return false
		})

	if !exist {
		raise_error(msg.ErrInvalidExploreId)
	}
	a.user_info().ExploreInfo.Status[index].CompleteAt =
		time.Unix(a.user_info().ExploreInfo.Status[index].CompleteAt, 0).Add(time.Duration(-1*sec) * time.Second).Unix()
	return a
}

func (a *Actor) explore_start(chara_id string, stage_id int) *Actor {

	if a.user_info().ExploreInfo.ChargeAvailable < 1 {
		raise_error(msg.ErrAvalibleCountNotEnough)
	}

	index, exist := find_explore_index(
		a.user_info().ExploreInfo.Status,
		func(v *odm.ExploreStatus) bool {
			if v.CharacterId == chara_id && v.IsExploring == false {
				return true
			}
			return false
		})

	if !exist {
		raise_error(msg.ErrInvalidExploreId)
	}

	a.user_info().ExploreInfo.ChargeAvailable -= 1

	a.user_info().ExploreInfo.Status[index].ExploreStageId = stage_id
	a.user_info().ExploreInfo.Status[index].DispatchedAt = time.Now().Unix()
	a.user_info().ExploreInfo.Status[index].CompleteAt = time.Now().Add(calc_explore_time(stage_id)).Unix()
	a.user_info().ExploreInfo.Status[index].IsExploring = true
	return a
}

func (a *Actor) explore_end(chara_id string, stage_id int) (int, int, int64) {

	index, exist := find_explore_index(
		a.user_info().ExploreInfo.Status,
		func(v *odm.ExploreStatus) bool {
			if v.CharacterId == chara_id && v.IsExploring == true {
				return true
			}
			return false
		})

	if !exist {
		raise_error(msg.ErrInvalidExploreId)
	}

	use_time := time.Now().Unix() - a.user_info().ExploreInfo.Status[index].DispatchedAt

	if _, exist := a.user_info().Characters[chara_id]; !exist {
		raise_error(msg.ErrInvalidCharacterId)
	}

	hero_id := a.user_info().Characters[chara_id].HeroId
	hero_level := a.user_info().Characters[chara_id].Level

	a.user_info().ExploreInfo.Status[index].ExploreStageId = 0
	a.user_info().ExploreInfo.Status[index].DispatchedAt = 0
	a.user_info().ExploreInfo.Status[index].CompleteAt = 0
	a.user_info().ExploreInfo.Status[index].IsExploring = false
	return hero_id, hero_level, use_time
}

func calc_explore_time(stage_id int) time.Duration {
	info, exist := csv.Table_Adventure[stage_id]
	if !exist {
		raise_error(msg.ErrInvalidExploreId)
	}
	// return time.Minute * 5
	return time.Minute * time.Duration(info.Adventure_time)
}

func hero_info_by_id(heroid int) (csv.HeroData, bool) {
	info, exist := csv.Table_HeroData[heroid]
	return info, exist
}

func calc_explore_award(stage_id int, hero_id int, hero_level int, use_time_sec int64) []odm.PlainCell {
	info, exist := csv.Table_Adventure[stage_id]
	if !exist {
		raise_error(msg.ErrInvalidExploreId)
	}

	ret := []odm.PlainCell{}
	time_need_sec := int64(info.Adventure_time * 60)

	is_competed, complet_rate := func() (bool, float64) {
		if use_time_sec >= time_need_sec {
			return true, 1
		}
		return false, float64(use_time_sec) / float64(time_need_sec)
	}()

	special_add_rate, special_lucky := func() (float64, float64) {
		info, exist := csv.Table_HeroData[hero_id]
		if !exist {
			return 0, 0
		}
		return float64(info.SpecialtyAddRate), float64(info.Specialtylucky)
	}()

	{
		// 固定收益
		get_num := func(total int, rate float64) uint {
			k := (float64(total) * rate) * (special_add_rate + float64(hero_level-1)*0.01)
			return uint(math.Floor(k))
		}

		for k, v := range info.Adventure_get {
			ret = append(ret, odm.PlainCell{
				TypeId: v,
				Num:    get_num(info.Get_num[k], complet_rate),
			})
		}
	}

	// 额外收益
	if is_competed {
		get_num := func(count int) uint {
			k := float64(count) * (special_add_rate + float64(hero_level-1)*0.01)
			return uint(math.Floor(k))
		}

		f := func(typeid int, count []int) {
			if len(count) < 2 {
				return
			}
			ret = append(ret, odm.PlainCell{
				TypeId: typeid,
				Num:    get_num(util.RandomIn(count[0], count[1])),
			})
		}

		f(info.Extra_1, info.Extra_1_num)
		f(info.Extra_2, info.Extra_2_num)
		f(info.Extra_3, info.Extra_3_num)
		f(info.Extra_4, info.Extra_4_num)
	}

	// 幸运道具掉落
	if is_competed {
		rate := special_lucky + float64(info.Item_probability)

		if util.HitRate(rate) {
			ret = append(ret, odm.PlainCell{
				TypeId: info.Item,
				Num:    1,
			})
		}
	}

	return ret
}
