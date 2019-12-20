package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"gitlab.com/megatech/golib/log"
	sms "gitlab.com/megatech/sms-proxy/aliyun-sms-sdk"
)

const (
	ACCESSID  = "your_accessid"
	ACCESSKEY = "your_accesskey"
)

func sms_client() *sms.Client {
	return sms.New(ACCESSID, ACCESSKEY)
}

func main() {

	var address string

	flag.StringVar(&address, "l", "localhost:7000", "listen address IP:Port")

	flag.Parse()

	log.Info("sms-proxy start.")

	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {})
	r.POST("/sms/send", func(c *gin.Context) {
		msg, err := sms_client().Send("1380000****", "阿里云短信测试专用", "SMS_71390007", `{"code":"123456"}`)
		if err != nil {
			log.Error(err, msg)
		}
	})

	err := r.Run(address)
	if err != nil {
		log.Warn(err)
	}
	log.Info("sms-proxy stop.")
}
