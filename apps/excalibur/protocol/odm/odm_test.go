package odm

import (
	"fmt"
	"reflect"
	"testing"
	"unicode"

	jsoniter "github.com/json-iterator/go"
)

func uppercaseFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func lowercaseFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// 验证正反序列化ok
func Test_user_to_json(t *testing.T) {

	a := NewUserWith1CharacterDefault(123)
	s, err := jsoniter.MarshalToString(a)
	if err != nil {
		t.Error(err)
	}
	b := &User{}
	err = jsoniter.UnmarshalFromString(s, b)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(a, b) {
		t.Fail()
	}
	println(s)
	println()
	fmt.Printf("%#v", b)
}

// 验证正反序列化ok
func Test_odm_to_ident_json(t *testing.T) {

	user := NewUserWith1CharacterDefault(0)

	m := map[string]interface{}{
		"user": user,
	}

	b, err := jsoniter.MarshalIndent(m, "", "  ")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(b))
}
