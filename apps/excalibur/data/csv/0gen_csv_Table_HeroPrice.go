package csv

type HeroPrice struct {
	// 英雄ID
	Id int

	// 英雄名称
	Name string

	// 钻石价格
	Diamondprice int

	// 金币价格
	Goldprice int
}

var Table_HeroPrice = map[int]HeroPrice{}

func Load_Table_HeroPrice() {
	Table_HeroPrice = map[int]HeroPrice{
		20000: HeroPrice{
			Id:           20000,
			Name:         "精灵法师",
			Diamondprice: 0,
			Goldprice:    120000,
		},
		30000: HeroPrice{
			Id:           30000,
			Name:         "亡灵勇士",
			Diamondprice: 1024,
			Goldprice:    0,
		},
	}
}
