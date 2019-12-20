package demo

import (
	"net"

	nats "github.com/nats-io/go-nats"
	"gitlab.com/megatech/serverex/lib/log"
)

type TCPGateOption struct {
	ListenAddress string
	NATsAddress   string
}

var DefaultTCPGateOption = TCPGateOption{
	ListenAddress: ":11002",
	NATsAddress:   "nats://localhost:4222",
}

type TCPGate struct {
	ServiceInfo
	option       TCPGateOption
	stop, quit   chan struct{}
	agentManager AgentManager
}

func NewTCPGate(opt TCPGateOption) *TCPGate {
	gate := &TCPGate{
		ServiceInfo:  NewServiceInfo("gate"),
		option:       opt,
		stop:         make(chan struct{}),
		quit:         make(chan struct{}),
		agentManager: NewSimpleAgentManager(),
	}
	return gate
}

func (gate *TCPGate) Run() {

	log.Info(gate.option)
	mq, err := nats.Connect(gate.option.NATsAddress)
	if err != nil {
		log.Error(err)
		close(gate.quit)
		return
	}

	listen, err := net.Listen("tcp", gate.option.ListenAddress)
	if err != nil {
		log.Error(err)
		mq.Close()
		close(gate.quit)
		return
	}

	go func() {
		<-gate.stop
		listen.Close()
		mq.Close()
		gate.agentManager.StopAll()
		close(gate.quit)
	}()

	for {
		peer, err := listen.Accept()
		if err != nil {
			log.Error(err)
			break
		}
		agent := NewTCPDefaultAgent(peer, &MqConn{mq}, gate.agentManager)
		gate.agentManager.AddAgent(agent)
		agent.RunAsync()
	}
}

func (gate *TCPGate) Stop() {
	close(gate.stop)
	<-gate.quit
}
