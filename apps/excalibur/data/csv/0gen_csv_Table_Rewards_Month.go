package csv

type Rewards_Month struct {
	// 编号
	ID int

	// 道具奖励
	ItemRewards []int

	// 金币奖励
	GoldRewards int

	// 钻石奖励
	DiamondsRewards int
}

var Table_Rewards_Month = map[int]Rewards_Month{}

func Load_Table_Rewards_Month() {
	Table_Rewards_Month = map[int]Rewards_Month{
		1: Rewards_Month{
			ID:              1,
			ItemRewards:     []int{},
			GoldRewards:     100,
			DiamondsRewards: 0,
		},
		2: Rewards_Month{
			ID:              2,
			ItemRewards:     []int{},
			GoldRewards:     0,
			DiamondsRewards: 20,
		},
		3: Rewards_Month{
			ID:              3,
			ItemRewards:     []int{11012, 11012, 11012},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		4: Rewards_Month{
			ID:              4,
			ItemRewards:     []int{12000, 12000, 12000},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		5: Rewards_Month{
			ID:              5,
			ItemRewards:     []int{12001, 12001, 12001},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		6: Rewards_Month{
			ID:              6,
			ItemRewards:     []int{15000, 15000, 15000},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		7: Rewards_Month{
			ID:              7,
			ItemRewards:     []int{20057},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		8: Rewards_Month{
			ID:              8,
			ItemRewards:     []int{},
			GoldRewards:     200,
			DiamondsRewards: 0,
		},
		9: Rewards_Month{
			ID:              9,
			ItemRewards:     []int{},
			GoldRewards:     0,
			DiamondsRewards: 30,
		},
		10: Rewards_Month{
			ID:              10,
			ItemRewards:     []int{11012, 11012, 11012, 11012, 11012},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		11: Rewards_Month{
			ID:              11,
			ItemRewards:     []int{12000, 12000, 12000, 12000, 12000},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		12: Rewards_Month{
			ID:              12,
			ItemRewards:     []int{12001, 12001, 12001, 12001, 12001},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		13: Rewards_Month{
			ID:              13,
			ItemRewards:     []int{11015, 11015, 11015},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		14: Rewards_Month{
			ID:              14,
			ItemRewards:     []int{20056},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		15: Rewards_Month{
			ID:              15,
			ItemRewards:     []int{},
			GoldRewards:     300,
			DiamondsRewards: 0,
		},
		16: Rewards_Month{
			ID:              16,
			ItemRewards:     []int{},
			GoldRewards:     0,
			DiamondsRewards: 40,
		},
		17: Rewards_Month{
			ID:              17,
			ItemRewards:     []int{11012, 11012, 11012, 11012, 11012, 11012, 11012, 11012},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		18: Rewards_Month{
			ID:              18,
			ItemRewards:     []int{12000, 12000, 12000, 12000, 12000, 12000, 12000, 12000},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		19: Rewards_Month{
			ID:              19,
			ItemRewards:     []int{12001, 12001, 12001, 12001, 12001, 12001, 12001, 12001},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		20: Rewards_Month{
			ID:              20,
			ItemRewards:     []int{15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		21: Rewards_Month{
			ID:              21,
			ItemRewards:     []int{20054},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		22: Rewards_Month{
			ID:              22,
			ItemRewards:     []int{},
			GoldRewards:     400,
			DiamondsRewards: 0,
		},
		23: Rewards_Month{
			ID:              23,
			ItemRewards:     []int{},
			GoldRewards:     0,
			DiamondsRewards: 50,
		},
		24: Rewards_Month{
			ID:              24,
			ItemRewards:     []int{11012, 11012, 11012, 11012, 11012, 11012, 11012, 11012, 11012, 11012, 11012, 11012, 11012, 11012, 11012},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		25: Rewards_Month{
			ID:              25,
			ItemRewards:     []int{12000, 12000, 12000, 12000, 12000, 12000, 12000, 12000, 12000, 12000, 12000, 12000, 12000, 12000, 12000},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		26: Rewards_Month{
			ID:              26,
			ItemRewards:     []int{12001, 12001, 12001, 12001, 12001, 12001, 12001, 12001, 12001, 12001, 12001, 12001, 12001, 12001, 12001},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		27: Rewards_Month{
			ID:              27,
			ItemRewards:     []int{15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		28: Rewards_Month{
			ID:              28,
			ItemRewards:     []int{20055},
			GoldRewards:     0,
			DiamondsRewards: 0,
		},
		29: Rewards_Month{
			ID:              29,
			ItemRewards:     []int{},
			GoldRewards:     500,
			DiamondsRewards: 0,
		},
		30: Rewards_Month{
			ID:              30,
			ItemRewards:     []int{},
			GoldRewards:     0,
			DiamondsRewards: 60,
		},
		31: Rewards_Month{
			ID:              31,
			ItemRewards:     []int{},
			GoldRewards:     0,
			DiamondsRewards: 120,
		},
	}
}
