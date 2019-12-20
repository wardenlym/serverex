package repo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"gitlab.com/megatech/serverex/apps/excalibur/payment"
	"gitlab.com/megatech/serverex/lib/log"
)

func (db *DB) Order() *OpOrder {
	return &OpOrder{db}
}

func (db *DB) c_apple_order() *mgo.Collection {
	return db.session.DB("").C("order_log_applepay")
}

func (db *DB) c_google_order() *mgo.Collection {
	return db.session.DB("").C("order_log_googlepay")
}

type OpOrder struct {
	db *DB
}

type orderinfo struct {
	orderid int64
	uid     int64
	typeid  string
}

func make_map_cause_bug1(v *payment.Receipt) map[string]interface{} {
	m := map[string]interface{}{
		"original_purchase_date_pst": v.A,
		"purchase_date_ms":           v.B,
		"unique_identifier":          v.C,
		"original_transaction_id":    v.D,
		"bvrs":                        v.E,
		"transaction_id":              v.TransactionID,
		"quantity":                    v.G,
		"unique_vendor_identifier":    v.H,
		"item_id":                     v.I,
		"version_external_identifier": v.J,
		"bid": v.K,
		"is_in_intro_offer_period":  v.L,
		"product_id":                v.ProductID,
		"purchase_date":             v.N,
		"is_trial_period":           v.O,
		"purchase_date_pst":         v.P,
		"original_purchase_date":    v.Q,
		"original_purchase_date_ms": v.R,
	}
	return m
}

func (p *OpOrder) InsertTempAppleOrderLog(userid int64, productid string, sandbox bool, transcation_id string, v *payment.Receipt) error {
	_, e := p.db.c_apple_order().Upsert(bson.M{"transcation_id": transcation_id}, map[string]interface{}{"uid": userid, "type": "applepay", "productId": productid, "sandbox": sandbox, "transcation_id": transcation_id, "data": make_map_cause_bug1(v)})
	if e != nil {
		log.Error(e, userid, productid, sandbox, transcation_id, v)
	}
	return e
}

func (p *OpOrder) IsAppleTranscationExist(transcation_id string) bool {

	old := map[string]interface{}{}
	e := p.db.c_apple_order().Find(bson.M{"transcation_id": transcation_id}).One(&old)
	if e != nil {
		return false
	}
	return true
}

func (p *OpOrder) InsertGooglePayOrderLog(userid int64, productid string, order_id string, v interface{}) error {
	_, e := p.db.c_google_order().Upsert(bson.M{"order_id": order_id}, map[string]interface{}{"uid": userid, "type": "googlepay", "productId": productid, "order_id": order_id, "data": v})
	if e != nil {
		log.Error(e, userid, productid, order_id, v)
	}
	return e
}
