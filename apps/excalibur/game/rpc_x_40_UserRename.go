package game

import (
	"bytes"
	"net/http"
	"time"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/lib/log"
)

var filter_client = &http.Client{Timeout: time.Duration(1) * time.Second}

func (g *GameService) UserRename(req *msg.Request, arg *msg.UserRenameRequest) *msg.UserRenameResponse {

	user := g.get_user(req.Uid)

	if user.Diamond < consts.ChangeNameDiamond {
		return &msg.UserRenameResponse{Err: msg.ErrDiamondNotEnough}
	}

	res, err := filter_client.Post("http://127.0.0.1:25001/simple/is_valid", "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(arg.NewName)))
	log.Info(err, res.StatusCode)
	if err == nil {
		if res.StatusCode == 400 {
			log.Warn(arg.NewName)
			return &msg.UserRenameResponse{Err: msg.ErrInvalidUserInput}
		}
	}

	user.Nickname = arg.NewName
	user.Diamond -= consts.ChangeNameDiamond

	g.save_user(user)
	return &msg.UserRenameResponse{
		Nickname: user.Nickname,
		Diamond:  user.Diamond,
	}
}
