package db

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Test_db1(t *testing.T) {
	//	"<user>:<password>/<database>?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

}

func Test_1(t *testing.T) {

	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

}

func Test_2(t *testing.T) {
	a := hash(123, "abcabc")
	fmt.Println(a)

}

func Test_23(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	phone := "15242492179"
	var uid int64 = 85287832687480832
	new_password := "111111"
	err = db.Model(&PhoneAuthInfo{}).Update(PhoneAuthInfo{Phone: phone, PasswordHash: hash(uid, new_password)}).Error
	fmt.Println(err)
}
