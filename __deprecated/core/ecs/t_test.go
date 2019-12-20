package ecs

import (
	"fmt"
	"reflect"
	"testing"
)

type A interface{}

type B struct {
	A
}

func Test_reflect(t *testing.T) {
	var a A = B{}
	var b A = &B{}

	fmt.Println(TypeOf(a))
	fmt.Println(TypeOf(b))
}

type AAA struct {
	A int
}

type BBB struct {
	A int
}

type CCC struct {
	AAA
	BBB
	A int
}

func Test_typecast(t *testing.T) {
	c := CCC{
		AAA{1},
		BBB{2},
		3,
	}

	fmt.Println(c.A)
}

func isInheritComponet(v interface{}) bool {
	var t reflect.Type
	if T := reflect.TypeOf(v); T.Kind() == reflect.Ptr {
		t = T.Elem()
	} else {
		t = T
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous && f.Type == reflect.TypeOf(Componet{}) {
			return true
		}
	}
	return false
}

func Test_isinherit(t *testing.T) {

	type EEEE struct {
		Componet
	}

	a := &EEEE{}
	b := EEEE{}
	var c interface{} = b

	fmt.Println(isInheritComponet(a))
	fmt.Println(isInheritComponet(b))

	fmt.Println(isInheritComponet(c))
}

func Test_maxuint(t *testing.T) {
	a := uint32(0)
	fmt.Println(a)

	fmt.Println(a - 1)
	var b uint32 = 4294967295
	fmt.Println(b)
	fmt.Println(b + 1)

}
