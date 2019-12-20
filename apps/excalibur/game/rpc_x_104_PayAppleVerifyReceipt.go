package game

import (
	"errors"

	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/payment"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/lib/log"
)

const is_vip_yueka = "vip.month"
const is_vip_godbless = "vip.godbless"

func (g *GameService) PayAppleVerifyReceipt(req *msg.Request, arg *msg.PayAppleVerifyReceiptRequest) *msg.PayAppleVerifyReceiptResponse {

	var sku_type string
	var diamond uint

	err := try_do(func() {
		actor := g.use_actor(req.Uid)

		if g.config.GateHttpOption.DevModeEnable == false {
			arg.UseSandbox = false
		}
		receipt, err := payment.VerifyReceipt(arg.Receipt, arg.UseSandbox)
		if err != nil {
			log.Error(err, receipt)
			receipt, err = payment.VerifyReceipt(arg.Receipt, !arg.UseSandbox)
			if err != nil {
				log.Error(err, receipt)
				raise_error(msg.ErrVerifyReceiptFailure)
			}
		}

		if g.db.Order().IsAppleTranscationExist(receipt.TransactionID) {
			log.Error("receipted: ", receipt)
			raise_error(msg.ErrTransactionIdReceipted)
		}
		var total_cnt uint

		func(v *payment.Receipt) {
			g.db.Order().InsertTempAppleOrderLog(req.Uid, v.ProductID, arg.UseSandbox, v.TransactionID, v)
			log.Info(v.TransactionID)
			log.Info(v.ProductID)
			sku_type = apple_product_sku_type(v.ProductID)
			if sku_type == is_vip_godbless {
				log.Info("处理神佑")
				actor.enable_vip_godbless()
				return
			}
			if sku_type == is_vip_yueka {
				log.Info("处理钻石月卡")
				actor.enable_vip_yueka_material(consts.VipYueKaMaterialExpireDays)
				return
			}

			diamond, err := apple_pay_product_dimond_info(v.ProductID)
			if err != nil {
				log.Error(err)
				raise_error(msg.ErrInvalidProductId)
				return
			}
			total_cnt += diamond
		}(receipt)

		actor.incr_diamond(total_cnt)
		diamond = actor.user_info().Diamond
		actor.save()

		func(v *payment.Receipt) {
			g.db.Order().InsertTempAppleOrderLog(req.Uid, v.ProductID, arg.UseSandbox, v.TransactionID, v)
		}(receipt)

	})
	if err != msg.Success {
		return &msg.PayAppleVerifyReceiptResponse{Err: err}
	}
	return &msg.PayAppleVerifyReceiptResponse{
		SkuType: sku_type,
		Diamond: diamond,
	}
}

func apple_product_sku_type(product_id string) string {

	for _, v := range csv.Table_SkuApple {
		if product_id == v.SkuId {
			return v.SkuType
		}
	}
	return "diamond"
}

func apple_pay_product_dimond_info(product_id string) (uint, error) {

	for _, v := range csv.Table_SkuApple {
		if product_id == v.SkuId {
			return uint(v.DiamondNum), nil
		}
	}
	return 0, errors.New("product id error: " + product_id)
}
