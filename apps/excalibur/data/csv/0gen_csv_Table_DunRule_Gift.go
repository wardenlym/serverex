package csv

type DunRule_Gift struct {
	// uid
	Uid int

	// 规则名
	Name string

	// 文字描述
	Intro []string

	// 怪物数量
	Monster_num int

	// 礼物id
	Gift_id int

	// 礼物数量
	Gift_num int
}

var Table_DunRule_Gift = map[int]DunRule_Gift{}

func Load_Table_DunRule_Gift() {
	Table_DunRule_Gift = map[int]DunRule_Gift{
		1: DunRule_Gift{
			Uid:         1,
			Name:        "雪中送炭",
			Intro:       []string{"[f9e0a9]每杀死[-]{0}[f9e0a9]名怪物[-]|获得{0}"},
			Monster_num: 10,
			Gift_id:     30002,
			Gift_num:    1,
		},
	}
}
