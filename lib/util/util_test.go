package util

import (
	"fmt"
	"testing"
	"time"

	"github.com/bwmarrin/snowflake"
)

func Test_snowflakeID(t *testing.T) {
	now := time.Now().UnixNano() / 1000000
	fmt.Println("now: ", now)
	id := int64(NewSnowflakeID())
	fmt.Println(id)
	fmt.Println(Int64ToBase58(id))
	s := snowflake.ID(id)
	fmt.Println(s.Base2())
	fmt.Println(36, s.Base36())
	fmt.Println(58, s.Base58())
	fmt.Println(64, s.Base64())
	fmt.Println(s.String())
	fmt.Println(time.Unix(s.Time()/1000, 0), s.Node(), s.Step())
	b, _ := s.MarshalJSON()
	fmt.Println(string(b))

}

func Test_snowflakeID1(t *testing.T) {

	id := int64(NewSnowflakeID())
	fmt.Println(id)

	a := Int64ToBase36(id)
	fmt.Println(a)
	b := Base36ToInt64(a)
	fmt.Println(b)
}

func Test_snowflakeID2(t *testing.T) {
	for i := 1; i < 10; i++ {
		id := NewSnowflakeID()

		fmt.Println(i, id, Int64ToBase58(id))
	}
}

func Test_bbb(t *testing.T) {
	println(hostname)
	println(Pid())

}

func Test_ccc(t *testing.T) {
	println(RandomIn(1, 3))
	println(RandomIn(1, 3))
	println(RandomIn(1, 3))
	println(RandomIn(1, 3))
	println(RandomIn(1, 3))
	println(RandomIn(1, 3))
	println(RandomIn(1, 3))
	println(RandomIn(1, 3))
	println(RandomIn(1, 3))
	println(RandomIn(1, 3))
}

func Test_ddd(t *testing.T) {
	for i := 0; i < 100; i++ {
		println(RandomIn(1, 5))
		println(RandomIn(1, 5))
	}
}

func Test_eee(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandomIndexList(10, 5))
	}

	for i := 0; i < 10; i++ {
		fmt.Println(RandomIndexList(5, 3))
	}
}
