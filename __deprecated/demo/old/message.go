package demo

import "time"

type MessageID int64
type MessageCode int

func (m MessageCode) String() string {
	return "abc"
}

const (
	_ = MessageCode(iota)
	AgentConnected
	AgentDisconnect
	AgentInactive
	AgentDeactive
)

type MessageHead struct {
	ID        MessageID
	Timestamp time.Time
	From      string
	To        string
	Code      MessageCode
}

type Message struct {
	MessageHead
	Raw  []byte
	Data interface{}
}

func Marshal(msg *Message) []byte {
	return nil
}

func Unmarshal([]byte) *Message {
	return nil
}
