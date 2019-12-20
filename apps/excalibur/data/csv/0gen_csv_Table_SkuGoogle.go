package csv

type SkuGoogle struct {
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

var Table_SkuGoogle = map[int]SkuGoogle{}

func Load_Table_SkuGoogle() {
	Table_SkuGoogle = map[int]SkuGoogle{
		1: SkuGoogle{
			Uid:        1,
			SkuId:      "com.sdbean.excalibur.6yuan",
			SkuType:    "diamond",
			SkuName:    "小袋钻石",
			DiamondNum: 60,
			RMB:        6,
			USD:        0.99,
			Icon:       "zuanshi_1",
			AltasPath:  "UI/EquAtlas/EquAtlas_ref",
		},
		2: SkuGoogle{
			Uid:        2,
			SkuId:      "com.sdbean.excalibur.18yuan",
			SkuType:    "diamond",
			SkuName:    "中袋钻石",
			DiamondNum: 180,
			RMB:        18,
			USD:        2.99,
			Icon:       "zuanshi_2",
			AltasPath:  "UI/EquAtlas/EquAtlas_ref",
		},
		3: SkuGoogle{
			Uid:        3,
			SkuId:      "com.sdbean.excalibur.30yuan",
			SkuType:    "diamond",
			SkuName:    "大袋钻石",
			DiamondNum: 300,
			RMB:        30,
			USD:        4.99,
			Icon:       "zuanshi_3",
			AltasPath:  "UI/EquAtlas/EquAtlas_ref",
		},
		4: SkuGoogle{
			Uid:        4,
			SkuId:      "com.sdbean.excalibur.60yuan",
			SkuType:    "diamond",
			SkuName:    "小堆钻石",
			DiamondNum: 630,
			RMB:        60,
			USD:        9.99,
			Icon:       "zuanshi_5",
			AltasPath:  "UI/EquAtlas/EquAtlas_ref",
		},
		5: SkuGoogle{
			Uid:        5,
			SkuId:      "com.sdbean.excalibur.128yuan",
			SkuType:    "diamond",
			SkuName:    "中堆钻石",
			DiamondNum: 1350,
			RMB:        128,
			USD:        19.9,
			Icon:       "zuanshi_6",
			AltasPath:  "UI/EquAtlas/EquAtlas_ref",
		},
		6: SkuGoogle{
			Uid:        6,
			SkuId:      "com.sdbean.excalibur.328yuan",
			SkuType:    "diamond",
			SkuName:    "大堆钻石",
			DiamondNum: 3450,
			RMB:        328,
			USD:        49.9,
			Icon:       "zuanshi_7",
			AltasPath:  "UI/EquAtlas/EquAtlas_ref",
		},
		7: SkuGoogle{
			Uid:        7,
			SkuId:      "com.sdbean.excalibur.vip.month",
			SkuType:    "vip.month",
			SkuName:    "月卡",
			DiamondNum: 300,
			RMB:        30,
			USD:        4.99,
			Icon:       "",
			AltasPath:  "",
		},
		8: SkuGoogle{
			Uid:        8,
			SkuId:      "com.sdbean.excalibur.vip.godbless",
			SkuType:    "vip.godbless",
			SkuName:    "神佑",
			DiamondNum: 980,
			RMB:        96,
			USD:        13.99,
			Icon:       "",
			AltasPath:  "",
		},
	}
}
