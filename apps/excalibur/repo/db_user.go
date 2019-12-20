package repo

import (
	"errors"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/log"
)

func (db *DB) c_user() *mgo.Collection {
	return db.session.DB("").C("user")
}

func (db *DB) GetUser(uid int64) (*odm.User, error) {
	user := odm.User{}
	err := db.c_user().Find(bson.M{"uid": uid}).One(&user)
	// log.Warn("find user", user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &user, nil
}

func (db *DB) GetUserOrUpsert(uid int64) (*odm.User, error) {

	user, err := db.GetUser(uid)
	if user != nil {
		// log.Warn(user.Characters["1"].BattleAttribute)
	}
	if err == mgo.ErrNotFound {
		log.Error(err)
		newUser := odm.NewUserWith1CharacterDefault(uid)
		now := time.Now()
		t := now.Format(time.RFC3339)
		newUser.CreatedAt = t
		newUser.UpdatedAt = t
		newUser.CreateAtUnix = now.Unix()
		newUser.UpdatedAtUnix = now.Unix()

		info, err := db.c_user().Upsert(bson.M{"uid": uid}, newUser)
		log.Debug(err, info)
		if err != nil {
			return nil, err
		}
		return newUser, nil
	}
	return user, err
}

func (db *DB) CreateNewUserIfNotExist(uid int64) error {

	_, err := db.GetUser(uid)
	if err == mgo.ErrNotFound {
		newUser := odm.NewUserWith1CharacterDefault(uid)
		now := time.Now()
		t := now.Format(time.RFC3339)
		newUser.CreatedAt = t
		newUser.UpdatedAt = t
		newUser.CreateAtUnix = now.Unix()
		newUser.UpdatedAtUnix = now.Unix()
		info, err := db.c_user().Upsert(bson.M{"uid": uid}, newUser)
		log.Debug(err, info)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *DB) UpsertUser(uid int64, user *odm.User) error {

	t := time.Now().Format(time.RFC3339)
	user.CreatedAt = t
	user.UpdatedAt = t
	info, err := db.c_user().Upsert(bson.M{"uid": uid}, user)
	if err != nil {
		log.Error(err, info)
		return err
	}
	return nil
}

func (db *DB) UpdateUser(uid int64, user *odm.User) error {

	now := time.Now()
	t := now.Format(time.RFC3339)
	user.UpdatedAt = t
	user.UpdatedAtUnix = now.Unix()

	err := db.c_user().Update(bson.M{"uid": uid}, user)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// GetUserInfo_OrInitIfNotExit 不成功时会返回nil ? error?
func (db *DB) DeleteUser(uid int64) error {

	err := db.c_user().Remove(bson.M{"uid": uid})
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// 根据uid和字段名更新某个field，如果不存在则插入
func (db *DB) UpsertUserField(uid int64, field string, v interface{}) error {

	info, err := db.c_user().Upsert(
		bson.M{"uid": uid},
		bson.M{"$set": bson.M{field: v, "updatedAt": time.Now().Format(time.RFC3339)}})
	if err != nil {
		log.Error(err, info)
		return err
	}
	return nil
}

// 根据uid和字段名获取某个field
func (db *DB) QueryUserField(uid int64, field string, v interface{}) error {
	err := db.c_user().Find(bson.M{"uid": uid}).Select(bson.M{field: 1}).One(v)
	if err != nil {
		log.Error(uid, field, err)
		return err
	}
	return nil
}

func (db *DB) GetCharacter(uid int64, cid string) (*odm.Character, error) {

	user := odm.User{}
	err := db.c_user().Find(bson.M{"uid": uid}).One(&user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	//log.Warn("find user", user)
	character, ok := user.Characters[cid]
	if !ok {
		return nil, errors.New("character id not exist.")
	}
	return &character, nil
}
