package gameplay

import (
	"errors"
	"math/rand"

	"github.com/json-iterator/go"
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/odm"
	"gitlab.com/megatech/serverex/lib/util"
)

func (g *game) enterStage(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)

	var o msg.EnterStageRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}

	info, err := g.db.GetUser(req.Uid)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	err = g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".lastStageInfo", o.StageInfo)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	mapinfo, err, errMsg := makeMapInfo(o.StageInfo.StageId)
	if err != nil {
		res.Data = &msg.EnterStageResponse{
			IsSuccess: false,
			ErrMsg:    errMsg,
		}
		return res
	}

	err = g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".lastMapInfo", mapinfo)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	res.Data = &msg.EnterStageResponse{
		IsSuccess:       true,
		MapInfo:         mapinfo,
		BattleAttribute: info.Characters[o.CharacterId].BattleAttribute,
	}
	return res
}

func makeMapInfo(stageId int) (odm.MapInfo, error, string) {
	info, exist := csv.Table_StageData[stageId]
	if !exist {
		return odm.MapInfo{}, errors.New(""), "stage id not exist"
	}

	npclist := []odm.Npc{}

	for _, v := range info.NormalMonsterList {
		npclist = append(npclist, odm.Npc{
			NpcId:     len(npclist),
			NpcTypeId: v,
		})
	}
	for _, v := range info.EliteMonsterList {
		npclist = append(npclist, odm.Npc{
			NpcId:     len(npclist),
			NpcTypeId: v,
		})
	}
	for _, v := range info.BossList {
		npclist = append(npclist, odm.Npc{
			NpcId:     len(npclist),
			NpcTypeId: v,
		})
	}

	return odm.MapInfo{
		MapId: func() int {
			if len(info.MapId) == 1 {
				return info.MapId[0]
			} else {
				return util.RandomIn(info.MapId[0], info.MapId[1])
			}
		}(),
		EntranceType: info.EntranceType,
		Seed:         int(rand.Int31()),
		NpcList:      npclist,
	}, nil, ""
}

func (g *game) exitStage(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.ExitStageRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}
	res.Data = &msg.ExitStageResponse{}
	return res
}
