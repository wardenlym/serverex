package game

import (
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
)

func (g *GameService) PickupLoot(req *msg.Request, arg *msg.PickupLootRequest) *msg.PickupLootResponse {
	user := g.get_user(req.Uid)
	chara := user.Characters[arg.CharacterId]

	topick := []odm.Cell{}
	for _, index := range arg.From {
		if index >= 0 && index < len(chara.DungeonStatus.UnpickedLoot) &&
			chara.DungeonStatus.UnpickedLoot[index].Item != nil {
			topick = append(topick, chara.DungeonStatus.UnpickedLoot[index])
		}
	}
	newcells, _ := cell_move_from_to(topick, chara.Bag.Cells)

	bag_diff := cell_diff(chara.Bag.Cells, newcells)

	chara.Bag.Cells = newcells
	chara.DungeonStatus.UnpickedLoot = []odm.Cell{}
	user.Characters[arg.CharacterId] = chara
	g.save_user(user)
	return &msg.PickupLootResponse{
		BagDiff: bag_diff,
	}
}
