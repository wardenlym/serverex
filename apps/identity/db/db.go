package db

import (
	"crypto/md5"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gitlab.com/megatech/serverex/lib/log"
)

type DB interface {
	CreateGuestUser(token string, uid int64, entry *UserInfo) error
	FindGuestUser(token string) (*UserInfo, error)
	FindUser(uid int64) (*UserInfo, error)
	BindPhone(phone string, uid int64, password string) error

	CreatePhoneUser(phone string, password string, uid int64, entry *UserInfo) error
	FindPhoneUser(phone string) (*UserInfo, error)
	ValidatePhoneUser(phone, password string) bool
	ModifyPassword(phone string, new_password string, uid int64) error
}

func New(dsn string) DB {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}
	return &sqldb{
		db: db,
	}
}

type sqldb struct {
	db *gorm.DB
}

func (p *sqldb) CreateGuestUser(token string, uid int64, entry *UserInfo) error {
	if err := p.db.Create(&GuestAuthInfo{
		Token: token,
		Uid:   uid,
	}).Error; err != nil {
		log.Error(err)
		return err
	}

	if err := p.db.Create(entry).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (p *sqldb) FindGuestUser(token string) (*UserInfo, error) {
	var auth GuestAuthInfo
	err := p.db.Where(&GuestAuthInfo{Token: token}).First(&auth).Error
	if err != nil {
		return nil, err
	}

	var user UserInfo
	err = p.db.Where(&UserInfo{Uid: auth.Uid}).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *sqldb) FindUser(uid int64) (*UserInfo, error) {

	var user UserInfo
	err := p.db.Where(&UserInfo{Uid: uid}).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *sqldb) BindPhone(phone string, uid int64, password string) error {

	if err := p.db.Create(&PhoneAuthInfo{
		Phone:        phone,
		Uid:          uid,
		PasswordHash: hash(uid, password),
	}).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func hash(uid int64, password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(uid)+"."+password)))
}

func (p *sqldb) CreatePhoneUser(phone string, password string, uid int64, entry *UserInfo) error {

	if err := p.db.Create(&PhoneAuthInfo{
		Phone:        phone,
		PasswordHash: hash(uid, password),
		Uid:          uid,
	}).Error; err != nil {
		log.Error(err)
		return err
	}

	if err := p.db.Create(entry).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (p *sqldb) FindPhoneUser(phone string) (*UserInfo, error) {
	var auth PhoneAuthInfo
	err := p.db.Where(&PhoneAuthInfo{Phone: phone}).First(&auth).Error
	if err != nil {
		return nil, err
	}

	var user UserInfo
	err = p.db.Where(&UserInfo{Uid: auth.Uid}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *sqldb) ValidatePhoneUser(phone, password string) bool {

	var auth PhoneAuthInfo
	err := p.db.Where(&PhoneAuthInfo{Phone: phone}).First(&auth).Error
	if err != nil {
		return false
	}

	var user UserInfo
	err = p.db.Where(&UserInfo{Uid: auth.Uid}).First(&user).Error
	if err != nil {
		return false
	}

	if auth.PasswordHash == hash(auth.Uid, password) {
		return true
	} else {
		return false
	}

}

func (p *sqldb) ModifyPassword(phone string, new_password string, uid int64) error {
	var auth PhoneAuthInfo
	err := p.db.Where(&PhoneAuthInfo{Phone: phone}).First(&auth).Error
	if err != nil {
		log.Error(err)
		return err
	}
	auth.PasswordHash = hash(uid, new_password)
	err = p.db.Model(&PhoneAuthInfo{}).Update(auth).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return err
}
