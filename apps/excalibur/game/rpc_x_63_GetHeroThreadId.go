package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) GetHeroThreadId(req *msg.Request, arg *msg.GetHeroThreadIdRequest) *msg.GetHeroThreadIdResponse {

	switch arg.HeroId {
	case 10000, 20000, 30000:
		return &msg.GetHeroThreadIdResponse{ThreadId: int64(arg.HeroId)}
	default:
		return &msg.GetHeroThreadIdResponse{Err: msg.ErrInvalidHeroId}
	}
}
