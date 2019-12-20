package fn

import (
	"fmt"
	"reflect"
)

func Apply(fn interface{}, v ...interface{}) {
	args := []reflect.Value{}
	for _, i := range v {
		args = append(args, reflect.ValueOf(i))
	}
	reflect.ValueOf(fn).Call(args)
}

type ActionT struct {
	fn    reflect.Value
	argsT []reflect.Type
	argsV []reflect.Value
}

func Action(f interface{}) *ActionT {
	fnv := reflect.ValueOf(f)
	fnt := reflect.TypeOf(f)

	if fnt.Kind() != reflect.Func {
		panic(fmt.Sprintf("%s=%v", fnt.Kind(), f))
	}

	t := reflect.ValueOf(f).Type()
	tt := []reflect.Type{}
	for i := 0; i < t.NumIn(); i++ {
		tt = append(tt, t.In(i))
	}

	return &ActionT{
		fn:    fnv,
		argsT: tt,
		argsV: []reflect.Value{},
	}
}

func (a *ActionT) Bind(args ...interface{}) {
	for _, i := range args {
		n := len(a.argsV)
		if len(a.argsT) < n {
			return
		}

		t, v := reflect.TypeOf(i), reflect.ValueOf(i)
		if t != a.argsT[n] {
			panic(fmt.Sprintf("%s %s=%v", t.String(), a.argsT[n].String(), i))
		}

		a.argsV = append(a.argsV, v)
	}
}

func (a *ActionT) Invoke() {
	a.fn.Call(a.argsV)
}
