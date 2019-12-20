package dto

import (
	"github.com/json-iterator/go"
)

type ResponseBody struct {
	Code    StatusCode  `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (p ResponseBody) String() string {
	j, _ := jsoniter.MarshalToString(p)
	return j
}

type StatusCode int

const (
	StatusOK                    StatusCode = 0
	ErrInvalidInput             StatusCode = 1000
	ErrGuestAccountNotExist     StatusCode = 1001
	ErrAccountNotGuest          StatusCode = 1002
	ErrInvalidGuestBindInfo     StatusCode = 1003
	ErrAccountNotExist          StatusCode = 1004
	ErrInvalidAccountOrPassword StatusCode = 1005
	ErrInvalidResetPassToken    StatusCode = 1006
	ErrDuplicateUserName        StatusCode = 1007
	ErrInvalidVerifyCode        StatusCode = 2001
	ErrSendVerifyCodeNotAllowed StatusCode = 2002
	ErrInvalidGoogleToken       StatusCode = 3001
	ErrInvalidFacebookToken     StatusCode = 4001
)
