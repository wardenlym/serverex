package repo

import (
	"testing"

	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

// 评审包生成3个账号，并携带特定装备
func Test_add_test_user(t *testing.T) {
	db := NewDB()
	err := db.Open("mongodb://127.0.0.1:27017/excalibur_approval")
	if err != nil {
		t.Error(err)
	}

	user10000 := odm.NewUserWithApprovalCharacters(10000, 1)
	user10000.Diamond = 500

	user20000 := odm.NewUserWithApprovalCharacters(20000, 7)
	user20000.Diamond = 5000

	user30000 := odm.NewUserWithApprovalCharacters(30000, 12)
	user30000.Diamond = 50000

	err = db.UpsertUser(10000, user10000)
	if err != nil {
		t.Error(err)
	}

	err = db.UpsertUser(20000, user20000)
	if err != nil {
		t.Error(err)
	}

	err = db.UpsertUser(30000, user30000)
	if err != nil {
		t.Error(err)
	}
}
