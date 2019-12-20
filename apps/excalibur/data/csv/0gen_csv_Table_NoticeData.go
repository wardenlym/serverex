package csv

type NoticeData struct {
	// 公告ID
	Id int

	// 标题
	Title string

	// 内容
	Content string

	// 携带金币
	Gold int

	// 携带钻石
	Diamond int

	// 携带礼品
	Gift []int

	// 发送时间
	Sendingtime string

	// 发送人
	Sender string
}

var Table_NoticeData = map[int]NoticeData{}

func Load_Table_NoticeData() {
	Table_NoticeData = map[int]NoticeData{
		30007: NoticeData{
			Id:          30007,
			Title:       "《无尽秘境》游戏公告",
			Content:     "《无尽秘境》游戏开始测试，为此开发小组给每位前来挑战的冒险者准备了金币*5000 钻石*30000 回城卷轴*5 全知宝石*5的礼包，希望能对您有所帮助，感谢各位冒险者对《无尽秘境》的支持。",
			Gold:        5000,
			Diamond:     30000,
			Gift:        []int{30097, 30097, 30097, 30097, 30097, 30098, 30098, 30098, 30098, 30098},
			Sendingtime: "2018-09-04T19:00:00",
			Sender:      "《无尽秘境》办事处",
		},
	}
}
