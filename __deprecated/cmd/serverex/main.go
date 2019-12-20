package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/pelletier/go-toml"
	"gitlab.com/megatech/serverex/cmd"

	"github.com/urfave/cli"
)

const (
	// Version .
	Version = "0.0.1"
)

type Option struct {
	AppName string
	Alias   string
	Exe     string
	Arg     string
}

type Config struct {
	BaseDir string
	Apps    []Option
}

func base_dir() string {
	return filepath.FromSlash(os.Getenv("GOPATH") + "/src/gitlab.com/megatech/serverex/")
}

func exe_name(basedir string, appname string, exe string) string {
	return filepath.FromSlash(basedir + "apps/" + appname + "/bin/" + cmd.Exename(exe))
}

var sample Config = Config{
	BaseDir: base_dir(),
	Apps: []Option{
		Option{
			AppName: "excalibur",
			Alias:   "ex",
			Exe:     "excalibur_gate_http",
			Arg:     "-mode=dev",
		},
	},
}

var config Config = sample

func get_dir() string {
	u, e := user.Current()
	if e != nil {
		return ""
	}

	return u.HomeDir
}

func find_exe(appname string) (bool, string, string, string) {
	for _, v := range config.Apps {
		if appname == v.Alias || appname == v.AppName {
			return true, exe_name(config.BaseDir, v.AppName, v.Exe), v.Arg, v.AppName
		}
	}
	return false, "false", "", ""
}

func command(app *cli.App) []cli.Command {

	start := cli.Command{
		Name:      "start",
		Usage:     "Start servers of a app.",
		UsageText: app.Name + " start [app]",
		Action: func(c *cli.Context) error {
			if ok, exe, args, _ := find_exe(c.Args().First()); ok {
				cmd.Start(exe, args)
			} else {
				fmt.Println("input a correct app name.")
			}
			return nil
		},
	}
	stop := cli.Command{
		Name:      "stop",
		Usage:     "Stop servers of a app.",
		UsageText: app.Name + " stop [app]",
		Action: func(c *cli.Context) error {
			if ok, _, _, appname := find_exe(c.Args().First()); ok {
				cmd.Stop(appname)
				time.Sleep(500)
			} else {
				fmt.Println("input a correct app name.")
			}
			return nil
		},
	}
	foreground := cli.Command{
		Name:      "foreground",
		Usage:     "Start servers of a app with default service at foreground.",
		UsageText: app.Name + " foreground [app]",
		Action: func(c *cli.Context) error {
			if ok, exe, args, _ := find_exe(c.Args().First()); ok {
				cmd.RunFreground(exe, args)
			} else {
				fmt.Println("input a correct app name.")
			}
			return nil
		},
	}

	infra := cli.Command{
		Name: "infra",
	}
	showconf := cli.Command{
		Name: "showconf",
		Action: func(c *cli.Context) error {
			s, _ := toml.Marshal(sample)
			fmt.Println(string(s))
			return nil
		},
	}
	return []cli.Command{start, stop, foreground, infra, showconf}
}

func main() {
	app := cli.NewApp()
	app.Name = "serverex"
	app.Usage = "A command line tool for megatech-backend apps management."
	app.UsageText = app.Name + " command [command options...]"
	app.Version = Version

	app.Commands = command(app)
	app.Run(os.Args)
}

// func main() {
// 	app := cli.NewApp()
// 	app.Usage = "Command line tool for app management."
// 	app.UsageText = app.Name + " command [command options...]"
// 	app.Version = Version

// 	appCmd := cli.Command{
// 		Name:  "app",
// 		Usage: "Shows app info",
// 		Action: func(c *cli.Context) error {
// 			fmt.Println("excalibur")
// 			fmt.Println("demo-server")
// 			return nil
// 		},
// 	}

// 	startCmd := cli.Command{
// 		Name:      "start",
// 		Usage:     "Start servers of a app.",
// 		UsageText: app.Name + " start [app]",
// 		Flags: []cli.Flag{
// 			cli.StringFlag{
// 				Name:  "mode",
// 				Usage: "-mode=dev,beta,prod",
// 			},
// 		},
// 		Action: func(c *cli.Context) error {

// 			mode := "dev"
// 			if c.String("mode") != "" {
// 				mode = c.String("mode")
// 			}

// 			switch c.Args().First() {
// 			case "excalibur", "ex":
// 				// ./serverex stop ex -dev
// 				fmt.Println("starting excalibur_api.")
// 				if mode == "dev" {
// 					fmt.Println("dev mode on.")
// 					cmd.IfNotExistRunMongodbDaemon("./infra/mongodb")
// 				}

// 				cmd.BlockingStartExcaliburAPI(mode)
// 			case "mongo", "mongodb", "mongod":
// 				fmt.Println("starting mongodb.")
// 				cmd.RunMongodbDaemon("./infra/mongodb")

// 			case "demo", "demo-server":
// 				cmd.StartDemoServer(c.Args().First())

// 			default:
// 				fmt.Println("input a correct app name.")
// 			}
// 			return nil
// 		},
// 	}
// 	stopCmd := cli.Command{
// 		Name:      "stop",
// 		Usage:     "Stop servers of a app.",
// 		UsageText: app.Name + " stop [app]",
// 		Action: func(c *cli.Context) error {
// 			switch c.Args().First() {
// 			case "excalibur", "ex":
// 				// ./serverex stop ex
// 				fmt.Println("stoping excalibur-api.")
// 				cmd.StopExcaliburAPI()
// 			case "mongo", "mongodb", "mongod":
// 				fmt.Println("try stopping mongodb.")
// 				cmd.TryStopMongodbDaemon()
// 			case "demo", "demo-server":
// 				cmd.TryStopDemoServer()
// 			default:
// 				fmt.Println("input a correct app name.")
// 			}
// 			return nil
// 		},
// 	}
// 	statusCmd := cli.Command{
// 		Name:      "status",
// 		Aliases:   []string{"st"},
// 		Usage:     "Shows server status.",
// 		UsageText: app.Name + " stop [app]",
// 		Action: func(c *cli.Context) error {
// 			fmt.Println("\nservers status:")
// 			cmd.ShowStatus()
// 			return nil
// 		},
// 	}
// 	restartCmd := cli.Command{
// 		Name:      "restart",
// 		Usage:     "Restart server.",
// 		UsageText: app.Name + " restart [app]",
// 		Flags: []cli.Flag{
// 			cli.StringFlag{
// 				Name:  "mode",
// 				Usage: "-mode=dev,beta,prod",
// 			},
// 		},
// 		Action: func(c *cli.Context) error {

// 			mode := "dev"
// 			if c.String("mode") != "" {
// 				mode = c.String("mode")
// 			}

// 			switch c.Args().First() {
// 			case "excalibur", "ex":
// 				if mode == "dev" {
// 					fmt.Println("dev mode on.")
// 					cmd.IfNotExistRunMongodbDaemon("./infra/mongodb")
// 				}
// 				fmt.Println("stoping excalibur_api.")
// 				cmd.StopExcaliburAPI()
// 				time.Sleep(100)
// 				fmt.Println("restarting excalibur_api.")
// 				cmd.StartExcaliburAPI(mode)
// 			case "demo", "demo-server":
// 				cmd.TryStopDemoServer()
// 				cmd.StartDemoServer(c.Args().First())
// 			default:
// 				fmt.Println("input a correct app name.")
// 			}
// 			return nil
// 		},
// 	}

// 	app.Commands = []cli.Command{appCmd, startCmd, stopCmd, statusCmd, restartCmd}
// 	app.Run(os.Args)
// }
