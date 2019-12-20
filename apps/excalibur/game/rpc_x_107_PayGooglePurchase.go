package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/payment"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/lib/log"
)

func (g *GameService) PayGooglePurchase(req *msg.Request, arg *msg.PayGooglePurchaseRequest) *msg.PayGooglePurchaseResponse {

	err, info := payment.UnmarshalGooglePurchaseData(arg.PurchaseInfo)
	log.Info(err, info)
	// err = payment.GoogleValidateToken(info.PackageName, info.ProductId, info.PurchaseToken)
	log.Info(err)
	return &msg.PayGooglePurchaseResponse{}
}
