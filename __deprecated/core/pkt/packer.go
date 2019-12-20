package pkt

import (
	"encoding/binary"
	"io"
)

func NewLengthPrefixPacker() EncoderDecoder {
	return &lengthPrefixPacker{}
}

//   2B    2B   < 2^16-1=65535B
// Code   Len   Data
type lengthPrefixPacker struct{}

func (*lengthPrefixPacker) Decode(r io.Reader) (*Pkt, error) {

	head := make([]byte, 4)
	_, e := r.Read(head)
	if e != nil {
		return nil, e
	}
	code := binary.LittleEndian.Uint16(head[0:2])
	len := binary.LittleEndian.Uint16(head[2:4])
	data := make([]byte, len)

	_, e = r.Read(data)
	if e != nil {
		return nil, e
	}
	return &Pkt{
		Code: CodeType(code),
		Data: data,
	}, nil
}

const maxLen = 2 ^ 16 - 2 - 2

func (*lengthPrefixPacker) Encode(p *Pkt, w io.Writer) error {

	var length int
	if len(p.Data) < maxLen {
		length = len(p.Data)
	} else {
		length = maxLen
	}

	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(p.Code))
	w.Write(buf)
	binary.LittleEndian.PutUint16(buf, uint16(length))
	w.Write(buf)
	w.Write(p.Data)
	return nil
}
