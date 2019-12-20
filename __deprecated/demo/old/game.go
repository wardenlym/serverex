package demo

import (
	"github.com/nats-io/go-nats"
	"gitlab.com/megatech/serverex/lib/log"
)

type DemoGameOption struct {
	ListenAddress string
	NATsAddress   string
}

var DefaultDemoGameOption = DemoGameOption{
	NATsAddress: "nats://lcoalhost:4222",
}

type DemoGame struct {
	ServiceInfo
	stop, quit chan struct{}
}

func NewDemoGame(opt DemoGameOption) *DemoGame {
	game := &DemoGame{
		ServiceInfo: NewServiceInfo("game"),
		stop:        make(chan struct{}),
		quit:        make(chan struct{}),
	}
	return game
}

func (game *DemoGame) Run() {

	recvMq, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Error(err)
		close(game.quit)
		return
	}
	sendMq, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Error(err)
		close(game.quit)
		return
	}

	recvMq.Subscribe(SubjectGateToGame, func(msg *nats.Msg) {
		log.Debug("game", msg.Data)
		sendMq.Publish(SubjectGameToGate, []byte("PONG "+string(msg.Data)))
		sendMq.Flush()
	})

	go func() {
		<-game.stop
		close(game.quit)
	}()
}

func (game *DemoGame) Stop() {
	close(game.stop)
	<-game.quit
}
