package repo

import (
	"github.com/globalsign/mgo"
)

func (db *DB) OperationActivity() *OpOperationActivity {
	return &OpOperationActivity{db}
}

func (db *DB) c_operation_activity() *mgo.Collection {
	return db.session.DB("").C("operation_activity")
}

type OpOperationActivity struct {
	db *DB
}

type OperationActivity struct {
}

func (p *OperationActivity) Get(userid int64) OperationActivity {
	return OperationActivity{}
}

func (p *OperationActivity) Set(userid int64, record OperationActivity) {
}
