package gate

import (
	"io"

	"gitlab.com/megatech/serverex/svc"
)

type Parser struct{}

func (parser *Parser) Decode(io.Reader) (error, svc.Atom, interface{}) {
	panic("not implemented")
}

func (parser *Parser) Encode(interface{}) []byte {
	panic("not implemented")
}
