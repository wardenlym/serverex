package natsmq

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/nats-io/gnatsd/server"
	nats "github.com/nats-io/go-nats"
	"gitlab.com/megatech/serverex/core/msg"
	"gitlab.com/megatech/serverex/core/svc"
	exlog "gitlab.com/megatech/serverex/lib/log"
)

// 默认的基于NATS的mq,这个是可以跨机通讯的
// 建议性能没问题的情况都尽量用这个

type Option struct {
	URL string
}

type natsconn struct {
	conn *nats.Conn
}

func NewClient(opt Option) svc.MqClient {
	var err error
	var c *nats.Conn
	for i := 0; i < 5; i++ {
		c, err = nats.Connect(opt.URL, nats.ReconnectHandler(func(*nats.Conn) {
			exlog.Info("Reconnected")
		}))

		if c != nil && c.IsConnected() {
			exlog.Info("connected")
			break
		}
		time.Sleep(time.Second)
		exlog.Info("reconnect: ", i, err)
	}

	if err != nil {
		exlog.Error(err)
		return nil
	}

	p := &natsconn{
		conn: c,
	}
	return p
}

func (c *natsconn) Publish(subj string, msg *msg.Msg) {
	buf := bytes.NewBuffer(nil)
	_ = gob.NewEncoder(buf).Encode(msg)
	c.conn.Publish(subj, buf.Bytes())
}

func (c *natsconn) Subscribe(subj string, f func(*msg.Msg)) {
	c.conn.Subscribe(subj, func(natsMsg *nats.Msg) {
		msg := msg.Msg{}
		buf := bytes.NewBuffer(msg.Data)
		gob.NewDecoder(buf).Decode(&msg)
		f(&msg)
	})
}

func (c *natsconn) Close() {
	c.conn.Close()
}

type natsmq struct {
	svc.Service
}

func New() svc.Mq {
	p := &natsmq{}
	p.Service = svc.NewBasicService("nats", p.loop)
	return p
}

func (p *natsmq) loop(stop chan bool, quit chan error) {

	log := exlog.NewLoggerWithName(fmt.Sprintf("%d.%s", p.ID(), p.Name()))

	opts := &server.Options{}
	// 默认端口监听 4222

	// -m 注意 这个是nats monitor的端口 默认没有开启
	opts.HTTPPort = 8222
	// 防止nats拦截系统信号
	opts.NoSigs = true

	gnatsd := server.New(opts)
	log.Info("start.")
	func() {
		go gnatsd.Start()
	}()
	log.Info("running.")

	<-stop

	log.Info("shuting down.")
	gnatsd.Shutdown()
	log.Info("stopped.")

	close(quit)
}
