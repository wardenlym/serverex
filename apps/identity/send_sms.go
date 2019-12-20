package identity

import (
	"gitlab.com/megatech/serverex/lib/log"
	sms "gitlab.com/megatech/sms-proxy/aliyun-sms-sdk"
)

const (
	ACCESSID  = "LTAIgsAzU0sNKxKi"
	ACCESSKEY = "0yZchD84SQFnc3gKd4DkKenWUYawTb"
)

func smsClient() *sms.Client {
	return sms.New(ACCESSID, ACCESSKEY)
}

func SendSMS(phoneNumber, signName, templateCode, templateParam string) {
	msg, err := smsClient().Send(phoneNumber, signName, templateCode, `{"code":`+templateParam+`}`)
	if err != nil {
		log.Error(err, msg.Message)
	}
}
