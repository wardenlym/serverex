package game

import "time"

var consts = struct {
	OnSalesRefreshInterval     time.Duration
	OnSalesDisplayCount        int
	BagExpandCount             int
	StashExpandCount           int
	FirstStageId               int
	ChangeNameDiamond          uint
	ExploreChangeDiamand       uint
	VipYueKaMaterialExpireDays int
	VipYueKaMaterialDiamond    uint
}{
	OnSalesRefreshInterval:     6 * time.Hour, //市场自然刷新时间
	OnSalesDisplayCount:        8,             //市场显示个数
	BagExpandCount:             15,            //背包每次扩充数量
	StashExpandCount:           5,             //仓库扩充数量
	FirstStageId:               1001,          //地城首层
	ChangeNameDiamond:          100,           //改名耗费钻石
	ExploreChangeDiamand:       100,           //派出冒险增加次数耗费钻石
	VipYueKaMaterialExpireDays: 30,            //vip材料月卡有效时间
	VipYueKaMaterialDiamond:    50,            //vip材料月卡领取钻石
}

func stageid_to_chapterid(stageid int) int {
	return stageid / 1000 * 1000
}
