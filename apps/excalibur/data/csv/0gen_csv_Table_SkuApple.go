package csv

type SkuApple struct {
	// uid
	Uid int

	// 商品id
	SkuId string

	// 商量类型
	SkuType string

	// 商品名称
	SkuName string

	// 钻石数量
	DiamondNum int

	// 人民币价格
	RMB float32

	// 美元价格
	USD float32

	// 图标名称
	Icon string

	// 图集路径
	AltasPath string
}

var Table_SkuApple = map[int]SkuApple{}

func Load_Table_SkuApple() {
	Table_SkuApple = map[int]SkuApple{
		1: SkuApple{
			Uid:        1,
			SkuId:      "excalibur.sdbean.6rmb",
			SkuType:    "diamond",
			SkuName:    "小袋钻石",
			DiamondNum: 60,
			RMB:        6,
			USD:        0.99,
			Icon:       "zuanshi_1",
			AltasPath:  "UI/EquAtlas/EquAtlas_ref",
		},
		2: SkuApple{
			Uid:        2,
			SkuId:      "excalibur.sdbean.18rmb",
			SkuType:    "diamond",
			SkuName:    "中袋钻石",
			DiamondNum: 180,
			RMB:        18,
			USD:        2.99,
			Icon:       "zuanshi_2",
			AltasPath:  "UI/EquAtlas/EquAtlas_ref",
		},
		3: SkuApple{
			Uid:        3,
			SkuId:      "excalibur.sdbean.30rmb",
			SkuType:    "diamond",
			SkuName:    "大袋钻石",
			DiamondNum: 300,
			RMB:        30,
			USD:        4.99,
			Icon:       "zuanshi_3",
			AltasPath:  "UI/EquAtlas/EquAtlas_ref",
		},
		4: SkuApple{
			Uid:        4,
			SkuId:      "excalibur.sdbean.60rmb",
			SkuType:    "diamond",
			SkuName:    "小堆钻石",
			DiamondNum: 630,
			RMB:        60,
			USD:        9.99,
			Icon:       "zuanshi_5",
			AltasPath:  "UI/EquAtlas/EquAtlas_ref",
		},
		5: SkuApple{
			Uid:        5,
			SkuId:      "excalibur.sdbean.128rmb",
			SkuType:    "diamond",
			SkuName:    "中堆钻石",
			DiamondNum: 1350,
			RMB:        128,
			USD:        19.9,
			Icon:       "zuanshi_6",
			AltasPath:  "UI/EquAtlas/EquAtlas_ref",
		},
		6: SkuApple{
			Uid:        6,
			SkuId:      "excalibur.sdbean.328rmb",
			SkuType:    "diamond",
			SkuName:    "大堆钻石",
			DiamondNum: 3450,
			RMB:        328,
			USD:        49.9,
			Icon:       "zuanshi_7",
			AltasPath:  "UI/EquAtlas/EquAtlas_ref",
		},
		7: SkuApple{
			Uid:        7,
			SkuId:      "excalibur.sdbean.yueka",
			SkuType:    "vip.month",
			SkuName:    "月卡",
			DiamondNum: 300,
			RMB:        30,
			USD:        4.99,
			Icon:       "",
			AltasPath:  "",
		},
		8: SkuApple{
			Uid:        8,
			SkuId:      "excalibur.sdbean.shenyou",
			SkuType:    "vip.godbless",
			SkuName:    "神佑",
			DiamondNum: 980,
			RMB:        98,
			USD:        13.99,
			Icon:       "",
			AltasPath:  "",
		},
	}
}
