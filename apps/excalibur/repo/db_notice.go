package repo

import (
	"errors"
	"sort"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/log"
)

func (db *DB) Notice() *OpNotice {
	return &OpNotice{db}
}

func (db *DB) c_notice() *mgo.Collection {
	return db.session.DB("").C("notice")
}

type OpNotice struct {
	db *DB
}

func add_cell(ids []int) []odm.Cell {
	m := map[int]uint{}
	for _, id := range ids {
		m[id]++
	}
	c := []odm.Cell{}
	for typeid, count := range m {
		c = append(c, odm.NewCell(typeid, count))
	}
	return c
}

func parse_time(s string) int64 {
	t, e := time.Parse(time.RFC3339, s+"+08:00")
	if e != nil {
		return time.Now().Unix()
	}
	return t.Unix()
}

func newNoticeList() []odm.NoticeInfo {

	list := []odm.NoticeInfo{}
	for _, v := range csv.Table_NoticeData {
		a := odm.NoticeInfo{
			Id:       v.Id,
			Title:    v.Title,
			Content:  v.Content,
			Gold:     uint(v.Gold),
			Diamond:  uint(v.Diamond),
			Gift:     add_cell(v.Gift),
			Sender:   v.Sender,
			SendTime: parse_time(v.Sendingtime),
		}
		list = append(list, a)
	}
	return list
}

func if_need_merge_newer_notice_list(old []odm.NoticeInfo) []odm.NoticeInfo {

	parse_time := func(s string) int64 {
		t, e := time.Parse(time.RFC3339, s+"+08:00")
		if e != nil {
			return time.Now().Unix()
		}
		return t.Unix()
	}

	oldest_id := 0
	for _, v := range old {
		if v.Id > oldest_id {
			oldest_id = v.Id
		}
	}

	for id, v := range csv.Table_NoticeData {
		if id > oldest_id {
			old = append(old, odm.NoticeInfo{
				Id:       v.Id,
				Title:    v.Title,
				Content:  v.Content,
				Gold:     uint(v.Gold),
				Diamond:  uint(v.Diamond),
				Gift:     add_cell(v.Gift),
				Sender:   v.Sender,
				SendTime: parse_time(v.Sendingtime),
			})
			log.Info(old)
		}
	}
	return old
}

func (p *OpNotice) GetNoticeList(uid int64) []odm.NoticeInfo {
	info := odm.UserNoticeInfo{
		NoticeInfos: []odm.NoticeInfo{},
	}
	err := p.db.c_notice().Find(bson.M{"uid": uid}).One(&info)
	if err != nil {
		log.Error(err)
		if err == mgo.ErrNotFound {
			info.NoticeInfos = newNoticeList()
			info.Uid = uid
			p.db.c_notice().Upsert(bson.M{"uid": uid}, info)
		}
	} else {
		info.NoticeInfos = if_need_merge_newer_notice_list(info.NoticeInfos)
		p.db.c_notice().Upsert(bson.M{"uid": uid}, info)
	}

	return sort_notice_info(info.NoticeInfos)
}

type sortable_notice []odm.NoticeInfo

func (c sortable_notice) Len() int {
	return len(c)
}

func (c sortable_notice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c sortable_notice) Less(i, j int) bool {

	calc_weight := func(v odm.NoticeInfo) int {
		k := 0
		if v.Deleted {
			k += 100000
		}
		if v.Confirmed {
			k += 10000
		}
		if v.Read {
			k += 1000
		}
		k += v.Id
		return k
	}

	weight_i := calc_weight(c[i])
	weight_j := calc_weight(c[j])
	if weight_i == weight_j {
		return c[i].SendTime > c[i].SendTime
	}
	return weight_i < weight_j
}
func sort_notice_info(l []odm.NoticeInfo) []odm.NoticeInfo {
	sort.Sort(sortable_notice(l))
	return l
}

func (p *OpNotice) GetNotice(uid int64, id int) (odm.NoticeInfo, error) {

	info := odm.UserNoticeInfo{
		NoticeInfos: []odm.NoticeInfo{},
	}
	err := p.db.c_notice().Find(bson.M{"uid": uid}).One(&info)
	if err != nil {
		log.Error(err)
		return odm.NoticeInfo{}, err
	}
	for _, v := range info.NoticeInfos {
		if v.Id == id {
			return v, nil
		}
	}

	return odm.NoticeInfo{}, errors.New("not found")
}

func (p *OpNotice) ConfirmNotice(uid int64, id int) {
	info := odm.UserNoticeInfo{
		NoticeInfos: []odm.NoticeInfo{},
	}
	err := p.db.c_notice().Find(bson.M{"uid": uid}).One(&info)
	if err != nil {
		log.Error(err)
		return
	}
	for i, v := range info.NoticeInfos {
		if v.Id == id {
			info.NoticeInfos[i].Confirmed = true
			info.NoticeInfos[i].Read = true
			_, err := p.db.c_notice().Upsert(bson.M{"uid": uid}, info)
			if err != nil {
				log.Error(err)
			}
			return
		}
	}
	log.Error("not found")
}

func (p *OpNotice) MarkAsReadNotice(uid int64, id int) {
	info := odm.UserNoticeInfo{
		NoticeInfos: []odm.NoticeInfo{},
	}
	err := p.db.c_notice().Find(bson.M{"uid": uid}).One(&info)
	if err != nil {
		log.Error(err)
		return
	}
	for i, v := range info.NoticeInfos {
		if v.Id == id {
			info.NoticeInfos[i].Read = true
			_, err := p.db.c_notice().Upsert(bson.M{"uid": uid}, info)
			if err != nil {
				log.Error(err)
			}
			return
		}
	}
	log.Error("not found")
}

func (p *OpNotice) DeleteNotice(uid int64, id int) {
	info := odm.UserNoticeInfo{
		NoticeInfos: []odm.NoticeInfo{},
	}
	err := p.db.c_notice().Find(bson.M{"uid": uid}).One(&info)
	if err != nil {
		log.Error(err)
		return
	}
	for i, v := range info.NoticeInfos {
		if v.Id == id {
			// remove i
			info.NoticeInfos[i].Deleted = true
			_, err := p.db.c_notice().Upsert(bson.M{"uid": uid}, info)
			if err != nil {
				log.Error(err)
			}
			return
		}
	}
	log.Error("not found")
}
