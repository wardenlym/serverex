package fn

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_a(t *testing.T) {
	b, err := ioutil.ReadFile("test.go")
	if err != nil {
		t.Error(err)
	}
	Compile(b)
}

func a1() {
	fmt.Println("a1", 1111)
}

type S struct {
	SS int
}

func a2(a int, b string, c S) {
	fmt.Println("a2", a, b, c.SS)
}

func a3() func(a int, b string, c S) {
	e := 1
	f := "abc"
	return func(a int, b string, c S) {
		fmt.Println("a3", a, b, c.SS, e, f)
	}
}

func Test_b(t *testing.T) {

	aa1 := Action(a1)
	aa2 := Action(a2)
	aa3 := Action(a3())

	aa2.Bind(2222, "abc2222")
	aa2.Bind(S{22221})

	aa3.Bind(3333)
	aa3.Bind("abc3333")
	aa3.Bind(S{33331})

	aa1.Invoke()
	aa2.Invoke()
	aa3.Invoke()

	Apply(a3(), 4444, "abc4444", S{44441})

}

func Test_c(t *testing.T) {

	a := []int{1}
	fmt.Println(a)
	func(aa []int) {
		aa = append(aa, 2)
		fmt.Println(aa)
	}(a)
	fmt.Println(a)

	func(aa *[]int) {
		*aa = append(*aa, 2)
		fmt.Println(*aa)
	}(&a)
	fmt.Println(a)
}
