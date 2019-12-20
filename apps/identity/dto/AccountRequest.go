package dto

type RegisterInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	PassWord    string `json:"passWord"`
	VerifyCode  string `json:"verifyCode"`
	NickName    string `json:"nickName"`
}

type PhoneVerificationInfo struct {
	PhoneNumber string `json:"phoneNumber"`
}

type GetTokenInfo struct {
	PassWord    string `json:"passWord"`
	PhoneNumber string `json:"phoneNumber"`
}

type VerifyPhoneInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	VerifyCode  string `json:"verifyCode"`
}

type ResetPasswordInfo struct {
	NewPassword string `json:"newPassword"`
}
