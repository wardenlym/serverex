package gameplay

import (
	"time"

	"github.com/json-iterator/go"
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/odm"
	"gitlab.com/megatech/serverex/lib/util"
)

func makeOnSaleInfoFromTable() []odm.OnSaleInfo {
	a := []odm.OnSaleInfo{}

	for _, v := range csv.Table_MarketData {
		a = append(a, odm.OnSaleInfo{
			GoodsId:      v.Goods_id,
			TypeId:       v.Id,
			NumTotal:     uint(v.Inventory),
			NumAvailable: uint(v.Inventory),
			NumSQ:        uint(v.Sq_number_value),

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
		})
	}

	return a
}

func (g *game) listOnSale(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.ListOnSaleRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}

	info, err := g.db.GetUser(req.Uid)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	onsales := info.OnSaleInfos
	lastUpdateTime := info.OnSaleLastUpdateTime

	now := time.Now()

	if lastUpdateTime == 0 || time.Unix(lastUpdateTime, 0).Add(10*time.Minute).Unix() <= now.Unix() { // for test
		sampleOnsale := makeOnSaleInfoFromTable()

		if len(info.OnSaleInfos) == 0 {
			g.db.UpsertUserField(req.Uid, "onSaleInfos", sampleOnsale)
			onsales = sampleOnsale
			lastUpdateTime = time.Now().Unix()
			g.db.UpsertUserField(req.Uid, "onSaleInfos", sampleOnsale)
			g.db.UpsertUserField(req.Uid, "onSaleLastUpdateTime", time.Now().Unix())
		}
	}

	res.Data = &msg.ListOnSaleResponse{
		LastTime:        lastUpdateTime,
		CurrentTime:     now.Unix(),
		NextRefreshTime: time.Unix(lastUpdateTime, 0).Add(10 * time.Minute).Unix(),
		OnSaleInfos:     onsales,
	}
	return res
}

func (g *game) refreshOnSale(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.RefreshOnSaleRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}
	g.db.UpsertUserField(req.Uid, "onSaleInfos", makeOnSaleInfoFromTable())
	res.Data = &msg.RefreshOnSaleResponse{IsSuccess: true}
	return res
}

func (g *game) purchase(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.PurchaseRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}

	info, err := g.db.GetUser(req.Uid)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	// 暂时没有检测消耗金币的正确数字
	if info.Diamond < o.CostDiamond {
		res.Data = &msg.PurchaseResponse{ErrMsg: "钻石不足"}
		return res
	}
	if info.Gold < o.CostGold {
		res.Data = &msg.PurchaseResponse{ErrMsg: "金币不足"}
		return res
	}

	gold := info.Gold - o.CostGold
	diamond := info.Diamond - o.CostDiamond

	onsales := info.OnSaleInfos
	ok, index, errmsg := func() (bool, int, string) {
		for i, goods := range onsales {
			if o.GoodsId == goods.GoodsId && o.TypeId == goods.TypeId {
				if goods.NumAvailable >= o.Num {
					return true, i, ""
				} else {
					return false, -1, "库存数量不足"
				}
			}
		}
		return false, -1, "商品ID错误"
	}()

	if !ok {
		res.Data = &msg.PurchaseResponse{ErrMsg: errmsg}
		return res
	}

	sq_num := onsales[index].NumSQ
	onsales[index].NumAvailable = onsales[index].NumAvailable - o.Num

	cells := info.Characters[o.CharacterId].Bag.Cells
	ok, index = func() (bool, int) {
		for i, c := range cells {
			if c.Item == nil || c.Num == 0 {
				return true, i
			}
		}
		return false, -1
	}()
	if !ok {
		res.Data = &msg.PurchaseResponse{ErrMsg: "背包空间不足"}
		return res
	}

	cells[index] = odm.Cell{
		Index: index,
		Item: &odm.Item{
			Guid:            util.NewSnowflakeID(),
			TypeId:          o.TypeId,
			StarLevel:       0,
			StarUpgradeInfo: []int{},
		},
		Num: sq_num,
	}

	// 更新数据稍后做成一次性update的
	g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".bag.cells", cells)
	g.db.UpsertUserField(req.Uid, "gold", gold)
	g.db.UpsertUserField(req.Uid, "diamond", diamond)
	g.db.UpsertUserField(req.Uid, "onSaleInfos", onsales)
	bag := info.Characters[o.CharacterId].Bag
	bag.Cells = cells
	res.Data = &msg.PurchaseResponse{
		IsSuccess: true,
		Gold:      gold,
		Diamond:   diamond,
		BagStatus: bag,
	}
	return res
}

//
func (g *game) listDiamondPrice(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.ListDiamondPriceRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}

	list := []odm.DiamondPrice{}

	for _, v := range csv.Table_RechargeData {
		list = append(list, odm.DiamondPrice{
			GoodsId:  v.GoodsId,
			Num:      uint(v.Num),
			PriceRMB: uint(v.PricerRMB),
			Discount: uint(v.Discount)})
	}

	res.Data = &msg.ListDiamondPriceResponse{
		DiamondPrices: list,
	}
	return res
}

func (g *game) submitOrder(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.SubmitOrderRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}
	res.Data = &msg.SubmitOrderResponse{}
	return res
}

func (g *game) queryOrderStatus(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.QueryOrderStatusRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}
	res.Data = &msg.QueryOrderStatusResponse{
		OrderInfo: odm.OrderInfo{},
	}
	return res
}
