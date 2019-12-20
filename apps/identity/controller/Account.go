package service

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/megatech/serverex/apps/identity"
	"gitlab.com/megatech/serverex/apps/identity/db"
	"gitlab.com/megatech/serverex/apps/identity/dto"
	"gitlab.com/megatech/serverex/lib/log"
	"gitlab.com/megatech/serverex/lib/util"
)

func (s *service) AccountRegister(c *gin.Context) {
	uid := util.NewSnowflakeID()

	user := &db.UserInfo{
		Uid:      uid,
		Nickname: "phone." + util.Int64ToBase58(uid),
	}

	var req dto.RegisterInfo
	if err := parseBody(c, &req); err != nil {
		err_resp(c, dto.ErrInvalidInput)
		return
	}

	if !s.code.Validate(req.PhoneNumber, req.VerifyCode) {
		err_resp(c, dto.ErrInvalidVerifyCode)
		return
	}

	if err := s.db.CreatePhoneUser(req.PhoneNumber, req.PassWord, uid, user); err != nil {
		log.Error(err)
		err_resp(c, dto.ErrDuplicateUserName)
		return
	}

	ok_resp(c)
}

func (s *service) AccountSendVerifyCode(c *gin.Context) {
	var req dto.PhoneVerificationInfo
	if err := parseBody(c, &req); err != nil {
		err_resp(c, dto.ErrInvalidInput)
		return
	}
	verifyCode := s.code.Add(req.PhoneNumber)
	identity.SendSMS(req.PhoneNumber, s.opt.AliSMSSignName, s.opt.AliSMSTemplateCode, verifyCode)
	ok_resp(c)
}

func (s *service) AccountToken(c *gin.Context) {
	var req dto.GetTokenInfo
	if err := parseBody(c, &req); err != nil {
		err_resp(c, dto.ErrInvalidInput)
		return
	}

	user, err := s.db.FindPhoneUser(req.PhoneNumber)
	if err != nil {
		err_resp(c, dto.ErrAccountNotExist)
		return
	}

	if !s.db.ValidatePhoneUser(req.PhoneNumber, req.PassWord) {
		err_resp(c, dto.ErrInvalidAccountOrPassword)
		return
	}

	uid := user.Uid

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

	ok_resp_body(c, dto.ResponseBody{
		Data: dto.CustomJWTRespnose{
			AccessToken: jwt_token,
			ExpiresIn:   60 * 24 * 7 * 365 * 10,
			TokenType:   "Bearer"}})
}

func (s *service) AccountLogin(c *gin.Context) {
	err, uid := identity.DecodeJWT(c.GetHeader("Authorization"), s.opt.JWTOption.SecurityKey)
	if err != nil {
		err_auth_resp(c)
		return
	}
	if user, err := s.db.FindUser(uid); err != nil {
		err_resp(c, dto.ErrAccountNotExist)
		return
	} else {
		ok_resp_body(c,
			dto.ResponseBody{
				Data: dto.LoginResponse{
					UserId:   fmt.Sprint(user.Uid),
					Nickname: user.Nickname}})
	}
}

func (s *service) AccountLogout(c *gin.Context) {
	err, uid := identity.DecodeJWT(c.GetHeader("Authorization"), s.opt.JWTOption.SecurityKey)
	if err != nil {
		err_auth_resp(c)
		return
	}
	if _, err := s.db.FindUser(uid); err != nil {
		err_resp(c, dto.ErrAccountNotExist)
		return
	} else {
		ok_resp(c)
	}
}

func (s *service) AccountResetPassToken(c *gin.Context) {
	var req dto.VerifyPhoneInfo
	if err := parseBody(c, &req); err != nil {
		err_resp(c, dto.ErrInvalidInput)
		return
	}

	user, err := s.db.FindPhoneUser(req.PhoneNumber)
	if err != nil {
		err_resp(c, dto.ErrAccountNotExist)
		return
	}

	if !s.code.Validate(req.PhoneNumber, req.VerifyCode) {
		err_resp(c, dto.ErrInvalidVerifyCode)
		return
	}

	uid := user.Uid

	jwt_token, err := identity.CreateJWT(map[string]interface{}{
		"resetUserId": fmt.Sprint(uid),
		"phone":       req.PhoneNumber,
		"iss":         "MegaIdentityService",
		"aud":         "MegaIdentityService",
		"exp":         time.Now().Add(time.Millisecond * 10).Unix(),
	}, s.opt.JWTOption.SecurityKey)

	if err != nil {
		log.Error(err)
		err_resp(c, dto.ErrInvalidInput)
		return
	}

	ok_resp_body(c, dto.ResponseBody{
		Data: dto.CustomJWTRespnose{
			AccessToken: jwt_token,
			ExpiresIn:   60 * 24 * 7 * 365 * 10,
			TokenType:   "Bearer"}})

}

func (s *service) AccountResetPassword(c *gin.Context) {
	var req dto.ResetPasswordInfo
	if err := parseBody(c, &req); err != nil {
		err_resp(c, dto.ErrInvalidInput)
		return
	}

	err, uid, phone := identity.DecodeJWTResetPassword(c.GetHeader("Authorization"), s.opt.JWTOption.SecurityKey)
	if err != nil {
		err_auth_resp(c)
		return
	}

	user, err := s.db.FindPhoneUser(phone)
	if err != nil || uid != user.Uid {
		err_resp(c, dto.ErrAccountNotExist)
		return
	}

	err = s.db.ModifyPassword(phone, req.NewPassword, uid)
	if err != nil {
		c.Status(500)
		return
	}
	ok_resp(c)
}
