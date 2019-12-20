package svc

import "gitlab.com/megatech/serverex/core/pkt"

type Gate interface {
	Service
	pkt.Encoder
	pkt.Decoder
	MqClient
}
