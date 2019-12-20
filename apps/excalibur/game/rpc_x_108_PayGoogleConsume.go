package game

import (
	"errors"

	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/lib/log"
)

func (g *GameService) PayGoogleConsume(req *msg.Request, arg *msg.PayGoogleConsumeRequest) *msg.PayGoogleConsumeResponse {

	err := try_do(func() {
		actor := g.use_actor(req.Uid)
		// err := payment.GoogleValidateToken(arg.PackageName, arg.ProductId, arg.PurchaseToken)
		// log.Info(err)

		var total_cnt uint

		func() {
			v := arg
			log.Info(v.ProductId)
			sku_type := google_product_sku_type(v.ProductId)
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

			diamond, err := google_pay_product_dimond_info(v.ProductId)
			if err != nil {
				log.Error(err)
				raise_error(msg.ErrInvalidProductId)
				return
			}
			total_cnt += diamond
		}()

		actor.incr_diamond(total_cnt)
		// diamond = actor.user_info().Diamond

		actor.save()
	})

	if err != msg.Success {
		return &msg.PayGoogleConsumeResponse{Err: err}
	}
	return &msg.PayGoogleConsumeResponse{
		// SkuType: sku_type,
		// Diamond: diamond,
	}
}

func google_pay_product_dimond_info(product_id string) (uint, error) {

	for _, v := range csv.Table_SkuGoogle {
		if product_id == v.SkuId {
			return uint(v.DiamondNum), nil
		}
	}
	return 0, errors.New("product id error: " + product_id)
}

func google_product_sku_type(product_id string) string {

	for _, v := range csv.Table_SkuGoogle {
		if product_id == v.SkuId {
			return v.SkuType
		}
	}
	return "diamond"
}
