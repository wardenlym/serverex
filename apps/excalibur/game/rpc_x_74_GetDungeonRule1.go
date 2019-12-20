package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) GetDungeonRule1(req *msg.Request, arg *msg.GetDungeonRule1Request) *msg.GetDungeonRule1Response {

	bag_diff := []odm.LocatedCell{}

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)

		get_gift_typeid_n := func() (int, uint) {
			info, exist := csv.Table_DunRule_Gift[1]
			if !exist {
				raise_error(msg.ErrInvalidDungeonRule)
			}
			return info.Gift_id, uint(info.Gift_num)
		}

		award := []odm.Cell{cell_new_with_typeid_n(get_gift_typeid_n())}

		actor.put_to_bag(award)
		bag_diff = actor.diff_bag()
		actor.save()
	})

	if err != msg.Success {
		return &msg.GetDungeonRule1Response{Err: err}
	}

	return &msg.GetDungeonRule1Response{
		BagDiff: bag_diff,
	}
}
