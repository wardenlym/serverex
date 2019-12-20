package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) GetThreadHotComment(req *msg.Request, arg *msg.GetThreadHotCommentRequest) *msg.GetThreadHotCommentResponse {

	a := g.db.Comment().GetThreadHotComment(arg.ThreadId, arg.TopN)
	return &msg.GetThreadHotCommentResponse{Comments: a}
}
