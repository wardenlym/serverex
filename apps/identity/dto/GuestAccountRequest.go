package dto

type GuestLoginRequest struct {
	MachineCode string `json:"machineCode"`
}

type GuestBindRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	VerifyCode  string `json:"verifyCode"`
	PassWord    string `json:"passWord"`
}
