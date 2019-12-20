package csv

type RuneopenData struct {
	// 英雄阶数
	Hero_level int

	// 人类
	Hero_10000 []int

	// 精灵
	Hero_20000 []int

	// 亡灵
	Hero_30000 []int
}

var Table_RuneopenData = map[int]RuneopenData{}

func Load_Table_RuneopenData() {
	Table_RuneopenData = map[int]RuneopenData{
		3: RuneopenData{
			Hero_level: 3,
			Hero_10000: []int{0, 10},
			Hero_20000: []int{20, 10},
			Hero_30000: []int{0, 1},
		},
		4: RuneopenData{
			Hero_level: 4,
			Hero_10000: []int{20, 1},
			Hero_20000: []int{0, 21},
			Hero_30000: []int{2, 3},
		},
		5: RuneopenData{
			Hero_level: 5,
			Hero_10000: []int{11, 21},
			Hero_20000: []int{11, 1},
			Hero_30000: []int{4, 5},
		},
		6: RuneopenData{
			Hero_level: 6,
			Hero_10000: []int{2, 12},
			Hero_20000: []int{22, 12},
			Hero_30000: []int{6, 7},
		},
		7: RuneopenData{
			Hero_level: 7,
			Hero_10000: []int{22, 3},
			Hero_20000: []int{2, 23},
			Hero_30000: []int{8, 9},
		},
		8: RuneopenData{
			Hero_level: 8,
			Hero_10000: []int{13, 23},
			Hero_20000: []int{13, 3},
			Hero_30000: []int{10, 11},
		},
		9: RuneopenData{
			Hero_level: 9,
			Hero_10000: []int{4, 14},
			Hero_20000: []int{24, 14},
			Hero_30000: []int{12, 13},
		},
		10: RuneopenData{
			Hero_level: 10,
			Hero_10000: []int{24, 5},
			Hero_20000: []int{4, 25},
			Hero_30000: []int{14, 15},
		},
		11: RuneopenData{
			Hero_level: 11,
			Hero_10000: []int{15, 25},
			Hero_20000: []int{15, 5},
			Hero_30000: []int{16, 17},
		},
		12: RuneopenData{
			Hero_level: 12,
			Hero_10000: []int{6, 16},
			Hero_20000: []int{26, 16},
			Hero_30000: []int{18, 19},
		},
		13: RuneopenData{
			Hero_level: 13,
			Hero_10000: []int{26, 7},
			Hero_20000: []int{6, 27},
			Hero_30000: []int{20, 21},
		},
		14: RuneopenData{
			Hero_level: 14,
			Hero_10000: []int{17, 27},
			Hero_20000: []int{17, 7},
			Hero_30000: []int{22, 23},
		},
		15: RuneopenData{
			Hero_level: 15,
			Hero_10000: []int{8, 18},
			Hero_20000: []int{28, 18},
			Hero_30000: []int{24, 25},
		},
		16: RuneopenData{
			Hero_level: 16,
			Hero_10000: []int{28, 9},
			Hero_20000: []int{8, 29},
			Hero_30000: []int{26, 27},
		},
		17: RuneopenData{
			Hero_level: 17,
			Hero_10000: []int{19, 29},
			Hero_20000: []int{19, 9},
			Hero_30000: []int{28, 29},
		},
	}
}
