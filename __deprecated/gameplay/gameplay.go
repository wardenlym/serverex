package gameplay

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"

	"gitlab.com/megatech/serverex/apps/excalibur"
	"gitlab.com/megatech/serverex/apps/excalibur/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/repo"
	"gitlab.com/megatech/serverex/lib/log"
)

type GamePlay interface {
	Handle(c *gin.Context)
	Stop()
}

func NewGamePlay(config excalibur.Config) GamePlay {
	p := &game{
		db:     repo.NewDB(),
		config: config,
	}
	p.db.Open(config.MongoAddress)
	p.initHandle()
	return p
}

type game struct {
	dispatch map[msg.Code]func(req *msg.Request) *msg.Response
	db       *repo.DB
	config   excalibur.Config
}

func (p *game) Stop() {
	if p.db != nil {
		p.db.Close()
	}
}

func (g *game) Handle(c *gin.Context) {

	req, err := g.parseRequest(c)
	if err != nil {
		log.Error(err)
		c.Status(400)
		return
	}

	res := g.handleRequest(req)

	log.Debug("->", fmt.Sprintf("#%d ", req.Seq), req.Code, c.Request.URL, req.Data)

	s, err := jsoniter.MarshalToString(res)
	if err != nil {
		log.Error(err)
		c.Status(500)
		return
	}

	log.Debug("<-", fmt.Sprintf("#%d ", req.Seq), req.Code, s)
	c.String(200, s)
}

func (g *game) parseRequest(c *gin.Context) (*msg.Request, error) {

	var err error

	code, err := strconv.Atoi(c.Query("code"))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	uid, err := strconv.ParseInt(c.Query("uid"), 10, 64)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	seq, _ := strconv.Atoi(c.Writer.Header().Get("Seqid"))
	data, _ := c.GetRawData()

	return &msg.Request{
		Code: msg.Code(code),
		Uid:  uid,
		Seq:  seq,
		Data: string(data),
	}, nil
}

func (g *game) handleRequest(req *msg.Request) *msg.Response {
	f := g.dispatch[req.Code]
	if f != nil {
		return f(req)
	}
	return notImplemented(req)
}

func notImplemented(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	res.Err = msg.ErrNotImplemented
	res.ErrMsg = "code:" + fmt.Sprint(int(req.Code)) + " not implemented"
	return res
}

func (g *game) initHandle() {
	g.dispatch = map[msg.Code]func(req *msg.Request) *msg.Response{}

	getname := func(v interface{}) string {
		full := runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
		a := strings.Split(full, "/")
		return strings.TrimSuffix(a[len(a)-1], ")-fm")
	}

	set := func(k msg.Code, v func(req *msg.Request) *msg.Response) {
		g.dispatch[k] = v
		log.Printf("Msg => Func: %19s => %#v", k, getname(v))
	}

	set(msg.C_GetUserInfo, g.getUserInfo)
	set(msg.C_EnterStage, g.enterStage)
	set(msg.C_ExitStage, g.exitStage)
	set(msg.C_BattleStart, g.battleStart)
	set(msg.C_BattleEnd, g.battleEnd)
	set(msg.C_DestroyItem, g.destroyItem)
	set(msg.C_ConsumItem, g.consumItem)
	set(msg.C_Equip, g.equip)
	set(msg.C_Unequip, g.unequip)
	set(msg.C_ListOnSale, g.listOnSale)
	set(msg.C_RefreshOnSale, g.refreshOnSale)
	set(msg.C_Purchase, g.purchase)
	set(msg.C_ListDiamondPrice, g.listDiamondPrice)
	set(msg.C_SubmitOrder, g.submitOrder)
	set(msg.C_QueryOrderStatus, g.queryOrderStatus)
	set(msg.C_Craft, g.craft)
	set(msg.C_Decompose, g.decompose)
	set(msg.C_Enhance, g.enhance)
	set(msg.C_LevelUp, g.levelUp)
	set(msg.C_JobUpgrade, g.jobupgrade)
	set(msg.C_RuneEquip, g.runeEquip)
	set(msg.C_RuneUnequip, g.runeUnequip)

	if g.config.ApiOption.DevModeEnable {
		set(msg.C_GM_ResetUser, g.gm_ResetUser)
		set(msg.C_GM_CreateItem, g.gm_CreateItem)
		set(msg.C_GM_SetMyMoney, g.gm_SetMyMoney)
		set(msg.C_GM_ResetJob, g.gm_ResetJob)
	}
}

func (g *game) getUserInfo(req *msg.Request) *msg.Response {
	res := msg.FromRequest(req)
	info, err := g.db.GetUserOrUpsert(req.Uid)
	if err != nil {
		res.Err = msg.ErrDBError
		res.ErrMsg = err.Error()
		return res
	}

	res.Data = &msg.GetUserInfoResponse{
		User: info,
	}
	return res
}
