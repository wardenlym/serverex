package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"gitlab.com/megatech/serverex/apps/excalibur"
	"gitlab.com/megatech/serverex/apps/excalibur/game"
	"gitlab.com/megatech/serverex/lib/log"
	"gitlab.com/megatech/serverex/lib/util"
)

var (
	VERSION   string = "dev"
	GITHASH   string
	BUILT     string
	GOVERSION string
)

func main() {
	var err error
	var mode string
	var cfg string
	var showver bool
	flag.BoolVar(&showver, "v", false, "show version info then exit")
	flag.StringVar(&mode, "mode", "dev", "config mode")
	flag.StringVar(&cfg, "c", "", "use a specific config file")
	flag.Parse()

	if showver {
		fmt.Printf("%s %s %s %s\n", VERSION, GITHASH, BUILT, GOVERSION)
		os.Exit(0)
	}

	log.EnableSwitchBySIGUSR2(func(debugMode bool) {
		if debugMode {
			gin.SetMode(gin.DebugMode)
			log.Info("# Enable Debug Log")
		} else {
			gin.SetMode(gin.ReleaseMode)
			log.Info("# Disable Debug Log")
		}
	})

	log.Info("--------------------")
	log.Info("excalibur start.")
	log.Info(fmt.Sprintf("complie info: %s %s %s %s\n", VERSION, GITHASH, BUILT, GOVERSION))

	if mode == "prod" {
		log.EnableDebug(false)
		gin.SetMode(gin.ReleaseMode)
	}

	config := excalibur.Config{}

	if cfg != "" {
		err = util.ParseTOML(cfg, &config)
	} else {
		err = util.ParseTOMLByMode(mode, &config)
	}

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", config)

	// 加載數據

	// data.LoadData()

	// 业务逻辑
	engine := gin.New()
	engine.Use(util.GinLoggerMiddleware(), gin.Recovery())

	// 每次请求会附带一个顺序的调用id，可以用来跟踪调试
	var seqID int32
	engine.Use(func(c *gin.Context) {
		c.Header("Seqid", fmt.Sprint(atomic.AddInt32(&seqID, 1)))
		c.Next()
	})

	api := engine.Group("/excalibur/api")

	// 游戏玩法部分
	gs := game.NewGameService(config)
	gs.Run()
	api.POST("/gs", gs.Handle)

	// 服务启停
	srv := &http.Server{
		Addr:    config.GateHttpOption.ListenAddress,
		Handler: engine,

		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// 监听退出信号
	go func(srv *http.Server) {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
		s := <-c
		log.Warn("### signal received:", s)
		log.Info("# waiting http stop.")
		waitingShutdown(srv)
	}(srv)

	log.Info("excalibur running.")
	err = func() error {
		if config.GateHttpOption.UseTLS {
			log.Info("using TLS")
			return srv.ListenAndServeTLS("./conf/server.crt", "./conf/server.key")
		} else {
			return srv.ListenAndServe()
		}
	}()
	if err != nil {
		log.Warn(err)
	}

	log.Info("# waiting gs stop.")
	gs.Stop()
	gs.WaitExit()

	log.Info("# all service exited.")
	log.Info("excalibur shutdown.")
}

func waitingShutdown(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Warn(err)
	}
}
