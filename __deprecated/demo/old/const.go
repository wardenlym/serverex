package demo

import "time"

const (
	SubjectGateToGame     = "demo.gate.to.game"
	SubjectGameToGate     = "demo.game.to.gate"
	SocketReadTimeout     = time.Millisecond * 2000
	SocketWriteTimeout    = time.Millisecond * 200
	SocketConnectTimeout  = time.Millisecond * 200
	GateHeartBeatInterval = time.Second * 6
)
