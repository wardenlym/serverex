package demo

import (
	"github.com/nats-io/gnatsd/server"
	"gitlab.com/megatech/serverex/lib/log"
)

type NATS struct {
	ServiceInfo
	stop, quit chan struct{}
}

func NewNATS() *NATS {
	return &NATS{
		ServiceInfo: NewServiceInfo("nats"),
		stop:        make(chan struct{}),
		quit:        make(chan struct{}),
	}

}

func (p *NATS) Run() {

	opts := &server.Options{}
	// 默认端口监听 4222

	// -m 注意 这个是nats monitor的端口 默认没有开启
	opts.HTTPPort = 8222
	// 防止nats拦截系统信号
	opts.NoSigs = true

	gnatsd := server.New(opts)
	log.Info("gnatsd start.")
	func() {
		go gnatsd.Start()
	}()
	log.Info("gnatsd running.")
	<-p.stop
	log.Info("gnatsd stop.")
	gnatsd.Shutdown()
	log.Info("gnatsd shutdown.")

	close(p.quit)
}

func (p *NATS) Stop() {
	close(p.stop)
	<-p.quit
}
