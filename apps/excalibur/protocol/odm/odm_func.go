package odm

import (
	"time"

	"gitlab.com/megatech/serverex/lib/util"
)

func NewUser() *User {
	p := &User{
		CharacterIds: []string{},
		Characters:   map[string]Character{},
	}
	return p
}

func NewUserWith1CharacterDefault(uid int64) *User {
	p := &User{
		Uid:          uid,
		Nickname:     "游客" + util.Int64ToBase36(uid),
		CreatedAt:    "2018-01-01T00:00:00+08:00",
		UpdatedAt:    "2018-01-01T00:00:00+08:00",
		Diamond:      1000,
		OnSaleInfos:  []OnSaleInfo{},
		CharacterIds: NewCharacterIdsList(),
		Characters: map[string]Character{
			"1": NewCharacters(10000),
		},
		Stash:              NewStash(),
		ExploreInfo:        NewExploreInfo("1", 10000),
		BestDungeonRanking: []DungeonRankingInfo{},
		StatInfo:           StatInfo{CreateTime: time.Now().Unix()},
	}
	return p
}

func NewCharacterIdsList() []string {
	return []string{"1"}
}

func NewCharacters(heroid int) Character {
	return Character{
		HeroId:          heroid,
		Level:           1,
		Bag:             NewBagWithNewbeeEquipment(),
		DungeonStatus:   NewDungeonStatus(),
		BattleAttribute: NewAttribute(heroid),
		RuneTree:        NewRuneTree(),
	}
}

func NewUserWithApprovalCharacters(uid int64, level int) *User {
	p := &User{
		Uid:          uid,
		Nickname:     "noname",
		CreatedAt:    "2018-01-01T00:00:00+08:00",
		UpdatedAt:    "2018-01-01T00:00:00+08:00",
		Diamond:      1000,
		OnSaleInfos:  []OnSaleInfo{},
		CharacterIds: []string{"1"},
		Characters: map[string]Character{
			"1": NewApprovalCharacters(10000, level),
		},
		Stash: NewStash(),
	}
	return p
}

func NewApprovalCharacters(heroid int, level int) Character {
	return Character{
		HeroId:          heroid,
		Level:           level,
		Bag:             NewBagWithApprovalEquipment(),
		DungeonStatus:   NewDungeonStatus(),
		BattleAttribute: NewAttribute(heroid),
		RuneTree:        NewRuneTree(),
	}
}

func NewAttribute(id int) Attribute {
	return Attribute{
		JobTree: NewJobTree(id),
	}
}

func NewEmptyBag() Bag {
	return Bag{
		Equipments: NewCells(9),
		Cells:      NewCells(15),
	}
}

func NewCell(typeid int, num uint) Cell {
	return Cell{
		Item: NewItem(typeid),
		Num:  num,
	}
}

func NewItem(typeid int) *Item {
	return &Item{
		Guid:            util.NewSnowflakeID(),
		TypeId:          typeid,
		StarLevel:       0,
		StarUpgradeInfo: []int{},
	}
}

func NewBagWithNewbeeEquipment() Bag {
	b := Bag{
		Equipments: NewCells(9),
		Cells:      NewCells(15),
	}
	b.Cells[0] = NewCell(30002, 5)
	b.Cells[1] = NewCell(30016, 5)
	b.Cells[2] = NewCell(57335, 1)
	b.Cells[3] = NewCell(12000, 3)
	b.Cells[4] = NewCell(12001, 3)
	b.Cells[5] = NewCell(15000, 5)
	b.Cells[6] = NewCell(15001, 5)
	return b
}

func NewBagWithApprovalEquipment() Bag {
	b := Bag{
		Equipments: NewCells(9),
		Cells:      NewCells(15),
	}
	b.Cells[0] = NewCell(40009, 1)
	b.Cells[1] = NewCell(40010, 1)
	return b
}

func NewStash() Stash {
	return Stash{
		Gold:  100,
		Cells: NewCells(10),
	}
}

func NewCells(n int) []Cell {
	c := make([]Cell, n)
	return c
}

func NewRuneTree() RuneTree {
	c := []RuneCell{}

	for i := 0; i < 30; i++ {
		c = append(c, RuneCell{
			SlotColor: (i / 10) + 1,
			Enable: func() bool {
				// if (i % 10) < 5 {
				// 	return true
				// }
				return false
			}(),
		})
	}

	return RuneTree{c}
}

const (
	S_InTown                    string = "S_InTown"
	S_InDungeon                 string = "S_InDungeon"
	S_InStage                   string = "S_InStage"
	S_InBattle                  string = "S_InBattle"
	S_Dead                      string = "S_Dead"
	Battle_Victory              string = "Battle_Victory"
	Battle_Dead                 string = "Battle_Dead"
	Battle_Escape               string = "Battle_Escape"
	PersonalRanking_Gold        string = "PersonalRanking_Gold"
	PersonalRanking_Kill        string = "PersonalRanking_Kill"
	PersonalRanking_Achievement string = "PersonalRanking_Achievement"
)

func NewDungeonStatus() DungeonStatus {
	return DungeonStatus{
		State:         S_InTown,
		StageProgress: []StageInfo{},
		CurrentStageInfo: StageInfo{
			NpcList: []Npc{},
			BoxList: []Box{},
		},
		KilledNpcList: []Npc{},
		OpenBoxList:   []Npc{},
		UnpickedLoot:  []Cell{},
		LootTotal:     []Cell{},
	}
}

func NewJobTree(id int) JobTree {
	return JobTree{
		JobStatus: []JobStatus{
			JobStatus{
				JobId:   id,
				JobStar: 1,
			},
		},
	}
}

func NewExploreInfo(id string, heroid int) ExploreInfo {
	info := ExploreInfo{
		ChargeTotal: 5,
		Status:      []*ExploreStatus{},
	}

	info.Status = append(info.Status, &ExploreStatus{
		CharacterId: id,
	})

	return info
}
