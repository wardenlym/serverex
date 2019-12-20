package csv

type HeroData struct {
	// 英雄ID
	Id int

	// 英雄名称
	Name string

	// 描述
	Desc string

	// 职业
	Job_name string

	// 种族
	Race string

	// 阶数
	Level int

	// 最大攻击力
	Attack_max_value int

	// 最小攻击力
	Attack_min_value int

	// 血量上限
	Hp int

	// 能量
	Mp int

	// 防御最大值
	Def_max_value int

	// 防御最小值
	Def_min_value int

	// 暴击率
	Crit float32

	// 暴击伤害
	Critdamage float32

	// 火属性攻击
	Fire_attack_value int

	// 冰属性攻击
	Ice_attack_value int

	// 毒属性攻击
	Poison_attack_value int

	// 光属性攻击
	Light_attack_value int

	// 暗属性攻击
	Dark_attack_value int

	// 火属性防御
	Fire_def_value int

	// 冰属性防御
	Ice_def_value int

	// 毒属性防御
	Poison_def_value int

	// 光属性防御
	Light_def_value int

	// 暗属性防御
	Dark_def_value int

	// 幸运值
	Lucky int

	// 额外金币
	Extra_gold int

	// 自带被动技能
	Skill []int

	// 击杀值
	Killsvalue int

	// 攻击速度
	AttackSpeed int

	// 特性
	Specialty string

	// 特性加成
	SpecialtyAddRate float32

	// 特性金币
	SpecialtyGold int

	// 特性幸运
	Specialtylucky float32

	// Spine_路径
	Spine_path string

	// 默认皮肤
	Default_skine []int
}

var Table_HeroData = map[int]HeroData{}

func Load_Table_HeroData() {
	Table_HeroData = map[int]HeroData{
		10000: HeroData{
			Id:                  10000,
			Name:                "安杜因",
			Desc:                "人类王国边境领主的儿子，在一次巧合下，登上了开往前线的马车。",
			Job_name:            "平民",
			Race:                "人族",
			Level:               1,
			Attack_max_value:    14,
			Attack_min_value:    8,
			Hp:                  1300,
			Mp:                  100,
			Def_max_value:       0,
			Def_min_value:       0,
			Crit:                0,
			Critdamage:          1.5,
			Fire_attack_value:   0,
			Ice_attack_value:    0,
			Poison_attack_value: 1,
			Light_attack_value:  2,
			Dark_attack_value:   0,
			Fire_def_value:      0,
			Ice_def_value:       0,
			Poison_def_value:    0,
			Light_def_value:     0,
			Dark_def_value:      0,
			Lucky:               0,
			Extra_gold:          0,
			Skill:               []int{1025, 1026},
			Killsvalue:          0,
			AttackSpeed:         100,
			Specialty:           "贪婪|额外产出金币100",
			SpecialtyAddRate:    0.1,
			SpecialtyGold:       5,
			Specialtylucky:      0,
			Spine_path:          "Model/Player/Actor",
			Default_skine:       []int{40014, 40013, 40016, 40015, 40017, 40012, 40019, 40020, 40018},
		},
		20000: HeroData{
			Id:                  20000,
			Name:                "特蕾西",
			Desc:                "奥苏安魔法议会的冰系指导，不过有限的天赋，限制了她的前途。",
			Job_name:            "法师",
			Race:                "精灵",
			Level:               1,
			Attack_max_value:    10,
			Attack_min_value:    9,
			Hp:                  1500,
			Mp:                  100,
			Def_max_value:       0,
			Def_min_value:       0,
			Crit:                0,
			Critdamage:          1.5,
			Fire_attack_value:   1,
			Ice_attack_value:    1,
			Poison_attack_value: 0,
			Light_attack_value:  1,
			Dark_attack_value:   0,
			Fire_def_value:      0,
			Ice_def_value:       0,
			Poison_def_value:    0,
			Light_def_value:     0,
			Dark_def_value:      0,
			Lucky:               0,
			Extra_gold:          0,
			Skill:               []int{2017, 2016},
			Killsvalue:          0,
			AttackSpeed:         100,
			Specialty:           "幸运|额外增加效率10%",
			SpecialtyAddRate:    0.2,
			SpecialtyGold:       0,
			Specialtylucky:      0,
			Spine_path:          "Model/Player/Actor2",
			Default_skine:       []int{40023, 40022, 40025, 40024, 40026, 40021, 40028, 40029, 40027},
		},
		30000: HeroData{
			Id:                  30000,
			Name:                "拉文霍德",
			Desc:                "拉文霍德庄园被毁灭后，过往的行人依然在废墟中看到黑影，久而久之传说甚嚣尘上，成为了著名的“Ravenholdt’s Ghost”。",
			Job_name:            "盗贼",
			Race:                "亡灵",
			Level:               1,
			Attack_max_value:    15,
			Attack_min_value:    7,
			Hp:                  1400,
			Mp:                  100,
			Def_max_value:       0,
			Def_min_value:       0,
			Crit:                0,
			Critdamage:          1.5,
			Fire_attack_value:   0,
			Ice_attack_value:    0,
			Poison_attack_value: 2,
			Light_attack_value:  0,
			Dark_attack_value:   2,
			Fire_def_value:      0,
			Ice_def_value:       0,
			Poison_def_value:    0,
			Light_def_value:     0,
			Dark_def_value:      0,
			Lucky:               0,
			Extra_gold:          0,
			Skill:               []int{3017, 3018},
			Killsvalue:          0,
			AttackSpeed:         100,
			Specialty:           "感知|幸运产出几率增加5%",
			SpecialtyAddRate:    0.1,
			SpecialtyGold:       0,
			Specialtylucky:      0.05,
			Spine_path:          "Model/Player/Actor3",
			Default_skine:       []int{40032, 40031, 40034, 40033, 40035, 40030, 40037, 40038, 40036},
		},
	}
}
