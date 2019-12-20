package repo

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/json-iterator/go"
	"gitlab.com/megatech/serverex/apps/excalibur/payment"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func Test_connectMongo(t *testing.T) {

	session, err := mgo.DialWithTimeout("localhost", time.Millisecond*100)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	type Person struct {
		Name  string
		Phone string
	}
	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)

}

func Test_DBGetUser(t *testing.T) {
	db := NewDB()
	err := db.Open("mongodb://localhost/abc_def")
	if err != nil {
		t.Error(err)
	}

	uid := int64(1234)
	user, err := db.GetUserOrUpsert(uid)
	if err != nil {
		t.Log(err.Error())
		t.Error(err)
	}

	user1234, err := db.GetUser(uid)
	if err != nil {
		t.Error(err)
	}

	a, _ := jsoniter.MarshalToString(user)
	b, _ := jsoniter.MarshalToString(user1234)
	if !reflect.DeepEqual(user, user1234) {
		t.Error(strings.Compare(a, b))
		t.Error("not equal")
		t.Error(user)
		t.Error(user1234)
	}
	db.DeleteUser(uid)
	db.Close()
}

func Test_GetSet(t *testing.T) {
	db := NewDB()
	err := db.Open("mongodb://localhost/test")
	if err != nil {
		t.Error(err)
	}

	uid := int64(1234)
	user, err := db.GetUserOrUpsert(uid)
	if err != nil {
		t.Log(err.Error())
		t.Error(err)
	}
	user.Nickname = "test name"
	db.UpsertUserField(uid, "nickname", user.Nickname)

	user1234, err := db.GetUser(uid)
	if err != nil {
		t.Error(err)
	}

	a, _ := jsoniter.MarshalToString(user)
	b, _ := jsoniter.MarshalToString(user1234)
	t.Log(a)
	t.Log(b)
	if !reflect.DeepEqual(user, user1234) {
		t.Error(strings.Compare(a, b))
		t.Error("not equal")
		t.Error(user)
		t.Error(user1234)
	}
	db.DeleteUser(uid)
	db.Close()
}

func Test_a(t *testing.T) {
	fmt.Println(string(1))
}
func Test_GetCharacter(t *testing.T) {
	db := NewDB()
	err := db.Open("mongodb://localhost/test")
	if err != nil {
		t.Error(err)
	}

	uid := int64(1234567)
	user, err := db.GetUserOrUpsert(uid)
	if err != nil {
		t.Log(err.Error())
		t.Error(err)
	}
	user.Nickname = "test name"
	db.UpsertUserField(uid, "nickname", user.Nickname)

	cid10000, err := db.GetCharacter(uid, user.CharacterIds[0])
	if err != nil {
		t.Error(err)
	}
	t.Log(fmt.Printf("%#v\n", cid10000))
	db.Close()
}

func Test_orderlog(t *testing.T) {
	db := NewDB()
	err := db.Open("mongodb://localhost/test")
	if err != nil {
		t.Error(err)
	}

	uid := int64(1234)
	db.Order().InsertTempAppleOrderLog(uid, "adfasdf", true, "adfasdfas", &payment.Receipt{})
	db.Close()
}

func Test_stage(t *testing.T) {

	fmt.Println(2015 / 1000 * 1000)
}

func Test_ls(t *testing.T) {

	a := func() func(int) int {
		n := 0
		f := func(arg int) int {
			return arg * n
		}
		n = 1
		return f
	}

	f := a()
	fmt.Println(f(1))
}

func Test_comment(t *testing.T) {
	db := NewDB()
	err := db.Open("mongodb://localhost/test")
	if err != nil {
		t.Error(err)
	}

	db.Comment().insert_new_comment(123, 234, "fuck", "whats up mother fucker.")

	db.Comment().insert_new_comment(456, 234, "fuck", "whats up mother fucker 2.")

	db.Close()
}

func Test_comment_range(t *testing.T) {
	db := NewDB()
	err := db.Open("mongodb://localhost/test")
	if err != nil {
		t.Error(err)
	}

	a := db.Comment().get_comment_range_by_parise(123, 0, 3)

	fmt.Println(a)
	db.Close()
}

func Test_comment_range_2(t *testing.T) {
	db := NewDB()
	err := db.Open("mongodb://localhost/test")
	if err != nil {
		t.Error(err)
	}

	a := db.Comment().get_comment_range_by_time(123, 0, 10)

	fmt.Println(a)
	db.Close()
}

func Test_comment_praise_comment(t *testing.T) {
	db := NewDB()
	err := db.Open("mongodb://localhost/test")
	if err != nil {
		t.Error(err)
	}

	db.Comment().praise_comment(65343627865362432)

	db.Close()
}

func Test_userp_get_set(t *testing.T) {
	db := NewDB()
	err := db.Open("mongodb://localhost/test")
	if err != nil {
		t.Error(err)
	}

	db.UserProperties().Set(456, "testkey1", "aaaaa")
	db.UserProperties().Set(456, "testkey2", "bbbbbbbb")

	a := db.UserProperties().Get(456, "testkey1")

	fmt.Println(a)

	a = db.UserProperties().Get(456, "testkey2")

	fmt.Println(a)

	a = db.UserProperties().Get(456, "testkey3")

	fmt.Println(a)

	db.Close()
}
