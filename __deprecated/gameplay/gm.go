package gameplay

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"gitlab.com/megatech/serverex/apps/excalibur/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/odm"
	"gitlab.com/megatech/serverex/lib/util"
)

func (g *game) gm_CreateUser(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	g.db.GetUserOrUpsert(req.Uid)
	res.Data = &msg.GM_CreateUserResponse{}
	return res
}

func (g *game) gm_ResetUser(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	g.db.DeleteUser(req.Uid)
	g.db.GetUserOrUpsert(req.Uid)
	res.Data = &msg.GM_ResetUserResponse{}
	return res
}

func (g *game) gm_CreateItem(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.GM_CreateItemRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}

	info, err := g.db.GetCharacter(req.Uid, o.CharacterId)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	bag := info.Bag

	index := func() int {
		for i, v := range bag.Cells {
			if v.Item == nil || v.Num == 0 {
				return i
			}
		}
		return bag.CellsCapacity
	}()

	if index >= bag.CellsCapacity {
		res.Data = &msg.GM_CreateItemResponse{IsSuccess: false, ErrMsg: "背包空间不足"}
		return res
	}

	newCell := &msg.GM_CreateItemResponse{
		IsSuccess: true,
		NewItem: odm.Item{
			Guid:            util.NewSnowflakeID(),
			TypeId:          o.TypeId,
			StarLevel:       o.StarLevel,
			StarUpgradeInfo: []int{},
		},
		Num:       o.Num,
		CellIndex: index,
	}

	cell := odm.Cell{
		Index: index,
		Num:   o.Num,
		Item:  &newCell.NewItem,
	}
	err = g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".bag.cells."+fmt.Sprint(index), cell)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	res.Data = newCell
	return res
}

func (g *game) gm_SetMyMoney(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.GM_SetMyMoneyRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}
	g.db.UpsertUserField(req.Uid, "gold", o.Gold)
	g.db.UpsertUserField(req.Uid, "diamond", o.Diamond)
	res.Data = &msg.GM_SetMyMoneyResponse{IsSuccess: true}
	return res
}

func (g *game) gm_ResetJob(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.GM_ResetJobRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}

	info, err := g.db.GetCharacter(req.Uid, o.CharacterId)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	attr := info.BattleAttribute
	attr.KillValue = int(o.KillValue)
	attr.JobTree = odm.JobTree{
		JobStatus: []odm.JobStatus{
			odm.JobStatus{
				JobId:   10000,
				JobStar: 1,
			},
		},
	}
	g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".battleAttribute", attr)
	res.Data = &msg.GM_ResetJobResponse{
		JobTree:   attr.JobTree,
		KillValue: uint(attr.KillValue),
	}
	return res
}
