package csv

type DunRule_GhostWd struct {
	// 骨戒等级
	Level int

	// 规则名
	Name string

	// 文字描述
	Intro []string

	// 骨戒的物品id
	Ring_id int

	// 怪物类型1
	Cat_one []int

	// 怪物类型2
	Cat_two []int

	// 怪物类型3
	Cat_three []int

	// 怪物类型4
	Cat_four []int
}

var Table_DunRule_GhostWd = map[int]DunRule_GhostWd{}

func Load_Table_DunRule_GhostWd() {
	Table_DunRule_GhostWd = map[int]DunRule_GhostWd{
		1: DunRule_GhostWd{
			Level:     1,
			Name:      "亡者世界",
			Intro:     []string{"[f9e0a9]收集亡者的灵魂，提升{0}的能力[-]"},
			Ring_id:   59000,
			Cat_one:   []int{1, 5},
			Cat_two:   []int{3, 5},
			Cat_three: []int{6, 5},
			Cat_four:  []int{9, 5},
		},
		2: DunRule_GhostWd{
			Level:     2,
			Name:      "亡者世界",
			Intro:     []string{"[f9e0a9]收集亡者的灵魂，提升{0}的能力[-]"},
			Ring_id:   59001,
			Cat_one:   []int{1, 5},
			Cat_two:   []int{3, 5},
			Cat_three: []int{6, 5},
			Cat_four:  []int{9, 5},
		},
		3: DunRule_GhostWd{
			Level:     3,
			Name:      "亡者世界",
			Intro:     []string{"[f9e0a9]收集亡者的灵魂，提升{0}的能力[-]"},
			Ring_id:   59002,
			Cat_one:   []int{1, 5},
			Cat_two:   []int{3, 5},
			Cat_three: []int{6, 5},
			Cat_four:  []int{9, 5},
		},
		4: DunRule_GhostWd{
			Level:     4,
			Name:      "亡者世界",
			Intro:     []string{"[f9e0a9]收集亡者的灵魂，提升{0}的能力[-]"},
			Ring_id:   59003,
			Cat_one:   []int{1, 5},
			Cat_two:   []int{3, 5},
			Cat_three: []int{6, 5},
			Cat_four:  []int{9, 5},
		},
		5: DunRule_GhostWd{
			Level:     5,
			Name:      "亡者世界",
			Intro:     []string{"[f9e0a9]收集亡者的灵魂，提升{0}的能力[-]"},
			Ring_id:   59004,
			Cat_one:   []int{1, 5},
			Cat_two:   []int{3, 5},
			Cat_three: []int{6, 5},
			Cat_four:  []int{9, 5},
		},
		6: DunRule_GhostWd{
			Level:     6,
			Name:      "亡者世界",
			Intro:     []string{"[f9e0a9]收集亡者的灵魂，提升{0}的能力[-]"},
			Ring_id:   59005,
			Cat_one:   []int{1, 5},
			Cat_two:   []int{3, 5},
			Cat_three: []int{6, 5},
			Cat_four:  []int{9, 5},
		},
		7: DunRule_GhostWd{
			Level:     7,
			Name:      "亡者世界",
			Intro:     []string{"[f9e0a9]收集亡者的灵魂，提升{0}的能力[-]"},
			Ring_id:   59006,
			Cat_one:   []int{1, 5},
			Cat_two:   []int{3, 5},
			Cat_three: []int{6, 5},
			Cat_four:  []int{9, 5},
		},
		8: DunRule_GhostWd{
			Level:     8,
			Name:      "亡者世界",
			Intro:     []string{"[f9e0a9]收集亡者的灵魂，提升{0}的能力[-]"},
			Ring_id:   59007,
			Cat_one:   []int{1, 5},
			Cat_two:   []int{3, 5},
			Cat_three: []int{6, 5},
			Cat_four:  []int{9, 5},
		},
		9: DunRule_GhostWd{
			Level:     9,
			Name:      "亡者世界",
			Intro:     []string{"[f9e0a9]收集亡者的灵魂，提升{0}的能力[-]"},
			Ring_id:   59008,
			Cat_one:   []int{1, 5},
			Cat_two:   []int{3, 5},
			Cat_three: []int{6, 5},
			Cat_four:  []int{9, 5},
		},
	}
}
