package game

import (
	"time"

	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/util"
)

func (a *Actor) list_onsale(now time.Time) {
	onsales := a.user.OnSaleInfos
	lastUpdateTime := a.user.OnSaleLastUpdateTime

	if lastUpdateTime == 0 || len(onsales) == 0 || // 第一次
		time.Unix(lastUpdateTime, 0).Add(consts.OnSalesRefreshInterval).Unix() <= now.Unix() { // for test

		if a.user.RankingInfo.BestStageId == 0 {
			a.user.RankingInfo.BestStageId = consts.FirstStageId
		}
		sampleOnsale := makeOnSaleInfoFromTable(a.user.RankingInfo.BestStageId)

		onsales = sampleOnsale
		lastUpdateTime = now.Unix()

		a.user.OnSaleInfos = onsales
		a.user.OnSaleLastUpdateTime = lastUpdateTime
	}
}

func (g *GameService) ListOnSale(req *msg.Request, arg *msg.ListOnSaleRequest) *msg.ListOnSaleResponse {

	var last_update_time int64
	var onsale_infos []odm.OnSaleInfo

	now := time.Now()

	err := try_do(func() {
		actor := g.use_actor(req.Uid)
		actor.list_onsale(now)
		last_update_time = actor.user.OnSaleLastUpdateTime
		onsale_infos = actor.user.OnSaleInfos
		actor.save()
	})

	if err != msg.Success {
		return &msg.ListOnSaleResponse{Err: err}
	}
	return &msg.ListOnSaleResponse{
		LastTime:        last_update_time,
		CurrentTime:     now.Unix(),
		NextRefreshTime: time.Unix(last_update_time, 0).Add(consts.OnSalesRefreshInterval).Unix(),
		OnSaleInfos:     onsale_infos,
	}
}

func makeOnSaleInfoFromTable(stageid int) []odm.OnSaleInfo {
	a := []odm.OnSaleInfo{}

	to_pick := []csv.MarketData{}

	for _, v := range csv.Table_MarketData {
		if stageid >= v.AvailableStageId {
			to_pick = append(to_pick, v)
		}
	}

	indices := util.RandomIndexList(len(to_pick), consts.OnSalesDisplayCount)

	for _, index := range indices {
		v := to_pick[index]
		a = append(a, new_onSaleInfo_by_MarketData(v))
	}

	return a
}

func new_onSaleInfo_by_MarketData(v csv.MarketData) odm.OnSaleInfo {
	return odm.OnSaleInfo{
		Guid:             util.NewSnowflakeID(),
		GoodsId:          v.Goods_id,
		TypeId:           v.Id,
		NumTotal:         uint(v.Inventory),
		NumAvailable:     uint(v.Inventory),
		NumSQ:            uint(v.Sq_number_value),
		AvailableStageId: v.Dungeon_number,
		CostGold: uint(func() int {
			if v.Buy_type == 1 {
				return v.Buy_price
			} else {
				return 0
			}
		}()),
		CostDiamond: uint(func() int {
			if v.Buy_type == 2 {
				return v.Buy_price
			} else {
				return 0
			}
		}()),
		Discount: uint(v.Discount),
	}
}
