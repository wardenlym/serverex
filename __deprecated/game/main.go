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
	"github.com/json-iterator/go"

	"gitlab.com/megatech/serverex/lib/log"
	"gitlab.com/megatech/serverex/lib/util"
	"gitlab.com/megatech/serverex/project/excalibur"
	"gitlab.com/megatech/serverex/project/excalibur/account"
	"gitlab.com/megatech/serverex/project/excalibur/game"
)

func main() {
	var err error
	var mode string
	var configFile string
	flag.StringVar(&configFile, "c", "", "config file path")
	flag.StringVar(&mode, "mode", "dev", "config mode")
	flag.Parse()

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
	log.Info("excalibur-api start.")

	if mode == "prod" {
		log.EnableDebug(false)
		gin.SetMode(gin.ReleaseMode)
	}

	config := excalibur.Config{}
	switch mode {
	case "":
		err = util.ParseTOML(configFile, &config)
	default:
		err = util.ParseTOMLByMode(mode, &config)
	}

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%#v", config)
	}

	engine := gin.New()
	engine.Use(util.GinLoggerMiddleware(), gin.Recovery())

	// 业务逻辑
	router := engine
	router.GET("/test", func(c *gin.Context) { c.JSON(200, map[string]string{"status": "running"}) })

	// 每次请求会附带一个顺序的调用id，可以用来跟踪调试
	var seqID int32
	router.Use(func(c *gin.Context) {
		c.Header("Seqid", fmt.Sprint(atomic.AddInt32(&seqID, 1)))
		c.Next()
	})

	// gateway 用来下发其他服务的入口
	router.GET("/gateway", gateway(config))

	api := router.Group("/excalibur/v1")

	// 账号系统部分
	account := account.NewAccount(config)
	api.POST("/account/register", account.Register)
	api.POST("/account/login", account.Login)

	// 游戏玩法部分
	game := game.NewGame(game.DefaultOption)
	go game.Run()

	// 服务启停
	srv := &http.Server{
		Addr:    config.ListenAddress,
		Handler: engine,
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

	log.Info("excalibur-api running.")
	err = srv.ListenAndServe()
	if err != nil {
		log.Warn(err)
	}

	log.Info("# waiting game stop.")
	game.Stop()
	<-game.Quit()
	account.Stop()
	log.Info("# all service exited.")
	log.Info("excalibur-api shutdown.")
}

func waitingShutdown(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Warn(err)
	}
}

func gateway(config excalibur.Config) func(*gin.Context) {

	return func(c *gin.Context) {
		list := map[string]interface{}{
			"login": config.PublicLoginServerAddress,
			"game":  config.PublicGameServerAddress,
		}
		s, err := jsoniter.MarshalToString(list)
		if err != nil {
			log.Error(err)
			c.Status(400)
			return
		}
		c.String(200, s)
	}
}
