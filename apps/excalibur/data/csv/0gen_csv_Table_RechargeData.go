package csv

type RechargeData struct {
	// 商品ID
	GoodsId int

	// 商品名称
	GoodsName string

	// 钻石数量
	Num int

	// 人民币价格
	PricerRMB int

	// 折扣
	Discount int

	// 图标名称
	Icon string

	// 图集路径
	AltasPath string

	// 产品id
	ProductId string
}

var Table_RechargeData = map[int]RechargeData{}

func Load_Table_RechargeData() {
	Table_RechargeData = map[int]RechargeData{
		1: RechargeData{
			GoodsId:   1,
			GoodsName: "小袋钻石",
			Num:       60,
			PricerRMB: 6,
			Discount:  100,
			Icon:      "zuanshi_1",
			AltasPath: "UI/EquAtlas/EquAtlas_ref",
			ProductId: "excalibur.sdbean.6rmb",
		},
		2: RechargeData{
			GoodsId:   2,
			GoodsName: "中袋钻石",
			Num:       180,
			PricerRMB: 18,
			Discount:  100,
			Icon:      "zuanshi_2",
			AltasPath: "UI/EquAtlas/EquAtlas_ref",
			ProductId: "excalibur.sdbean.18rmb",
		},
		3: RechargeData{
			GoodsId:   3,
			GoodsName: "大袋钻石",
			Num:       300,
			PricerRMB: 30,
			Discount:  100,
			Icon:      "zuanshi_3",
			AltasPath: "UI/EquAtlas/EquAtlas_ref",
			ProductId: "excalibur.sdbean.30rmb",
		},
		4: RechargeData{
			GoodsId:   4,
			GoodsName: "小堆钻石",
			Num:       630,
			PricerRMB: 60,
			Discount:  100,
			Icon:      "zuanshi_5",
			AltasPath: "UI/EquAtlas/EquAtlas_ref",
			ProductId: "excalibur.sdbean.60rmb",
		},
		5: RechargeData{
			GoodsId:   5,
			GoodsName: "中堆钻石",
			Num:       1350,
			PricerRMB: 128,
			Discount:  100,
			Icon:      "zuanshi_6",
			AltasPath: "UI/EquAtlas/EquAtlas_ref",
			ProductId: "excalibur.sdbean.128rmb",
		},
		6: RechargeData{
			GoodsId:   6,
			GoodsName: "大堆钻石",
			Num:       3450,
			PricerRMB: 328,
			Discount:  100,
			Icon:      "zuanshi_7",
			AltasPath: "UI/EquAtlas/EquAtlas_ref",
			ProductId: "excalibur.sdbean.328rmb",
		},
	}
}
