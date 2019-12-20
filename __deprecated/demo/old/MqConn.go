package demo

import nats "github.com/nats-io/go-nats"

// 简单屏蔽一下
type MqConn struct {
	*nats.Conn
}

func (*MqConn) Close() {
	panic("should not do this")
}
