package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
)

func try_do(f func()) (retcode msg.ErrCode) {
	retcode = msg.Success
	defer func() {
		if err := recover(); err != nil {
			switch errcode := err.(type) {
			case msg.ErrCode:
				retcode = errcode
			default:
				// 重新抛出其他异常
				panic(err)
			}
		}
	}()

	f()

	return retcode
}

func raise_error(e msg.ErrCode) {
	panic(e)
}
