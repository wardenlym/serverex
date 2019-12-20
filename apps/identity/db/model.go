package db

import (
	"github.com/jinzhu/gorm"
)

type UserInfo struct {
	gorm.Model
	Uid         int64 `gorm:"unique;unique_index;not null"`
	Nickname    string
	AppleCode   string
	MachineCode string
}

type GuestAuthInfo struct {
	gorm.Model
	Token string `gorm:"unique;index;not null"`
	Uid   int64
}

type PhoneAuthInfo struct {
	gorm.Model
	Phone        string `gorm:"unique;index;not null"`
	PasswordHash string
	Uid          int64
}
