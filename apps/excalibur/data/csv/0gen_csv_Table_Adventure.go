package csv

type Adventure struct {
	// 冒险ID
	Adventure_id int

	// 冒险点等级
	Adventure_lv int

	// 名字
	Adventure_name string

	// 描述
	Adventure_describe string

	// 消耗时间
	Adventure_time int

	// 产出材料
	Adventure_get []int

	// 产出数量
	Get_num []int

	// 额外掉落1
	Extra_1 int

	// 额外掉落1区间
	Extra_1_num []int

	// 额外掉落2
	Extra_2 int

	// 额外掉落2区间
	Extra_2_num []int

	// 额外掉落3
	Extra_3 int

	// 额外掉落3区间
	Extra_3_num []int

	// 额外掉落4
	Extra_4 int

	// 额外掉落4区间
	Extra_4_num []int

	// 默认效率
	Efficiency float32

	// 获得道具
	Item int

	// 概率
	Item_probability float32
}

var Table_Adventure = map[int]Adventure{}

func Load_Table_Adventure() {
	Table_Adventure = map[int]Adventure{
		1: Adventure{
			Adventure_id:       1,
			Adventure_lv:       0,
			Adventure_name:     "莱纳河下游",
			Adventure_describe: "莱纳河的下游比较安全，只会出现一些低阶魔物，但是仍然时常可以发现从上游漂流过来的一些宝物，应该可以找到一些进阶材料，过去看一看吧。。。",
			Adventure_time:     60,
			Adventure_get:      []int{14002, 14003, 14004},
			Get_num:            []int{2, 2, 2},
			Extra_1:            11000,
			Extra_1_num:        []int{0, 1},
			Extra_2:            11004,
			Extra_2_num:        []int{0, 1},
			Extra_3:            11008,
			Extra_3_num:        []int{0, 1},
			Extra_4:            0,
			Extra_4_num:        []int{},
			Efficiency:         1,
			Item:               40000,
			Item_probability:   0.05,
		},
		2: Adventure{
			Adventure_id:       2,
			Adventure_lv:       1,
			Adventure_name:     "古城废墟",
			Adventure_describe: "古城废墟已的魔物已经被清除的差不多了，应该不会很危险，但是这里还有很多没有探索到的地区，应该还可以发现一些强化用的材料。。。。",
			Adventure_time:     60,
			Adventure_get:      []int{12006, 12007},
			Get_num:            []int{3, 3},
			Extra_1:            12000,
			Extra_1_num:        []int{0, 2},
			Extra_2:            12001,
			Extra_2_num:        []int{0, 2},
			Extra_3:            12012,
			Extra_3_num:        []int{0, 2},
			Extra_4:            12013,
			Extra_4_num:        []int{0, 2},
			Efficiency:         1,
			Item:               40000,
			Item_probability:   0.05,
		},
		3: Adventure{
			Adventure_id:       3,
			Adventure_lv:       1,
			Adventure_name:     "凯旋桥",
			Adventure_describe: "凯旋桥是层级人类的最后一道防线，这里曾经爆发了人类和魔物的最后一次的战争，现在仍然可以看到当时战争的惨烈场景，去发现一些强化材料吧。。。",
			Adventure_time:     120,
			Adventure_get:      []int{13006, 13007},
			Get_num:            []int{6, 6},
			Extra_1:            13000,
			Extra_1_num:        []int{0, 4},
			Extra_2:            13001,
			Extra_2_num:        []int{0, 4},
			Extra_3:            13012,
			Extra_3_num:        []int{0, 4},
			Extra_4:            13013,
			Extra_4_num:        []int{0, 2},
			Efficiency:         1,
			Item:               40000,
			Item_probability:   0.05,
		},
		4: Adventure{
			Adventure_id:       4,
			Adventure_lv:       2,
			Adventure_name:     "战争遗迹",
			Adventure_describe: "曾经的战场，最大规模的战争在此爆发，无数的英雄在此陨落，现在仍不时有淘金者来此寻宝想要发一笔横财，现在还能搜集到一些制作材料。。。。",
			Adventure_time:     120,
			Adventure_get:      []int{14002, 14003, 14005},
			Get_num:            []int{4, 4, 4},
			Extra_1:            11001,
			Extra_1_num:        []int{0, 2},
			Extra_2:            11005,
			Extra_2_num:        []int{0, 2},
			Extra_3:            11009,
			Extra_3_num:        []int{0, 2},
			Extra_4:            0,
			Extra_4_num:        []int{},
			Efficiency:         1,
			Item:               40000,
			Item_probability:   0.05,
		},
		5: Adventure{
			Adventure_id:       5,
			Adventure_lv:       2,
			Adventure_name:     "莱纳河上游",
			Adventure_describe: "莱纳河的上游是原来魔物的聚集地，现在已经荒废了，见不到魔物的身影，能够搜集到当年魔物锻造武器的材料。。。",
			Adventure_time:     240,
			Adventure_get:      []int{12008, 12009},
			Get_num:            []int{12, 12},
			Extra_1:            12002,
			Extra_1_num:        []int{0, 8},
			Extra_2:            12003,
			Extra_2_num:        []int{0, 8},
			Extra_3:            12014,
			Extra_3_num:        []int{0, 8},
			Extra_4:            12015,
			Extra_4_num:        []int{0, 8},
			Efficiency:         1,
			Item:               40000,
			Item_probability:   0.05,
		},
		6: Adventure{
			Adventure_id:       6,
			Adventure_lv:       3,
			Adventure_name:     "熔岩地穴",
			Adventure_describe: "熔岩地穴十分炽热，但是这里仍然有很多可以探索的地方",
			Adventure_time:     120,
			Adventure_get:      []int{13008, 13009},
			Get_num:            []int{6, 6},
			Extra_1:            13002,
			Extra_1_num:        []int{0, 4},
			Extra_2:            13003,
			Extra_2_num:        []int{0, 4},
			Extra_3:            13014,
			Extra_3_num:        []int{0, 4},
			Extra_4:            13015,
			Extra_4_num:        []int{0, 4},
			Efficiency:         1,
			Item:               40000,
			Item_probability:   0.05,
		},
		7: Adventure{
			Adventure_id:       7,
			Adventure_lv:       3,
			Adventure_name:     "熔岩湖",
			Adventure_describe: "那罗亚火山喷发出来的岩浆在山下聚集成为一个炽热的湖泊，可以发现。。。",
			Adventure_time:     240,
			Adventure_get:      []int{14006, 14007, 14008},
			Get_num:            []int{8, 8, 8},
			Extra_1:            11002,
			Extra_1_num:        []int{0, 4},
			Extra_2:            11006,
			Extra_2_num:        []int{0, 4},
			Extra_3:            11010,
			Extra_3_num:        []int{0, 4},
			Extra_4:            0,
			Extra_4_num:        []int{},
			Efficiency:         1,
			Item:               40000,
			Item_probability:   0.05,
		},
		8: Adventure{
			Adventure_id:       8,
			Adventure_lv:       3,
			Adventure_name:     "那罗亚火山",
			Adventure_describe: "炽热的火山口，随时可能会喷发，这里是",
			Adventure_time:     480,
			Adventure_get:      []int{12010, 12011},
			Get_num:            []int{20, 20},
			Extra_1:            12004,
			Extra_1_num:        []int{0, 16},
			Extra_2:            12005,
			Extra_2_num:        []int{0, 16},
			Extra_3:            12016,
			Extra_3_num:        []int{0, 16},
			Extra_4:            12017,
			Extra_4_num:        []int{0, 16},
			Efficiency:         1,
			Item:               40000,
			Item_probability:   0.05,
		},
		9: Adventure{
			Adventure_id:       9,
			Adventure_lv:       4,
			Adventure_name:     "海底圣所",
			Adventure_describe: "海底圣所中",
			Adventure_time:     120,
			Adventure_get:      []int{13010, 13011},
			Get_num:            []int{6, 6},
			Extra_1:            13004,
			Extra_1_num:        []int{0, 4},
			Extra_2:            13005,
			Extra_2_num:        []int{0, 4},
			Extra_3:            13016,
			Extra_3_num:        []int{0, 4},
			Extra_4:            13017,
			Extra_4_num:        []int{0, 4},
			Efficiency:         1,
			Item:               40000,
			Item_probability:   0.05,
		},
		10: Adventure{
			Adventure_id:       10,
			Adventure_lv:       4,
			Adventure_name:     "荆棘路",
			Adventure_describe: "终于从海底圣所中出来了，后面是一条被玫瑰花铺满的长路看不到尽头，一条充满荆棘的道路中可以搜集。。。。。",
			Adventure_time:     240,
			Adventure_get:      []int{14009, 14010, 14011},
			Get_num:            []int{5, 5, 5},
			Extra_1:            11003,
			Extra_1_num:        []int{0, 3},
			Extra_2:            11007,
			Extra_2_num:        []int{0, 3},
			Extra_3:            11011,
			Extra_3_num:        []int{0, 3},
			Extra_4:            0,
			Extra_4_num:        []int{},
			Efficiency:         1,
			Item:               40000,
			Item_probability:   0.05,
		},
		11: Adventure{
			Adventure_id:       11,
			Adventure_lv:       4,
			Adventure_name:     "世界尽头",
			Adventure_describe: "天涯海角，世界的尽头，能够到达这里的勇者",
			Adventure_time:     480,
			Adventure_get:      []int{14012, 14013, 14014},
			Get_num:            []int{10, 10, 10},
			Extra_1:            11003,
			Extra_1_num:        []int{0, 3},
			Extra_2:            11007,
			Extra_2_num:        []int{0, 3},
			Extra_3:            11011,
			Extra_3_num:        []int{0, 3},
			Extra_4:            0,
			Extra_4_num:        []int{},
			Efficiency:         1,
			Item:               40000,
			Item_probability:   0.05,
		},
	}
}
