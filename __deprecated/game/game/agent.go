package game

import (
	"fmt"
	"io"
	"net"
	"sync"

	"gitlab.com/megatech/serverex/lib/log"
)

type AgentID int

type Agent struct {
	ID       AgentID
	peer     *net.TCPConn
	stop     <-chan bool
	wg       *sync.WaitGroup
	log      log.Logger
	onexit   func()
	incoming chan []byte
	outgoing chan []byte
}

func NewAgent(id AgentID, peer *net.TCPConn, stop <-chan bool, wg *sync.WaitGroup, onexit func()) *Agent {

	tcpaddr := peer.RemoteAddr().(*net.TCPAddr)
	log := log.NewLoggerWithName(fmt.Sprintf("%s.%d:%d", "agent", id, tcpaddr.Port))

	p := &Agent{
		ID:       id,
		peer:     peer,
		stop:     stop,
		wg:       wg,
		log:      log,
		onexit:   onexit,
		incoming: make(chan []byte, 128),
		outgoing: make(chan []byte, 128),
	}

	return p
}

func (p *Agent) loop() {

	p.wg.Add(1)
	defer p.wg.Done()

	go func() {
		for {
			b := make([]byte, 1024)
			n, err := p.peer.Read(b)
			if err != nil {
				if err == io.EOF {
					p.log.Info("recv:remote peer closed", err)
				} else {
					log.Error(n, err)
				}
				return
			}
			p.incoming <- b
		}

	}()

	func() {
		for {
			select {
			case <-p.stop:
				return
			case packet, ok := <-p.incoming:
				p.process(packet)
				if !ok {
					return
				}
			case tosend, ok := <-p.outgoing:
				if tosend != nil {
					n, err := p.peer.Write(tosend)
					if err != nil {
						if err == io.EOF {
							p.log.Info("send:remote peer closed", err)
						} else {
							log.Error(n, err)
						}
						return
					}
				}
				if !ok {
					return
				}
			}
		}
	}()

	p.peer.Close()
	close(p.outgoing)
	close(p.incoming)

	p.onexit()
}

func (p *Agent) send(packet []byte) {
	p.outgoing <- packet
}

func (p *Agent) process([]byte) {

}
