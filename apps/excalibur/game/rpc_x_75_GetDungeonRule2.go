package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func check_ghost_ring_type() bool {
	return false
}

func (g *GameService) GetDungeonRule2(req *msg.Request, arg *msg.GetDungeonRule2Request) *msg.GetDungeonRule2Response {

	bag_diff := []odm.LocatedCell{}
	equipments_diff := []odm.LocatedCell{}

	err := try_do(func() {
		actor := g.use_actor(req.Uid).character(arg.CharacterId)

		try_make_new_ghost_ring := func(c odm.Cell, current_id int, new_id int) odm.Cell {

			if c.Item == nil || c.Item.TypeId != current_id || c.Num != 1 {
				raise_error(msg.ErrInvalidDungeonRule)
			}
			c.Item.TypeId = new_id
			return c
		}

		get_ghost_ring_typeid_by_level := func(level int) (int, int) {

			next_level := level + 1
			id1 := -1
			id2 := -1

			if info, exist := csv.Table_DunRule_GhostWd[level]; exist {
				id1 = info.Ring_id
			}

			if info, exist := csv.Table_DunRule_GhostWd[next_level]; exist {
				id2 = info.Ring_id
			}

			if id2 == -1 {
				raise_error(msg.ErrInvalidDungeonRule)
			}

			return id1, id2
		}

		current, next := get_ghost_ring_typeid_by_level(arg.Level)

		if arg.Level == 0 {
			c := odm.NewCell(next, 1)
			actor.put_to_bag([]odm.Cell{c})

		} else {
			c := odm.Cell{}
			if arg.IsEquiped {
				c = actor.chara_info().Bag.Equipments[arg.CellIndex]
			} else {
				c = actor.chara_info().Bag.Cells[arg.CellIndex]
			}

			c = try_make_new_ghost_ring(c, current, next)

			if arg.IsEquiped {
				actor.chara_info().Bag.Equipments[arg.CellIndex] = c
			} else {
				actor.chara_info().Bag.Cells[arg.CellIndex] = c
			}
		}

		bag_diff = actor.diff_bag()
		equipments_diff = actor.diff_equipments()
		actor.save()
	})

	if err != msg.Success {
		return &msg.GetDungeonRule2Response{Err: err}
	}

	return &msg.GetDungeonRule2Response{
		BagDiff:        bag_diff,
		EquipmentsDiff: equipments_diff,
	}
}
