package repo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/log"
)

func (db *DB) Ranking() *OpRanking {
	return &OpRanking{db}
}

func (db *DB) c_ranking_personal() *mgo.Collection {
	return db.session.DB("").C("ranking_personal")
}

func (db *DB) c_ranking_dungeon() *mgo.Collection {
	return db.session.DB("").C("ranking_dungeon")
}

type OpRanking struct {
	db *DB
}

func (p *OpRanking) InsertDungeonRanking(info odm.DungeonRankingInfo) {
	p.insert_dungeon_ranking(info)
}

func (p *OpRanking) GetDungeonRankingList(chapterid int, from int, count int) []odm.DungeonRankingInfo {
	return p.get_dungeon_ranking_range_by_time(chapterid, from, count)
}

func (p *OpRanking) insert_dungeon_ranking(info odm.DungeonRankingInfo) {
	err := p.db.c_ranking_dungeon().Insert(info)
	if err != nil {
		log.Error(err)
	}
}

func (p *OpRanking) get_dungeon_ranking_range_by_time(chapterid, from, count int) []odm.DungeonRankingInfo {

	itr := p.db.c_ranking_dungeon().Find(bson.M{"chapterId": chapterid}).Sort("time").Skip(from).Limit(count).Iter()

	entry := odm.DungeonRankingInfo{}
	rankings := []odm.DungeonRankingInfo{}
	k := from + 1
	for itr.Next(&entry) {
		entry.Rank = from + k
		k++
		rankings = append(rankings, entry)
	}
	return rankings
}

//
func (p *OpRanking) InsertPersonalRanking(info odm.PersonalRankingInfo) {

	old := odm.PersonalRankingInfo{}
	err := p.db.c_ranking_personal().Find(bson.M{"uid": info.Uid}).One(&old)
	if err != nil {
		info.ThreadId = old.ThreadId
		p.db.c_ranking_personal().Upsert(bson.M{"uid": info.Uid}, info)
		if err != nil {
			log.Error(err)
		}
	} else {
		p.db.c_ranking_personal().Upsert(bson.M{"uid": info.Uid}, info)
		if err != nil {
			log.Error(err)
		}
	}
}

func (p *OpRanking) GetPersonalRankingList(typestr string, from int, count int) []odm.PersonalRankingInfo {

	sort_by := "gold"

	switch typestr {
	case odm.PersonalRanking_Gold:
		sort_by = "-gold"
	case odm.PersonalRanking_Kill:
		sort_by = "-kill"
	case odm.PersonalRanking_Achievement:
		sort_by = "-achievement"
	}

	itr := p.db.c_ranking_personal().Find(nil).Sort(sort_by).Skip(from).Limit(count).Iter()

	entry := odm.PersonalRankingInfo{}
	rankings := []odm.PersonalRankingInfo{}
	k := from + 1
	for itr.Next(&entry) {
		entry.Rank = from + k
		k++
		rankings = append(rankings, entry)
	}
	return rankings
}
