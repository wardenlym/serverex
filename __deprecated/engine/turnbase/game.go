package turnbase

import (
	"fmt"

	"gitlab.com/megatech/serverex/core/msg"
	"gitlab.com/megatech/serverex/core/svc"
	exlog "gitlab.com/megatech/serverex/lib/log"
)

type Option struct {
}

type game struct {
	svc.Service
	svc.MqClient
	Option
	incoming, outgoing chan *msg.Msg
	SceneManager
}

func NewGame(option Option, mqc svc.MqClient) svc.Game {
	p := &game{
		Option:   option,
		MqClient: mqc,
	}

	p.Service = svc.NewBasicService("game", p.loop)

	return p
}

func (g *game) Incoming() <-chan *msg.Msg {
	return g.incoming
}

func (g *game) Outgoing() chan<- *msg.Msg {
	return g.outgoing
}

func (gate *game) loop(stop chan bool, quit chan error) {

	log := exlog.NewLoggerWithName(fmt.Sprintf("%d.%s", gate.ID(), gate.Name()))
	defer func() {
		close(quit)
		log.Info("stopped.")
	}()

	log.Info("running.")

	for {
		select {
		case <-stop:
			log.Info("shutting down.")
			return
		}
	}
}
