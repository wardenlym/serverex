package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) ListDiamondPrice(req *msg.Request, arg *msg.ListDiamondPriceRequest) *msg.ListDiamondPriceResponse {

	list := []odm.DiamondPrice{}

	for _, v := range csv.Table_RechargeData {
		list = append(list, odm.DiamondPrice{
			GoodsId:  v.GoodsId,
			Num:      uint(v.Num),
			PriceRMB: uint(v.PricerRMB),
			Discount: uint(v.Discount)})
	}

	return &msg.ListDiamondPriceResponse{
		DiamondPrices: list,
	}
}
