package gate

import "gitlab.com/megatech/serverex/svc"

type Gate struct {
}

func NewGate() *Gate {
	return &Gate{}
}

func (gate *Gate) OnConnect(svc.Peer) {
	panic("not implemented")
}

func (gate *Gate) OnDisconnect(svc.PeerID) {
	panic("not implemented")
}

func (gate *Gate) HandleCall(svc.Peer, interface{}) interface{} {
	panic("not implemented")
}

func (gate *Gate) HandleCast(svc.Peer, interface{}) {
	panic("not implemented")
}

func (gate *Gate) Stop() {
	panic("not implemented")
}
