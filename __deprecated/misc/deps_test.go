package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/json-iterator/go"
	"github.com/urfave/cli"
)

// 因为golang会自动去掉不用的包导入,所以用这个文件来保持对常用包的依赖关系

func Test_deps_holder(t *testing.T) {
	r := gin.Default()
	fmt.Println(r)
	s, _ := jsoniter.MarshalToString(42)
	fmt.Println(s)
	a := cli.NewApp()
	fmt.Println(a)
	// https://qiita.com/rubytomato@github/items/4192e2340d81c0c39ca5
	session, err := mgo.DialWithTimeout("localhost:27017", time.Millisecond*100)
	if err != nil {
		t.Error(err)
	}
	defer session.Close()
}
