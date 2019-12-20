package tcpgate

import (
	"fmt"
	"net"
	"time"

	"gitlab.com/megatech/serverex/core/pkt"
	"gitlab.com/megatech/serverex/core/svc"
	exlog "gitlab.com/megatech/serverex/lib/log"
)

type Option struct {
	Listen string
}

type TCPGate struct {
	svc.Service
	pkt.Encoder
	pkt.Decoder
	svc.MqClient
	Option
}

func New(option Option, backend svc.MqClient) svc.Gate {
	p := &TCPGate{
		Option:   option,
		MqClient: backend,
		Encoder:  pkt.NewLengthPrefixPacker(),
		Decoder:  pkt.NewLengthPrefixPacker(),
	}
	p.Service = svc.NewBasicService("gate", p.loop)
	return p
}

func (gate *TCPGate) loop(stop chan bool, quit chan error) {
	log := exlog.NewLoggerWithName(fmt.Sprintf("%d.%s", gate.ID(), gate.Name()))

	defer close(quit)

	addr, err := net.ResolveTCPAddr("tcp", gate.Option.Listen)
	if err != nil {
		log.Error(gate.Option.Listen, err)
		return
	}

	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Error(addr, err)
		return
	}

	log.Info("listen at:", listen.Addr())
	acceptor := make(chan *net.TCPConn)

	go func() {
		for {
			peer, err := listen.AcceptTCP()
			if err != nil {
				log.Error(err)
				close(stop)
				return
			} else {
				log.Debug("accept:", peer.RemoteAddr())
				acceptor <- peer
			}
		}
	}()

	log.Info("running.")

	admin := NewAgentAdmin()

	func(agents *AgentAdmin) {
		period := time.NewTicker(time.Second * 10).C
		id := 0
		for {
			select {
			case <-period:
			case <-stop:
				log.Info("shutting down.")
				agents.StopAll()
				return
			case peer := <-acceptor:
				id = id + 1
				agents.Add(
					id,
					NewAgent(peer, id, gate.Encoder, gate.Decoder, gate.MqClient,
						func() {
							agents.Remove(id)
						}))
			}
		}
	}(admin)

	log.Info("stopped.")
}
