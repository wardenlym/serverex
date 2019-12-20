package game

import (
	"fmt"
	"net"
	"sync"
	"time"

	"gitlab.com/megatech/serverex/core/svc"
	"gitlab.com/megatech/serverex/lib/log"
)

type Option struct {
	Listen string
}

var DefaultOption = Option{
	Listen: ":11003",
}

type Game struct {
	svc.Service
	Option
}

func NewGame(option Option) *Game {
	p := &Game{Option: option}
	p.Service = svc.NewBasicService("game", p.loop)
	return p
}

func (p *Game) loop(stop chan bool, quit chan error) {
	log := log.NewLoggerWithName(fmt.Sprintf("%d.%s", p.ID(), p.Name()))

	defer close(quit)

	addr, err := net.ResolveTCPAddr("tcp", p.Option.Listen)
	if err != nil {
		log.Error(p.Option.Listen, err)
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
	wg := &sync.WaitGroup{}

	func(agents *AgentAdmin) {
		period := time.NewTicker(time.Second * 10).C
		for {
			select {
			case <-period:
				log.Debug("tick.")
			case <-stop:
				log.Info("shutting down.")
				wg.Wait()
				return
			case peer := <-acceptor:
				agents.RunAgentAsync(peer, stop, wg)
			}
		}
	}(admin)

	log.Info("stopped.")
}
