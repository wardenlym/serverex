package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) PraiseComment(req *msg.Request, arg *msg.PraiseCommentRequest) *msg.PraiseCommentResponse {

	g.db.Comment().PraiseComment(
		arg.ThreadId,
		arg.CommentId)

	return &msg.PraiseCommentResponse{}
}
