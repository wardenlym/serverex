package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) GetUserProperties(req *msg.Request, arg *msg.GetUserPropertiesRequest) *msg.GetUserPropertiesResponse {

	return &msg.GetUserPropertiesResponse{
		Doc: g.db.UserProperties().Get(req.Uid, arg.Key),
	}
}
