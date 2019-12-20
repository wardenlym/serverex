package repo

import "github.com/globalsign/mgo"

func (db *DB) c_battlelog() *mgo.Collection {
	return db.session.DB("").C("battlelog")
}
