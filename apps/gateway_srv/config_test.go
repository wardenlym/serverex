package main

import (
	"io/ioutil"
	"testing"

	"github.com/pelletier/go-toml"
)

func Test_codegen_write_default_config(t *testing.T) {

	c := Config{
		ServiceList: []ServiceInfo{

			ServiceInfo{
				ServerId:    "game-node-01",
				Version:     "1.0.1",
				Env:         "prod",
				ServiceType: "identity",
				Endpoint:    "https://excalibur.53site.com/identity/api",
			},

			ServiceInfo{
				ServerId:    "assets",
				Version:     "0.0.1",
				Env:         "all",
				ServiceType: "assets",
				// Endpoint:    "https://datacenter.53site.com/excalibur/assets/0.0.1/",
				Endpoint: "https://werewolf.53site.com/excalibur/api/",
			},

			ServiceInfo{
				ServerId:    "test-node-01",
				Version:     "0.0.1",
				Env:         "beta",
				ServiceType: "game",
				Endpoint:    "https://excalibur.53site.com/excalibur/beta/api/gs",
			},

			ServiceInfo{
				ServerId:    "game-node-01",
				Version:     "1.0.1",
				Env:         "approval",
				ServiceType: "game",
				Endpoint:    "https://excalibur.53site.com/excalibur/approval/api/gs",
			},

			ServiceInfo{
				ServerId:    "game-node-01",
				Version:     "1.0.1",
				Env:         "prod",
				ServiceType: "game",
				Endpoint:    "https://excalibur.53site.com/excalibur/prod/api/gs",
			},

			ServiceInfo{
				ServerId:    "game-node-01",
				Version:     "1.0.1",
				Env:         "taptap",
				ServiceType: "game",
				Endpoint:    "https://excalibur.53site.com/excalibur/taptap/api/gs",
			},
		},
	}

	b, err := toml.Marshal(c)
	if err != nil {
		t.Error(err)
	}
	ioutil.WriteFile("config.toml", b, 0644)
}
