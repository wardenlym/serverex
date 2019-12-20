package gameplay

import (
	"github.com/json-iterator/go"
	"gitlab.com/megatech/serverex/apps/excalibur/msg"
	"gitlab.com/megatech/serverex/lib/log"
)

func (g *game) battleStart(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)

	var o msg.BattleStartRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}
	log.Info("battle with Npc:", o.NpcId)

	_, err = g.db.GetUser(req.Uid)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}
	res.Data = &msg.BattleStartResponse{}
	return res
}

func (g *game) battleEnd(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)

	var o msg.BattleEndRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}
	log.Info("battle end  Npc:", o.NpcId)

	err = g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".battleAttribute", o.BattleAttribute)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}
	res.Data = &msg.BattleEndResponse{}
	return res
}
