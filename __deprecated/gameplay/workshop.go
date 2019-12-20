package gameplay

import (
	"math/rand"
	"time"

	"github.com/json-iterator/go"
	"github.com/kyokomi/lottery"
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
	"gitlab.com/megatech/serverex/apps/excalibur/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/odm"
	"gitlab.com/megatech/serverex/lib/log"
	"gitlab.com/megatech/serverex/lib/util"
)

func (g *game) craft(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.CraftRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}

	//
	info, err := g.db.GetUser(req.Uid)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	if info.Diamond < o.CostDiamond {
		res.Data = &msg.CraftResponse{ErrMsg: "钻石不足"}
		return res
	}
	if info.Gold < o.CostGold {
		res.Data = &msg.CraftResponse{ErrMsg: "金币不足"}
		return res
	}

	gold := info.Gold - o.CostGold
	diamond := info.Diamond - o.CostDiamond

	ok, cells := checkMaterials(o.Materials, info.Characters[o.CharacterId].Bag.Cells)

	if !ok {
		res.Data = &msg.CraftResponse{ErrMsg: "材料不足"}
		return res
	}

	ok, index := checkCellEnough(cells)
	if !ok {
		res.Data = &msg.CraftResponse{ErrMsg: "背包空间不足"}
		return res
	}

	cells[index] = odm.Cell{
		Index: index,
		Item: &odm.Item{
			Guid:            util.NewSnowflakeID(),
			TypeId:          o.ToCraftTypeId,
			StarLevel:       0,
			StarUpgradeInfo: []int{},
		},
		Num: 1,
	}

	// 更新数据稍后做成一次性update的
	g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".bag.cells", cells)
	g.db.UpsertUserField(req.Uid, "gold", gold)
	g.db.UpsertUserField(req.Uid, "diamond", diamond)

	bag := info.Characters[o.CharacterId].Bag
	bag.Cells = cells
	res.Data = &msg.CraftResponse{
		IsSuccess:    true,
		CraftedItems: cells[index],
		Gold:         gold,
		Diamond:      diamond,
		BagStatus:    bag,
	}
	return res
}

func checkCellEnough(cells []odm.Cell) (bool, int) {
	for i, c := range cells {
		if c.Item == nil || c.Num == 0 {
			return true, i
		}
	}
	return false, -1
}

func checkMaterials(materials []msg.Material, cells []odm.Cell) (bool, []odm.Cell) {
	requiredCnt := map[int]uint{}
	haveCnt := map[int]uint{}
	new_cells := make([]odm.Cell, len(cells))

	for _, material := range materials {
		requiredCnt[material.TypeId] = material.Num
		haveCnt[material.TypeId] = 0
	}

	for i, cell := range cells {
		if cell.Item == nil || cell.Num == 0 {
			new_cells[i] = cell
			continue
		}
		id := cell.Item.TypeId
		if _, yes := requiredCnt[id]; yes {
			if haveCnt[id] < requiredCnt[id] {
				// 尚未计算完消耗材料个数

				to_use := requiredCnt[id] - haveCnt[id]
				if to_use < cell.Num {
					// 当前格子消耗后已经够了
					haveCnt[id] = haveCnt[id] + to_use
					cell.Num = cell.Num - to_use
					new_cells[i] = cell
				} else if to_use == cell.Num {
					haveCnt[id] = haveCnt[id] + cell.Num
					cell.Num = 0
					cell.Item = nil
					new_cells[i] = cell
				} else {
					haveCnt[id] = haveCnt[id] + cell.Num
					cell.Num = 0
					cell.Item = nil
					new_cells[i] = cell
				}
			} else {
				new_cells[i] = cell
			}
		} else {
			new_cells[i] = cell
		}
	}

	for id, cnt := range requiredCnt {
		if haveCnt[id] < cnt {
			log.Warn("###################", id)
			return false, nil
		}
	}

	return true, new_cells
}

