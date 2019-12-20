package util

import (
	"fmt"
	"testing"
)

type Abc struct {
	A     int
	Bcd   string
	Def_a []int
	G_Fkj []string
	Xyzkj map[string]Bcdef `json:"aaa"`
}

type Bcdef struct {
}

func Test_a(t *testing.T) {
	s := MakeCSharpNamespace("", "nstest",
		ParseStructToClass(Abc{}),
		ParseStructToClass(Bcdef{}),
	)
	fmt.Println(s)
}
