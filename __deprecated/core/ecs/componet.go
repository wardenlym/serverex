package ecs

import "reflect"

type Componet struct{}
type ComponetType reflect.Type

func TypeOf(v interface{}) ComponetType {
	return reflect.TypeOf(v)
}

func IsComponetType(v interface{}) bool {
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
