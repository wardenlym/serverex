package game

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/lib/log"
)

// 尚未完全实现。。。

type Actor struct {
	g        *GameService
	old_user *odm.User
	user     *odm.User
	chara_id string
	chara    *odm.Character
}

func deepcopy(dst interface{}, src interface{}) {
	if dst == nil {
		panic(fmt.Errorf("dst cannot be nil"))
	}
	if src == nil {
		panic(fmt.Errorf("src cannot be nil"))
	}
	b, err := json.Marshal(src)
	if err != nil {
		panic(fmt.Errorf("Unable to marshal src: %s", err))
	}
	err = json.Unmarshal(b, dst)
	if err != nil {
		panic(fmt.Errorf("Unable to unmarshal into dst: %s", err))
	}
}

func (g *GameService) use_actor(uid int64) *Actor {
	old, e := g.db.GetUser(uid)
	if e != nil {
		log.Error(e)
		raise_error(msg.ErrInvalidUserId)
	}
	new := &odm.User{}
	deepcopy(new, old)
	return &Actor{g: g, old_user: old, user: new}
}

func (a *Actor) save() {
	if a.chara != nil {
		a.user.Characters[a.chara_id] = *a.chara
	}
	e := a.g.db.UpdateUser(a.user.Uid, a.user)
	if e != nil {
		log.Error("##**## DbSaveError:", e)
	}
}

func (a *Actor) user_info() *odm.User {
	return a.user
}

func (a *Actor) chara_info() *odm.Character {
	return a.chara
}

func (a *Actor) character(id string) *Actor {
	if _, exist := a.user.Characters[id]; exist {
		a.chara_id = id
		c := a.user.Characters[id]
		a.chara = &c
	} else {
		raise_error(msg.ErrInvalidCharacterId)
	}
	return a
}

func (a *Actor) next_character_id() string {
	last_id := a.user.CharacterIds[len(a.user.CharacterIds)-1]
	id, err := strconv.Atoi(last_id)
	if err != nil {
		raise_error(msg.ErrImpossiableLogic)
	}

	idstr := fmt.Sprint(id + 1)

	if _, exist := a.user.Characters[idstr]; exist {
		raise_error(msg.ErrImpossiableLogic)
	}
	return idstr
}

func (a *Actor) new_character(chara_id string, hero_id int) *Actor {

	if len(a.user.CharacterIds) >= 3 {
		return a
	}

	info, exist := csv.Table_HeroPrice[hero_id]
	if !exist {
		return a
	}

	a.cost_diamond(uint(info.Diamondprice))
	a.cost_gold(uint(info.Goldprice))

	a.user.CharacterIds = append(a.user.CharacterIds, chara_id)
	a.user.Characters[chara_id] = odm.NewCharacters(hero_id)

	a.user.ExploreInfo.Status = append(a.user.ExploreInfo.Status)
	return a
}
