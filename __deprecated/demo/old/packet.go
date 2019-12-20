package demo

import (
	"bytes"
	"io"
)

const (
	_ = byte(iota)
	PktMsg
	PktSyn
	PktAck
)

type Packet struct {
	// 2 byte head -> int16 len
	Code byte
	Data []byte
}

type Decoder interface {
	Decode(io.Reader, *bytes.Buffer) error
}

type Encoder interface {
	Encode(io.Writer, *bytes.Buffer) error
}

////////////////////////////////////////////////////////////////////////////////////

// 默认的 基于2字节int16长度头的协议
//  2     2    n
// len code data

func NewPacketDecoder() Decoder {
	return &lenPrefix{}
}

func NewPacketEncoder() Encoder {
	return &lenPrefix{}
}

type lenPrefix struct {
}

func (*lenPrefix) Decode(io.Reader, *bytes.Buffer) error {
	return nil
}

func (*lenPrefix) Encode(io.Writer, *bytes.Buffer) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////////

// 基于字符的 \r\n 分割协议 主要用于测试

// func NewTelnetDecoder() Decoder {
// 	return lineEnd{}
// }

// func NewTelnetEncoder() Encoder {
// 	return lineEnd{}
// }

// type lineEnd struct {
// }

// func (lineEnd) Decode(r io.Reader) (*Packet, error) {
// 	s, err := r.ReadString('\n')
// 	if err != nil {
// 		log.Error(err, s)
// 		return nil, err
// 	}
// 	s1 := strings.TrimRight(s, "\r\n")
// 	return &Packet{Code: 0, Data: []byte(s1)}, nil
// }

// func (lineEnd) Encode(d *Packet, w io.Writer) error {
// 	b := []byte(string(d.Data) + "\r\n")

// 	n, err := w.Write(b)
// 	if err != nil {
// 		log.Error(err, n, len(b))
// 		return err
// 	}
// 	return nil
// }
