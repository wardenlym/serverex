package csv

type Rewards_Week struct {
	// 编号
	ID int

	// 道具奖励
	ItemRewards []int

	// 金币奖励
	GoldRewards int

	// 钻石奖励
	DiamondsRewards int
}

var Table_Rewards_Week = map[int]Rewards_Week{}

func Load_Table_Rewards_Week() {
	Table_Rewards_Week = map[int]Rewards_Week{
		1: Rewards_Week{
			ID:              1,
			ItemRewards:     []int{},
			GoldRewards:     1000,
			DiamondsRewards: 0,
		},
		2: Rewards_Week{
			ID:              2,
			ItemRewards:     []int{57331},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		3: Rewards_Week{
			ID:              3,
			ItemRewards:     []int{10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015, 10015},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		4: Rewards_Week{
			ID:              4,
			ItemRewards:     []int{},
			GoldRewards:     0,
			DiamondsRewards: 100,
		},
		5: Rewards_Week{
			ID:              5,
			ItemRewards:     []int{30099, 30099, 30099},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		6: Rewards_Week{
			ID:              6,
			ItemRewards:     []int{15001, 15001, 15001, 15001, 15001, 15001, 15001, 15001, 15001, 15001},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		7: Rewards_Week{
			ID:              7,
			ItemRewards:     []int{11000, 11000, 11000, 11000, 11000, 11000, 11000, 11000, 11000, 11000},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
	}
}
