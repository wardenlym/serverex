package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/util"
)

func (g *GameService) GM_CreateItem(req *msg.Request, arg *msg.GM_CreateItemRequest) *msg.GM_CreateItemResponse {

	user := g.get_user(req.Uid)
	chara := user.Characters[arg.CharacterId]

	index := func() int {
		for i, v := range chara.Bag.Cells {
			if v.Item == nil || v.Num == 0 {
				return i
			}
		}
		return len(chara.Bag.Cells)
	}()

	if index >= len(chara.Bag.Cells) {
		return &msg.GM_CreateItemResponse{Err: msg.ErrBagCellNotEnough}
	}

	newItem := odm.Item{
		Guid:            util.NewSnowflakeID(),
		TypeId:          arg.TypeId,
		StarLevel:       arg.StarLevel,
		StarUpgradeInfo: []int{},
	}

	chara.Bag.Cells[index].Item = &newItem
	chara.Bag.Cells[index].Num = arg.Num
	user.Characters[arg.CharacterId] = chara

	g.save_user(user)

	return &msg.GM_CreateItemResponse{
		NewItem:   newItem,
		Num:       arg.Num,
		CellIndex: index,
	}
}

// func (g *GameService) GM_CreateItem2(req *msg.Request, arg *msg.GM_CreateItemRequest) *msg.GM_CreateItemResponse {

// 	user := g.get_user(req.Uid)
// 	chara := user.Characters[arg.CharacterId]

// 	index := func() int {
// 		for i, v := range chara.Bag.Cells {
// 			if v.Item == nil || v.Num == 0 {
// 				return i
// 			}
// 		}
// 		return len(chara.Bag.Cells)
// 	}()

// 	if index >= len(chara.Bag.Cells) {
// 		return &msg.GM_CreateItemResponse{IsSuccess: false, ErrMsg: "背包空间不足"}
// 	}

// 	newItem := odm.Item{
// 		Guid:            util.NewSnowflakeID(),
// 		TypeId:          arg.TypeId,
// 		StarLevel:       arg.StarLevel,
// 		StarUpgradeInfo: []int{},
// 	}

// 	chara.Bag.Cells[index].Item = &newItem
// 	chara.Bag.Cells[index].Num = arg.Num
// 	user.Characters[arg.CharacterId] = chara

// 	g.save_user(user)

// 	return &msg.GM_CreateItemResponse{
// 		IsSuccess: true,
// 		NewItem:   newItem,
// 		Num:       arg.Num,
// 		CellIndex: index,
// 	}
// }
