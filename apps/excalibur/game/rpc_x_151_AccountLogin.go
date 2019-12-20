package game

import (
	"time"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) AccountLogin(req *msg.Request, arg *msg.AccountLoginRequest) *msg.AccountLoginResponse {

	g.db.CreateNewUserIfNotExist(req.Uid)

	err := try_do(func() {
		actor := g.use_actor(req.Uid)

		if is_flushed_time_valid_everyday(actor.user.StatInfo.LastLoginTime) {
			actor.user.StatInfo.LoginDaysTotal++
		}

		actor.user.StatInfo.LastLoginTime = time.Now().Unix()

		actor.save()
	})
	if err != msg.Success {
		return &msg.AccountLoginResponse{Err: err}
	}
	return &msg.AccountLoginResponse{}
}
