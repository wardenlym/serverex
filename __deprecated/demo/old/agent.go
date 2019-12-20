package demo

import (
	"bytes"
	"net"
	"sync"
	"time"

	nats "github.com/nats-io/go-nats"
	"gitlab.com/megatech/serverex/lib/log"
	"gitlab.com/megatech/serverex/lib/util"
)

type AgentManager interface {
	AddAgent(Agent)
	DelAgent(Agent)
	GetById(id string) Agent
	StopAll()
}

type Agent interface {
	Id() string
	RunAsync()
	Stop()
}

type tcpAgent struct {
	agentId  string
	peerConn net.Conn
	mqConn   *MqConn
	decoder  Decoder
	encoder  Encoder
	manager  AgentManager

	stopmu *sync.Mutex
	stop   chan struct{}
}

// tcp / mq / len prefix parser
func NewTCPDefaultAgent(peer net.Conn, mq *MqConn, manager AgentManager) Agent {
	agent := &tcpAgent{
		agentId:  util.NewSnowflakeIDBase58(),
		peerConn: peer,
		mqConn:   mq,
		decoder:  NewPacketDecoder(),
		encoder:  NewPacketEncoder(),
		manager:  manager,
		stopmu:   &sync.Mutex{},
		stop:     make(chan struct{}),
	}
	return agent
}

func (agent *tcpAgent) Id() string {
	return agent.agentId
}

func (self *tcpAgent) RunAsync() {

	// recv
	go func() {

		for {
			buf := bytes.NewBuffer([]byte{})

			// self.peerConn.SetReadDeadline(time.Now().Add(SocketReadTimeout))
			err := self.decoder.Decode(self.peerConn, buf)

			if err != nil {
				if err, ok := err.(net.Error); ok && err.Timeout() {
					// TODO: see issue 2
					continue
				} else {
					log.Error(err)
					break
				}
			}

			packet := &Packet{Data: buf.Bytes()}

			log.Debug("recv:", string(packet.Data))
			err = self.mqConn.Publish(SubjectGateToGame, packet.Data)
			log.ErrorIf(err)
			err = self.mqConn.Flush()
			log.ErrorIf(err)
		}

		log.Info("agent recv exit:", self.agentId)
		self.Stop()
	}()

	// send
	go func() {
		ch := make(chan *nats.Msg)
		subinfo, err := self.mqConn.ChanSubscribe(SubjectGameToGate, ch)
		log.ErrorIf(err, subinfo)

		// first sync packet
		// self.peerConn.SetWriteDeadline(time.Now().Add(SocketWriteTimeout))
		// self.encoder.Encode(&Packet{Code: PktSyn}, self.peerConn)

		for {
			msg, ok := <-ch
			if !ok {
				break
			}
			log.Debug("send:", string(msg.Data))
			// self.peerConn.SetWriteDeadline(time.Now().Add(SocketWriteTimeout))
			err := self.encoder.Encode(self.peerConn, bytes.NewBuffer(msg.Data))
			if err != nil {
				if err, ok := err.(net.Error); ok && err.Timeout() {
					// TODO: see issue 2
					continue
				} else {
					log.Error(err)
					break
				}
			}

		}

		log.Info("agent send exit:", self.agentId)
		self.Stop()
	}()

	// controller
	ticker := time.NewTicker(time.Second * 10)
	go func() {
		func() {
			for {
				select {
				case <-ticker.C:
					log.Debug("agent running:", self.agentId)
				case <-self.stop:
					return
				}
			}
		}()

		log.Debug("agent controller stopping:", self.agentId)
		self.peerConn.Close()

		ticker.Stop()
		self.cleanup()
	}()

}

func (self *tcpAgent) Stop() {
	self.stopmu.Lock()
	if self.stop != nil {
		close(self.stop)
		self.stop = nil
	}
	self.stopmu.Unlock()
}

func (self *tcpAgent) cleanup() {
	self.manager.DelAgent(self)
}

func NewSimpleAgentManager() AgentManager {
	return &agentManager{
		agents: map[string]Agent{},
	}
}

type agentManager struct {
	sync.Mutex
	agents map[string]Agent
}

func (self *agentManager) AddAgent(a Agent) {
	self.Lock()
	self.agents[a.Id()] = a
	self.Unlock()
}

func (self *agentManager) DelAgent(a Agent) {
	self.Lock()
	delete(self.agents, a.Id())
	self.Unlock()
}

func (self *agentManager) GetById(id string) Agent {
	self.Lock()
	defer self.Unlock()
	return self.agents[id]
}

func (self *agentManager) StopAll() {
	self.Lock()
	defer self.Unlock()
	for _, v := range self.agents {
		v.Stop()
	}
}
