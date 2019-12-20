package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) GetThreadComment(req *msg.Request, arg *msg.GetThreadCommentRequest) *msg.GetThreadCommentResponse {

	a := g.db.Comment().GetThreadComment(arg.ThreadId, arg.IndexFrom, arg.Count)
	return &msg.GetThreadCommentResponse{Comments: a}
}
