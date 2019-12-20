package svc

import "gitlab.com/megatech/serverex/core/msg"

type Mq interface {
	Service
}

type MqClient interface {
	Publish(subj string, msg *msg.Msg)
	Subscribe(subj string, f func(*msg.Msg))
}