func (g *game) decompose(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.DecomposeRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}

	//
	info, err := g.db.GetUser(req.Uid)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	if info.Diamond < o.CostDiamond {
		res.Data = &msg.DecomposeResponse{ErrMsg: "钻石不足"}
		return res
	}
	if info.Gold < o.CostGold {
		res.Data = &msg.DecomposeResponse{ErrMsg: "金币不足"}
		return res
	}

	gold := info.Gold - o.CostGold
	diamond := info.Diamond - o.CostDiamond

	cells := info.Characters[o.CharacterId].Bag.Cells

	if o.DecomposedCellIndex > uint(len(cells)) ||
		cells[o.DecomposedCellIndex].Item == nil ||
		cells[o.DecomposedCellIndex].Num == 0 {
		res.Data = &msg.DecomposeResponse{ErrMsg: "物品不存在"}
		return res
	}

	ok, decomposeId, decomposeNum := func() (bool, int, uint) {
		for _, v := range csv.Table_Workshop {
			if v.Goods_id == cells[o.DecomposedCellIndex].Item.TypeId {
				return true, v.Reslove_materials[0], uint(v.Reslove_materials_num[0])
			}
		}
		return false, -1, 0
	}()

	if !ok {
		res.Data = &msg.DecomposeResponse{ErrMsg: "物品类型错误，不是符文"}
		return res
	}

	sq := uint(99)

	// 放置背包
	cells[o.DecomposedCellIndex].Item = nil
	cells[o.DecomposedCellIndex].Num = 0

	count := decomposeNum
	for i, v := range cells {
		if v.Item == nil || v.Num == 0 {
			n := uint(0)
			if count >= sq {
				n = sq
			} else {
				n = count
			}
			count = count - n
			cells[i].Num = n
			cells[i].Item = &odm.Item{
				Guid:            util.NewSnowflakeID(),
				TypeId:          decomposeId,
				StarLevel:       0,
				StarUpgradeInfo: []int{},
			}
			if count <= 0 {
				break
			}
		}
	}

	if count > 0 {
		res.Data = &msg.DecomposeResponse{ErrMsg: "背包空间不足"}
		return res
	}

	// 更新数据稍后做成一次性update的
	g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".bag.cells", cells)
	g.db.UpsertUserField(req.Uid, "gold", gold)
	g.db.UpsertUserField(req.Uid, "diamond", diamond)
	bag := info.Characters[o.CharacterId].Bag
	bag.Cells = cells
	res.Data = &msg.DecomposeResponse{
		IsSuccess: true,
		TypeId:    decomposeId,
		Num:       decomposeNum,
		Gold:      gold,
		Diamond:   diamond,
		BagStatus: bag,
	}
	return res
}

func (g *game) enhance(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	var o msg.EnhanceRequest
	err := jsoniter.UnmarshalFromString(req.Data, &o)
	if err != nil {
		res.Err = msg.ErrDecodeJSON
		res.ErrMsg = err.Error()
		return res
	}

	//
	info, err := g.db.GetUser(req.Uid)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	if info.Diamond < o.CostDiamond {
		res.Data = &msg.EnhanceResponse{ErrMsg: "钻石不足"}
		return res
	}
	if info.Gold < o.CostGold {
		res.Data = &msg.EnhanceResponse{ErrMsg: "金币不足"}
		return res
	}

	gold := info.Gold - o.CostGold
	diamond := info.Diamond - o.CostDiamond

	ok, cells := checkMaterials(o.Materials, info.Characters[o.CharacterId].Bag.Cells)

	if !ok {
		res.Data = &msg.EnhanceResponse{ErrMsg: "材料不足"}
		return res
	}

	ok, index := func(i int, cells []odm.Cell) (bool, int) {
		if i < 0 || i >= len(cells) {
			return false, -1
		}
		if cells[i].Item == nil || cells[i].Num == 0 {
			return false, -1
		}
		return true, i
	}(o.ToEnhanceIndex, cells)

	if !ok {
		res.Data = &msg.EnhanceResponse{ErrMsg: "被强化物品信息错误"}
		return res
	}
	//

	// [1,4]
	upgrade := upgradeRandom()
	//

	old := cells[index]
	new := old
	new.Item.StarLevel = new.Item.StarLevel + 1
	new.Item.StarUpgradeInfo = append(new.Item.StarUpgradeInfo, upgrade)
	cells[index] = new

	// 更新数据稍后做成一次性update的
	g.db.UpsertUserField(req.Uid, "characters."+o.CharacterId+".bag.cells", cells)
	g.db.UpsertUserField(req.Uid, "gold", gold)
	g.db.UpsertUserField(req.Uid, "diamond", diamond)

	bag := info.Characters[o.CharacterId].Bag
	bag.Cells = cells
	res.Data = &msg.EnhanceResponse{
		IsSuccess:    true,
		EnhancedItem: new,
		BagStatus:    bag,
		Gold:         gold,
		Diamond:      diamond,
	}
	return res
}

type upgrade_prob struct {
	star int
	prob int
}

func (d upgrade_prob) Prob() int {
	return d.prob
}

var lot lottery.Lottery
var up_prob []lottery.Interface = []lottery.Interface{
	upgrade_prob{1, 50},
	upgrade_prob{2, 30},
	upgrade_prob{3, 15},
	upgrade_prob{4, 5},
}

func init() {
	lot = lottery.New(rand.New(rand.NewSource(time.Now().UnixNano())))
}

func upgradeRandom() int {
	return up_prob[lot.Lots(up_prob...)].(upgrade_prob).star
}
