package service

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"gitlab.com/megatech/serverex/apps/identity"
	"gitlab.com/megatech/serverex/apps/identity/db"
	"gitlab.com/megatech/serverex/apps/identity/dto"
)

func New(opt identity.Config) *service {
	return &service{
		opt:  opt,
		db:   db.New(opt.DataSourceName),
		code: identity.NewCodeVerify(),
	}
}

type service struct {
	db   db.DB
	opt  identity.Config
	code identity.CodeCache
}

func (s *service) SetGinRouter(r gin.IRouter) {

	guestaccount := r.Group("/guestAccount")
	guestaccount.GET("/register", s.GuestRegister)
	guestaccount.POST("/login", s.GuestLogin)
	guestaccount.POST("/logout", s.GuestLogout)
	guestaccount.POST("/bindphone", s.GuestBindPhone)

	account := r.Group("/account")
	account.POST("/register", s.AccountRegister)
	account.POST("/SendVerifyCode", s.AccountSendVerifyCode)
	account.POST("/token", s.AccountToken)
	account.POST("/login", s.AccountLogin)
	account.POST("/logout", s.AccountLogout)
	account.POST("/ResetPassToken", s.AccountResetPassToken)
	account.POST("/ResetPassword", s.AccountResetPassword)
}

func parseBody(c *gin.Context, v interface{}) error {
	b, err := c.GetRawData()
	if err != nil {
		return err
	}
	err = jsoniter.Unmarshal(b, v)
	if err != nil {
		return err
	}
	return nil
}

func ok_resp(c *gin.Context) {
	c.String(200, dto.ResponseBody{}.String())
}

func ok_resp_body(c *gin.Context, body dto.ResponseBody) {
	c.String(200, body.String())
}

func err_auth_resp(c *gin.Context) {
	c.Status(401)
}

func err_resp(c *gin.Context, code dto.StatusCode) {
	c.String(400, dto.ResponseBody{Code: code}.String())
}

func err_resp_body(c *gin.Context, body dto.ResponseBody) {
	c.String(400, body.String())
}
