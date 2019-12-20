package csv

type Rewards_FirstPay struct {
	// 编号
	ID int

	// 道具奖励
	ItemRewards []int

	// 金币奖励
	GoldRewards int

	// 钻石奖励
	DiamondsRewards int
}

var Table_Rewards_FirstPay = map[int]Rewards_FirstPay{}

func Load_Table_Rewards_FirstPay() {
	Table_Rewards_FirstPay = map[int]Rewards_FirstPay{
		1: Rewards_FirstPay{
			ID:              1,
			ItemRewards:     []int{57332, 57333, 57334, 30111, 30111, 30111},
			GoldRewards:     3000,
			DiamondsRewards: 0,
		},
	}
}
