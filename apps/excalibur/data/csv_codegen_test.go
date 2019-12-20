package data

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/json-iterator/go"
	"gitlab.com/megatech/serverex/apps/excalibur/data/csv"
)

var UTF8_BOM = []byte{239, 187, 191}

func hasBOM(in []byte) bool {
	return bytes.HasPrefix(in, UTF8_BOM)
}

func stripBOM(in []byte) []byte {
	if hasBOM(in) {
		return bytes.TrimPrefix(in, UTF8_BOM)
	}
	return in
}

func fixType(typename string) string {
	switch typename {
	case "float":
		return "float32"
	case "int[]", "int[:]":
		return "[]int"
	case "float[]":
		return "[]float32"
	case "string[]":
		return "[]string"
	default:
		return typename
	}
}

func GenCSVStruct(filename string, lines []string, filter_use func(string, string) bool) string {
	comments := strings.Split(strings.Trim(lines[0], "\r\n"), ",")
	fieldname := strings.Split(strings.Trim(lines[1], "\r\n"), ",")
	types := strings.Split(strings.Trim(lines[2], "\r\n"), ",")

	if len(fieldname) != len(types) || len(fieldname) != len(comments) {
		panic("csv 行数错误")
	}

	src := fmt.Sprintf("type %s struct {", filename)
	for i := 0; i < len(fieldname); i++ {

		if !filter_use(filename, fieldname[i]) {
			continue
		}

		if i > 0 {
			src = src + "\n"
		}
		src = src + fmt.Sprintf("\n    // %s", comments[i])
		src = src + fmt.Sprintf("\n    %s %s", strings.Title(fieldname[i]), fixType(types[i]))
	}
	src = src + "\n}\n"
	return src
}

func genDataMapTable(name string, t *testing.T) {
	path := "./csv/" + name + ".csv"

	filename := "0gen_csv_Table_" + name + ".go"

	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(stripBOM(b)), "\n")

	src1 := GenCSVStruct(name, lines, filter_use_consume_size)

	src := genCSVtoMapDefine(name, lines, filter_use_consume_size)

	err = writeGolangFile("./csv/"+filename, genGolangFile("csv", src1+"\n"+src))
	if err != nil {
		t.Error(err, filename)
		ioutil.WriteFile("errr.go", []byte(src), 0644)
	}
}

func filter_use_consume_size(name string, field string) bool {

	if name == "EquipData" {
		if field == "quality" {
			return true
		} else {
			return false
		}
	}

	if name == "GoodsItem" {

		switch field {
		case "name", "intro", "icon", "icon2", "altasPath":
			return false
		}
	}

	if name == "LootoneData" {

		switch field {
		case "group_description":
			return false
		}
	}

	if name == "LoottwoData" {

		switch field {
		case "loot_type", "description", "remarks":
			return false
		}
	}

	if name == "MonsterData" {

		switch field {
		case "id", "killsvalue", "monster_type":
			return true
		default:
			return false
		}
	}

	return true
}

var Confirmed_CSV = []string{
	"StageData",
	"MarketData",
	"RuneData",
	"Workshop",
	"GoodsItem",
	"MonsterData",
	"RechargeData",
	"NoticeData",
	"LootoneData",
	"LoottwoData",
	"Adventure",
	"HeroData",
	"AchievementData",
	"Rewards_FirstPay",
	"Rewards_Week",
	"Rewards_Month",
	"DunRule_Gift",
	"DunRule_GhostWd",
	"RuneopenData",
	"Hero_levelup",
	"GiftPackage",
	"SkuApple",
	"SkuGoogle",
}

func Test_codegen_GenDataMapTable(t *testing.T) {

	for _, name := range Confirmed_CSV {
		genDataMapTable(name, t)
	}

	// genDataMapTable("StageData", t)
	// genDataMapTable("MarketData", t)
	// genDataMapTable("RuneData", t)
	// genDataMapTable("Workshop", t)
	// genDataMapTable("GoodsItem", t)
	// genDataMapTable("MonsterData", t)
	// genDataMapTable("RechargeData", t)
	// genDataMapTable("NoticeData", t)
	// genDataMapTable("LootoneData", t)
	// genDataMapTable("LoottwoData", t)
	// genDataMapTable("Adventure", t)
	// genDataMapTable("HeroData", t)
	// genDataMapTable("AchievementData", t)
	// genDataMapTable("Rewards_FirstPay", t)
	// genDataMapTable("Rewards_Week", t)
	// genDataMapTable("Rewards_Month", t)
	// genDataMapTable("DunRule_Gift", t)
	// genDataMapTable("DunRule_GhostWd", t)
	// genDataMapTable("RuneopenData", t)
	// genDataMapTable("Hero_levelup", t)
	// genDataMapTable("GiftPackage", t)
	// genDataMapTable("HeroPrice", t)

}

// func Test_codegen_GenData(t *testing.T) {

// 	genData("StageData", &csv.Table_StageData, t)
// 	genData("MarketData", &csv.Table_MarketData, t)
// 	genData("RuneData", &csv.Table_RuneData, t)
// 	genData("RechargeData", &csv.Table_RechargeData, t)
// 	genData("Workshop", &csv.Table_Workshop, t)
// 	genData("GoodsItem", &csv.Table_GoodsItem, t)
// 	genData("MonsterData", &csv.Table_MonsterData, t)
// 	genData("RechargeData", &csv.Table_RechargeData, t)
// 	genData("NoticeData", &csv.Table_NoticeData, t)
// 	genData("LootoneData", &csv.Table_LootoneData, t)
// 	genData("LoottwoData", &csv.Table_LoottwoData, t)
// 	genData("Adventure", &csv.Table_Adventure, t)
// 	genData("HeroData", &csv.Table_HeroData, t)
// 	genData("AchievementData", &csv.Table_AchievementData, t)
// 	genData("Rewards_FirstPay", &csv.Table_Rewards_FirstPay, t)
// 	genData("Rewards_Week", &csv.Table_Rewards_Week, t)
// 	genData("Rewards_Month", &csv.Table_Rewards_Month, t)
// 	genData("DunRule_Gift", &csv.Table_DunRule_Gift, t)
// 	genData("DunRule_GhostWd", &csv.Table_DunRule_GhostWd, t)

// }

func genData(name string, v interface{}, t *testing.T) {
	b, err := jsoniter.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Error(err)
	}
	ioutil.WriteFile("csvdata/"+name+".json", b, 0644)

	tpl := `
package csvdata

import "gitlab.com/megatech/serverex/apps/excalibur/data/csv"

var Table_` + name + ` map[int]csv.` + name + ``

	dst, err := format.Source([]byte(tpl))
	if err != nil {
		t.Error(err)
	}

	ioutil.WriteFile("csvdata/Table_"+name+".go", []byte(dst), 0644)
	if err != nil {
		t.Error(err)
	}
}

func Test_load_data(t *testing.T) {
	csv.LoadData()
}
