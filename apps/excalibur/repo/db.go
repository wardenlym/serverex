package repo

import (
	"time"

	"github.com/globalsign/mgo"
	"gitlab.com/megatech/serverex/lib/log"
)

type DB struct {
	session *mgo.Session
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) Open(connstr string) error {
	var err error
	if db.session != nil {
		db.session.Close()
	}
	db.session, err = mgo.DialWithTimeout(connstr, time.Second*2)
	if err != nil {
		return err
	}
	log.Info("open mongodb:", connstr)
	db.session.SetMode(mgo.Monotonic, true)
	return nil
}

func (db *DB) Close() {
	if db.session != nil {
		db.session.Close()
		db.session = nil
	}
}
