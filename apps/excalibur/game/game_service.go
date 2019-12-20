package game

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http/httputil"
	"runtime"
	"strconv"
	"strings"
	"sync"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/json-iterator/go"
	"gitlab.com/megatech/serverex/apps/excalibur"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/msg"
	"gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"
	"gitlab.com/megatech/serverex/apps/excalibur/repo"
	"gitlab.com/megatech/serverex/lib/log"
)

type GameService struct {
	db          *repo.DB
	config      excalibur.Config
	dbcache     *sync.Map
	loop_stoper chan bool
}

func NewGameService(config excalibur.Config) *GameService {
	g := &GameService{
		db:          repo.NewDB(),
		config:      config,
		dbcache:     &sync.Map{},
		loop_stoper: make(chan bool),
	}
	err := g.db.Open(config.MongoAddress)
	if err != nil {
		log.Fatal("mongodb: ", err)
	}
	return g
}

func (g *GameService) Run() {
	go g.loop()
}

func (g *GameService) Stop() {
	g.stop_loop()
}

func (g *GameService) WaitExit() {
}

func decode(c *gin.Context) (*msg.Request, error) {

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

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

func encode(o *msg.Response) (string, error) {
	return jsoniter.MarshalToString(o)
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contains dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastslash := bytes.LastIndex(name, slash); lastslash >= 0 {
		name = name[lastslash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}

// stack returns a nicely formated stack frame, skipping skip frames
func stack(skip int) []byte {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

func get_agent(c *gin.Context) string {
	n := c.GetHeader("User-Agent")
	if n == "" {
		return "ExcaliburClientSDK/0.0.1"
	}
	return n
}

func (g *GameService) Handle(c *gin.Context) {

	req, err := decode(c)
	if err != nil {
		log.Error(err)
		c.Status(400)
		return
	}

	agent := get_agent(c)

	if req.Code != msg.C_SetUserProperties && req.Code != msg.C_GetUserProperties {
		log.Debug(fmt.Sprintf("->#%d %s %s %s %s", req.Seq, code_to_string(req.Code), agent, c.Request.URL, req.Data))
	}

	g.check_jwt(c, req.Uid)

	// protect call
	ret := func(req *msg.Request) (err_ret *msg.Response) {

		defer func() {
			if err := recover(); err != nil {

				stack := stack(1)
				httprequest, _ := httputil.DumpRequest(c.Request, false)
				log.Printf("[Recovery] panic recovered:\n%s\n%s\n%s%s", string(httprequest), err, stack, string([]byte{27, 91, 48, 109}))

				if fmt.Sprint(err) == "runtime error: index out of range" {
					err_ret = msg.ErrFromRequest(req, msg.ErrInvalidCellIndex, errcode_to_string(msg.ErrInvalidCellIndex))
				} else {
					err_ret = msg.ErrFromRequest(req, -1, "")
				}
			}
		}()
		return g.call(req)
	}(req)

	s, err := encode(ret)
	if err != nil {
		log.Error(err)
		c.Status(500)
		return
	}

	if req.Code != msg.C_SetUserProperties && req.Code != msg.C_GetUserProperties {
		log.Debug(fmt.Sprintf("<-#%d %s %s", req.Seq, code_to_string(req.Code), s))
	}
	c.String(200, s)
}

func (g *GameService) call(req *msg.Request) *msg.Response {

	if g.config.GateHttpOption.DevModeEnable == false && req.Code >= 200 {
		return msg.Err(req, msg.ErrNotImplemented, "not implemented")
	}

	ret := g.rpc_dispatch(req)
	return ret
}

func (g *GameService) get_user(uid int64) *odm.User {
	user, e := g.db.GetUser(uid)
	if e != nil {
		if e == mgo.ErrNotFound {
			return nil
		}
		log.Error(e)
		return nil
	}
	return user
}

func (g *GameService) save_user(info *odm.User) {
	e := g.db.UpdateUser(info.Uid, info)
	if e != nil {
		log.Error("##**## DbSaveError:", e)
	}
}

var hmacSampleSecret = []byte("MegaIdentitySecurityKey")

func (g *GameService) check_jwt(c *gin.Context, url_uid int64) {
	token_str := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	token, err := jwt.Parse(token_str, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if token == nil || err != nil {
		log.Error(err)
		log.Error(token_str)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// log.Debug(claims)
		uid, _ := strconv.ParseInt(claims["userId"].(string), 10, 64)
		if uid != url_uid {
			log.Error(token_str)
			log.Error(claims)
			log.Error("Token UID Check Fail:", uid, url_uid)
		}
	} else {
		log.Error(err)
		log.Error(token_str)
		log.Error(claims)
	}
}
