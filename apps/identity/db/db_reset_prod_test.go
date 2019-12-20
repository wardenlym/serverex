package db

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Test_prod_reset_database(t *testing.T) {

	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/identity_prod?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	db.DropTableIfExists(&UserInfo{})
	db.DropTableIfExists(&GuestAuthInfo{})
	db.DropTableIfExists(&PhoneAuthInfo{})

	db.AutoMigrate(
		&UserInfo{},
		&GuestAuthInfo{},
		&PhoneAuthInfo{},
	)
}
