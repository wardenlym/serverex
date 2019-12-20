package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/log"
)

func activity_get_firstpay() (int, []odm.PlainCell, uint, uint) {
	info, exist := csv.Table_Rewards_FirstPay[1]
	if exist != true {
		return 0, []odm.PlainCell{}, 0, 0
	}
	reward := []odm.PlainCell{}

	for _, id := range info.ItemRewards {
		reward = append(reward, odm.PlainCell{TypeId: id, Num: 1})
	}

	return info.ID, reward, uint(info.DiamondsRewards), uint(info.GoldRewards)
}

func activity_get_sigin_first_week_end() int {
	return len(csv.Table_Rewards_Week)
}

func activity_get_sigin_first_week(day int) (int, []odm.PlainCell, uint, uint) {
	info, exist := csv.Table_Rewards_Week[day]
	if exist != true {
		return 0, []odm.PlainCell{}, 0, 0
	}

	reward := []odm.PlainCell{}
	log.Warn(reward)
	for _, id := range info.ItemRewards {
		reward = append(reward, odm.PlainCell{TypeId: id, Num: 1})
	}
	log.Warn(reward)
	reward = cell_plains_fold(reward)
	log.Warn(reward)
	return info.ID, reward, uint(info.DiamondsRewards), uint(info.GoldRewards)
}

func activity_get_sigin_month_end() int {
	return len(csv.Table_Rewards_Month)
}

func activity_get_sigin_month(day int) (int, []odm.PlainCell, uint, uint) {
	info, exist := csv.Table_Rewards_Month[day]
	if exist != true {
		return 0, []odm.PlainCell{}, 0, 0
	}

	reward := []odm.PlainCell{}
	log.Warn(reward)
	for _, id := range info.ItemRewards {
		reward = append(reward, odm.PlainCell{TypeId: id, Num: 1})
	}
	log.Warn(reward)
	reward = cell_plains_fold(reward)
	log.Warn(reward)
	return info.ID, reward, uint(info.DiamondsRewards), uint(info.GoldRewards)
}
