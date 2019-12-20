package gameplay

import (
	jsoniter "github.com/json-iterator/go"
	"gitlab.com/megatech/serverex/apps/excalibur/msg"
)

func (g *game) levelUp(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.LevelUpRequest
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

	if info.Level != o.CurrentLevel {
		res.Data = &msg.LevelUpResponse{ErrMsg: "请求等级错误"}
		return res
	}

	level := info.Level + 1
	bag := info.Bag

	ok, cells := checkMaterials(o.Materials, info.Bag.Cells)
	if !ok {
		res.Data = &msg.LevelUpResponse{ErrMsg: "材料不足"}
		return res
	}
	bag.Cells = cells

	g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".level", level)
	g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".bag", bag)

	res.Data = &msg.LevelUpResponse{
		IsSuccess:    true,
		CurrentLevel: level,
		BagStatus:    bag,
	}
	return res
}
