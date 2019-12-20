package csv

type GiftPackage struct {
	// 礼包ID
	Id int

	// 礼包道具
	ItemId []int

	// 范围最大值
	NumMax []int

	// 范围最小值
	NumMin []int
}

var Table_GiftPackage = map[int]GiftPackage{}

func Load_Table_GiftPackage() {
	Table_GiftPackage = map[int]GiftPackage{
		100001: GiftPackage{
			Id:     100001,
			ItemId: []int{15002},
			NumMax: []int{1000},
			NumMin: []int{1000},
		},
		100002: GiftPackage{
			Id:     100002,
			ItemId: []int{15002},
			NumMax: []int{5000},
			NumMin: []int{500},
		},
		100003: GiftPackage{
			Id:     100003,
			ItemId: []int{15003},
			NumMax: []int{1000},
			NumMin: []int{1000},
		},
		100004: GiftPackage{
			Id:     100004,
			ItemId: []int{15003},
			NumMax: []int{5000},
			NumMin: []int{500},
		},
	}
}
