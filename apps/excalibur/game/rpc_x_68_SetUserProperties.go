package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) SetUserProperties(req *msg.Request, arg *msg.SetUserPropertiesRequest) *msg.SetUserPropertiesResponse {
	g.db.UserProperties().Set(req.Uid, arg.Key, arg.Doc)
	return &msg.SetUserPropertiesResponse{}
}
