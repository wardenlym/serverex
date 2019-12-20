package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/megatech/serverex/apps/identity"
	"gitlab.com/megatech/serverex/apps/identity/db"
	"gitlab.com/megatech/serverex/apps/identity/dto"
	"gitlab.com/megatech/serverex/lib/log"
	"gitlab.com/megatech/serverex/lib/util"
)

func (s *service) GuestRegister(c *gin.Context) {
	uid := util.NewSnowflakeID()

	user := &db.UserInfo{
		Uid:      uid,
		Nickname: "guest." + util.Int64ToBase58(uid),
	}

	jwt_token, err := identity.CreateJWT(map[string]interface{}{
		"userId": fmt.Sprint(uid),
		"iss":    "MegaIdentityService",
		"aud":    "ExcaliburClient",
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	}, s.opt.JWTOption.SecurityKey)

	if err != nil {
		log.Error(err)
		err_resp(c, dto.ErrInvalidInput)
		return
	}

	err = s.db.CreateGuestUser(jwt_token, uid, user)
	if err != nil {
		log.Error(err)
		c.Status(500)
		return
	}

	ok_resp_body(c, dto.ResponseBody{
		Data: dto.CustomJWTRespnose{
			AccessToken: jwt_token,
			ExpiresIn:   60 * 24 * 7 * 365 * 10,
			TokenType:   "Bearer"}})
}

func (s *service) GuestLogin(c *gin.Context) {

	err, uid := identity.DecodeJWT(c.GetHeader("Authorization"), s.opt.JWTOption.SecurityKey)
	if err != nil {
		err_auth_resp(c)
		return
	}

	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	user, err := s.db.FindGuestUser(token)
	if err != nil {
		err_resp(c, dto.ErrGuestAccountNotExist)
		return
	}

	ok_resp_body(c,
		dto.ResponseBody{
			Data: dto.LoginResponse{
				UserId:   fmt.Sprint(uid),
				Nickname: user.Nickname}})
}

func (s *service) GuestLogout(c *gin.Context) {

	err, _ := identity.DecodeJWT(c.GetHeader("Authorization"), s.opt.JWTOption.SecurityKey)

	if err != nil {
		err_auth_resp(c)
		return
	}

	ok_resp(c)
}

func (s *service) GuestBindPhone(c *gin.Context) {
	err, uid := identity.DecodeJWT(c.GetHeader("Authorization"), s.opt.JWTOption.SecurityKey)

	if err != nil {
		err_auth_resp(c)
		return
	}

	var req dto.GuestBindRequest
	err = parseBody(c, &req)
	if err != nil {
		err_resp(c, dto.ErrInvalidInput)
		return
	}

	_, err = s.db.FindUser(uid)
	if err != nil {
		err_resp(c, dto.ErrAccountNotExist)
		return
	}

	if !s.code.Validate(req.PhoneNumber, req.VerifyCode) {
		err_resp(c, dto.ErrInvalidVerifyCode)
		return
	}

	err = s.db.BindPhone(req.PhoneNumber, uid, req.PassWord)
	if err != nil {
		err_resp(c, dto.ErrDuplicateUserName)
		return
	}

	ok_resp(c)
}
