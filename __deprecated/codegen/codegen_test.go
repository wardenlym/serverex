package protocol

import (
	"io/ioutil"
	"testing"
)

func Test_test1(t *testing.T) {
	b, e := ioutil.ReadFile("protocol.define")
	if e != nil {
		t.Error(e)
	}
	parse(b)
	// fmt.Println(parse(b))
}

func Test_main(t *testing.T) {

}
