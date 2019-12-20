package repo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"gitlab.com/megatech/serverex/lib/log"
)

func (db *DB) UserProperties() *OpUserProperties {
	return &OpUserProperties{db}
}

func (db *DB) c_user_properties() *mgo.Collection {
	return db.session.DB("").C("user_properties")
}

type OpUserProperties struct {
	db *DB
}

func (p *OpUserProperties) Set(uid int64, k, v string) {
	_, err := p.db.c_user_properties().Upsert(bson.M{"uid": uid}, bson.M{"$set": bson.M{k: v}})
	if err != nil {
		log.Error(err)
	}
}

func (p *OpUserProperties) Get(uid int64, k string) string {
	doc := map[string]interface{}{}
	err := p.db.c_user_properties().Find(bson.M{"uid": uid}).Select(bson.M{k: 1}).One(&doc)
	if err != nil {
		log.Error(err)
		return ""
	}

	s, exist := doc[k]
	if !exist {
		return ""
	}

	ret, ok := s.(string)
	if !ok {
		return ""
	}

	return ret
}
