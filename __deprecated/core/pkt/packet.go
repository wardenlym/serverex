package pkt

import (
	"io"
)

//go:generate stringer -type=CodeType
type CodeType uint16

// CodeType
const (
	_ CodeType = CodeType(iota)
	// 返回对端传来的相同数据作为回复
	C_Echo
	// 传递本端时间和版本给对端
	C_Handshake
	// 返回双方时间
	C_TimeSync
	// 通知有错误发生后立刻关闭连接
	C_Error
	// 合法的更高层业务消息
	C_Message
)

type PktHandshake struct {
	Version int64
	Time    int64
}

type PktTimeSync struct {
	ClientTime int64
	ServerTime int64
}

type Pkt struct {
	Code CodeType
	Data []byte
}

func New(code int16, data []byte) *Pkt {
	return &Pkt{CodeType(code), data}
}

func (self *Pkt) Clone() *Pkt {
	d := make([]byte, len(self.Data))
	copy(d, self.Data)
	return &Pkt{
		Code: self.Code,
		Data: d,
	}
}

type Encoder interface {
	Encode(*Pkt, io.Writer) error
}

type Decoder interface {
	Decode(io.Reader) (*Pkt, error)
}

type EncoderDecoder interface {
	Encoder
	Decoder
}
