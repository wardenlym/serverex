package game

import (
	"time"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) GetOperationActivityInfo(req *msg.Request, arg *msg.GetOperationActivityInfoRequest) *msg.GetOperationActivityInfoResponse {
	var operation_activity odm.OperationActivity
	err := try_do(func() {
		actor := g.use_actor(req.Uid)
		operation_activity = actor.user_info().OperationActivity
	})

	if err != msg.Success {
		return &msg.GetOperationActivityInfoResponse{Err: err}
	}

	return &msg.GetOperationActivityInfoResponse{
		ServerTime:        time.Now().Unix(),
		OperationActivity: operation_activity,
	}
}
