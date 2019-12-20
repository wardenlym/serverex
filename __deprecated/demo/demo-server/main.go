package main

import (
	"gitlab.com/megatech/serverex/engine/natsmq"
	"gitlab.com/megatech/serverex/engine/tcpgate"
	"gitlab.com/megatech/serverex/engine/turnbase"
	"gitlab.com/megatech/serverex/lib/log"
	"gitlab.com/megatech/serverex/lib/util"
)

type Config struct {
	Nats natsmq.Option
	Gate tcpgate.Option
	Game turnbase.Option
}

var config Config = Config{
	Nats: natsmq.Option{
		URL: "nats://127.0.0.1:4222",
	},
	Gate: tcpgate.Option{
		Listen: ":11002",
	},
	Game: turnbase.Option{},
}

func exit_if_nil(v interface{}) {
	if v == nil {
		log.Fatal(v)
	}
}

func main() {
	log.Info("demo server start.")

	// 消息通道
	mq := natsmq.New()
	go mq.Run()

	// 客户端连接
	mqc_gate := natsmq.NewClient(config.Nats)
	exit_if_nil(mqc_gate)

	gate := tcpgate.New(config.Gate, mqc_gate)
	go gate.Run()

	mqc_game := natsmq.NewClient(config.Nats)
	exit_if_nil(mqc_game)
	game := turnbase.NewGame(config.Game, mqc_game)
	go game.Run()

	s := <-util.WaitInterruptSignal()
	log.Warn("# signal received:", s)

	game.Stop()
	<-game.Quit()

	gate.Stop()
	<-gate.Quit()

	mq.Stop()
	<-mq.Quit()

	log.Info("demo server end.")
}
