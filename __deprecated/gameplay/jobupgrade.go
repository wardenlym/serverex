package gameplay

import (
	jsoniter "github.com/json-iterator/go"
	"gitlab.com/megatech/serverex/apps/excalibur/msg"
)

func (g *game) jobupgrade(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.JobUpgradeRequest
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

	if uint(attr.KillValue) < o.CostKillValue {
		res.Data = &msg.JobUpgradeResponse{ErrMsg: "击杀点数不足"}
		return res
	}
	killValue := uint(attr.KillValue) - o.CostKillValue

	is_upstar, index, star := func() (bool, int, uint) {
		for i, v := range attr.JobTree.JobStatus {
			if v.JobId == o.JobStatus.JobId {
				return true, i, v.JobStar + 1
			}
		}
		return false, -1, 0
	}()

	if is_upstar {
		attr.JobTree.JobStatus[index].JobStar = star
	} else {
		attr.JobTree.JobStatus = append(attr.JobTree.JobStatus, o.JobStatus)
	}

	attr.KillValue = int(killValue)

	g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".battleAttribute", attr)

	res.Data = &msg.JobUpgradeResponse{
		IsSuccess: true,
		JobTree:   attr.JobTree,
		KillValue: killValue,
	}
	return res
}
