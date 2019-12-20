package repo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"gitlab.com/megatech/serverex/lib/util"
)

func (db *DB) c_auth() *mgo.Collection {
	return db.session.DB("").C("account")
}

type Auth struct {
	LoginName  string `json:"loginName" bson:"loginName"`
	LoginToken string `json:"loginToken" bson:"loginToken"`
	Uid        int64  `json:"uid" bson:"uid"`
}

func (db *DB) FindInAuth(loginName, loginToken, appToken string) (int64, error) {
	// appToken 之后对应多个app

	var auth Auth
	err := db.c_auth().Find(bson.M{"loginName": loginName, "loginToken": loginToken}).One(&auth)
	if err != nil {
		return 0, err
	}
	return auth.Uid, nil
}

func (db *DB) InsertIntoAuth(loginName, loginToken string) (int64, error) {
	auth := Auth{
		LoginName:  loginName,
		LoginToken: loginToken,
		Uid:        util.NewSnowflakeID(),
	}
	err := db.c_auth().Insert(auth)
	if err != nil {
		return 0, err
	}
	return auth.Uid, nil
}
