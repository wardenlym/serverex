package tcpgate

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"sync"

	"gitlab.com/megatech/serverex/core/msg"
	"gitlab.com/megatech/serverex/core/pkt"
	"gitlab.com/megatech/serverex/core/svc"
	"gitlab.com/megatech/serverex/lib/log"
)

type AgentAdmin struct {
	sync.Mutex
	agents map[int]*Agent
}

func NewAgentAdmin() *AgentAdmin {
	return &AgentAdmin{
		agents: map[int]*Agent{},
	}
}

func (p *AgentAdmin) Add(id int, agent *Agent) {
	p.Lock()
	defer p.Unlock()
	p.agents[id] = agent
}

func (p *AgentAdmin) StopAll() {
	p.Lock()
	defer p.Unlock()
	for _, v := range p.agents {
		v.Stop()
	}
}

func (p *AgentAdmin) Remove(id int) {
	p.Lock()
	defer p.Unlock()
	delete(p.agents, id)
}

type Agent struct {
	id     int
	stop   chan bool
	logger log.Logger
	wg     *sync.WaitGroup
	onexit func()

	peer     *net.TCPConn
	incoming chan *pkt.Pkt
	outgoing chan *pkt.Pkt
	encoder  pkt.Encoder
	decoder  pkt.Decoder
	mqc      svc.MqClient
}

func (a *Agent) Stop() {
	close(a.stop)
}

func NewAgent(peer *net.TCPConn, id int, encoder pkt.Encoder, decoder pkt.Decoder, mqc svc.MqClient, onexit func()) *Agent {
	tcpaddr := peer.RemoteAddr().(*net.TCPAddr)
	logger := log.NewLoggerWithName(fmt.Sprintf("%s.%d.%d", "agent", id, tcpaddr.Port))
	agent := &Agent{
		id:     id,
		stop:   make(chan bool),
		logger: logger,
		wg:     &sync.WaitGroup{},
		onexit: onexit,

		incoming: make(chan *pkt.Pkt, 1024),
		outgoing: make(chan *pkt.Pkt, 1024),
		peer:     peer,
		encoder:  encoder,
		decoder:  decoder,
		mqc:      mqc,
	}

	// mq process
	go mq_subscribe(agent)
	go mq_publish(agent)

	// socket process
	go socket_recv(agent)
	go socket_send(agent)

	// monitor
	go monitor(agent)

	return agent
}

func socket_recv(a *Agent) {
	a.wg.Add(1)
	defer a.wg.Done()
	for {
		select {
		case <-a.stop:
			return
		default:
			pkt, err := a.decoder.Decode(a.peer)
			if err != nil {
				if err == io.EOF {
					a.logger.Info("remote peer closed", err)
					return
				} else {
					log.Error(err)
					return
				}
			}
			log.Debug(pkt.Code, pkt.Data)
			a.incoming <- pkt
		}
	}
}

func monitor(a *Agent) {
	a.wg.Wait()
	a.peer.Close()
	a.onexit()
	a.logger.Info("quit.")
}

func socket_send(a *Agent) {
	a.wg.Add(1)
	defer a.wg.Done()
	for {
		pkt, ok := <-a.outgoing
		if !ok {
			return
		}
		buf := bytes.NewBuffer([]byte{})
		a.encoder.Encode(pkt, buf)

		n, err := a.peer.Write(buf.Bytes())
		if err != nil {
			if err == io.EOF {
				log.Info("remote peer closed", n, err)
				return
			} else {
				log.Error(n, err)
				return
			}
		}
	}
}

func mq_publish(a *Agent) {
	a.wg.Add(1)
	defer a.wg.Done()
	for {
		select {
		case <-a.stop:
			return
		case pkt, ok := <-a.incoming:
			if !ok {
				return
			}
			pkt = pkt
			//  msg := parse(pkt)
			msg := &msg.Msg{}
			a.mqc.Publish("", msg)
		}
	}
}

func mq_subscribe(a *Agent) {
	a.mqc.Subscribe("", func(msg *msg.Msg) {
		pkt := &pkt.Pkt{}
		//  pkt := parse(msg)
		a.outgoing <- pkt
	})
}
