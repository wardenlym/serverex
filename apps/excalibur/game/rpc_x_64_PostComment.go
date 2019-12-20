package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func (g *GameService) PostComment(req *msg.Request, arg *msg.PostCommentRequest) *msg.PostCommentResponse {

	err := try_do(func() {
		actor := g.use_actor(req.Uid)
		g.db.Comment().PostComment(
			arg.ThreadId,
			actor.user_info().Uid,
			actor.user_info().Nickname,
			arg.Content)
	})
	if err != msg.Success {
		return &msg.PostCommentResponse{Err: err}
	}
	return &msg.PostCommentResponse{}
}
