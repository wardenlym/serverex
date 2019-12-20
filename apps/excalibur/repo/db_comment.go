package repo

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/log"
	"gitlab.com/megatech/serverex/lib/util"
)

func (db *DB) Comment() *OpComment {
	return &OpComment{db}
}

func (db *DB) c_comments() *mgo.Collection {
	return db.session.DB("").C("comments")
}

type OpComment struct {
	db *DB
}

func (p *OpComment) PostComment(threadid, uid int64, nickname, content string) {
	p.insert_new_comment(threadid, uid, nickname, content)
}

func (p *OpComment) PraiseComment(threadid, commentid int64) {
	p.praise_comment(commentid)
}

func (p *OpComment) GetThreadHotComment(threadid int64, topn int) []odm.Comment {
	from, count := 0, topn
	return p.get_comment_range_by_parise(threadid, from, count)
}

func (p *OpComment) GetThreadComment(threadid int64, from, count int) []odm.Comment {
	return p.get_comment_range_by_time(threadid, from, count)
}

func (p *OpComment) insert_new_comment(threadid, uid int64, nickname, content string) {
	comment := odm.Comment{
		CommentId: util.NewSnowflakeID(),
		ThreadId:  threadid,
		Uid:       uid,
		Date:      time.Now().Unix(),
		Nickname:  nickname,
		Content:   content,
	}
	err := p.db.c_comments().Insert(comment)
	if err != nil {
		log.Error(err)
	}
}

func (p *OpComment) praise_comment(commentid int64) {
	err := p.db.c_comments().Update(bson.M{"commentId": commentid}, bson.M{"$inc": bson.M{"praise": 1}})
	if err != nil {
		log.Error(err)
	}
}

func (p *OpComment) get_comment_range_by_parise(threadid int64, from, count int) []odm.Comment {
	itr := p.db.c_comments().Find(bson.M{"threadId": threadid}).Sort("-praise").Skip(from).Limit(count).Iter()

	entry := odm.Comment{}
	comments := []odm.Comment{}
	for itr.Next(&entry) {
		comments = append(comments, entry)
	}
	return comments
}

func (p *OpComment) get_comment_range_by_time(threadid int64, from, count int) []odm.Comment {
	itr := p.db.c_comments().Find(bson.M{"threadId": threadid}).Sort("-date").Skip(from).Limit(count).Iter()

	entry := odm.Comment{}
	comments := []odm.Comment{}
	for itr.Next(&entry) {
		comments = append(comments, entry)
	}
	return comments
}
