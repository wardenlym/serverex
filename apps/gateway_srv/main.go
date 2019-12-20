package main

import (
	"flag"
	"net/http"
	"sync"

	"github.com/json-iterator/go"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"

	"gitlab.com/megatech/serverex/lib/log"
	"gitlab.com/megatech/serverex/lib/util"
)

type ServiceInfo struct {
	ServerId    string
	Version     string
	Env         string
	ServiceType string
	Endpoint    string
}

type Config struct {
	ServiceList []ServiceInfo
}

var mu sync.Mutex
var data string

func _data() string {
	mu.Lock()
	defer mu.Unlock()
	return data
}

func update_config(c Config) {
	mu.Lock()
	defer mu.Unlock()
	b, err := jsoniter.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Error(err)
		return
	}

	// 一些编辑器可能有空操作
	if string(b) == data || string(b) == "{}" {
		return
	}

	data = string(b)
}

func main() {
	var err error
	var cfg_path string
	var address string
	flag.StringVar(&address, "a", "127.0.0.1:20000", "listen address")
	flag.StringVar(&cfg_path, "c", "config.toml", "use a specific config file")
	flag.Parse()

	var config Config
	err = util.ParseTOML(cfg_path, &config)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", config)
	update_config(config)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	watcher.Add(cfg_path)

	go func() {
		defer watcher.Close()
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					c := Config{}
					err = util.ParseTOML(cfg_path, &c)
					if err != nil {
						log.Error(err)
					} else {
						log.Info("modified:", c)
						update_config(c)
					}
				}
			case err := <-watcher.Errors:
				log.Error(err)
			}
		}
	}()

	engine := gin.New()
	engine.Use(util.GinLoggerMiddleware(), gin.Recovery())

	engine.GET("/gateway/endpoints", func(c *gin.Context) {
		c.String(200, _data())
	})

	srv := &http.Server{
		Addr:    address,
		Handler: engine,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Error(err)
	}
	return
}
