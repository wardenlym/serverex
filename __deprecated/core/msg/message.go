package msg

import (
	"time"
)

type MsgID int64

type PID string

//go:generate stringer -type=CodeType
type CodeType int16

// CodeType
const (
	_ CodeType = CodeType(-iota)
	C_Any
	C_Multicast
)

// CodeType
const (
	C_Invalid CodeType = CodeType(iota)
	C_AgentConnect
	C_AgentDisconnect
	C_AgentInactive
	C_AgentDeactive
)

type Head struct {
	ID    MsgID
	Stamp time.Time
	Src   PID
	Dst   PID
	Code  CodeType
}

type Msg struct {
	Head
	Data []byte
}

func New(src, dst string, code int16, data []byte) *Msg {
	return &Msg{
		Head: Head{
			Stamp: time.Now(),
			Src:   PID(src),
			Dst:   PID(dst),
			Code:  CodeType(code),
		},
		Data: data,
	}
}

func (self *Msg) Clone() *Msg {
	d := make([]byte, len(self.Data))
	copy(d, self.Data)
	return &Msg{
		Head: self.Head,
		Data: d,
	}
}

type Serializer interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte) (interface{}, error)
}

type Parser interface {
	Register(CodeType, Serializer)
	Marshal(Head, interface{}) (Msg, error)
	Unmarshal(Msg) (Head, interface{}, error)
}

type Dispatcher interface {
	Register(CodeType, func(Head, interface{}))
	Process(Msg)
}
