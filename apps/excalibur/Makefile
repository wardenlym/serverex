.PHONY: all test codegen build beta approval rollback_beta prod rollback_prod

PLATFORM=$(shell go env GOOS)
ifeq ($(PLATFORM), windows)
	ext=.exe
endif

PROJECT=excalibur

all: build

codegen:
	go generate ./...
	go test ./... -count=1 -run "_codegen_"
	go test ./... -v -count=1 -run "_codegen_"

build:
ifeq ($(PLATFORM), windows)
	cd gate_srv && go build -tags=jsoniter
else
	cd gate_srv && go build -tags=jsoniter -ldflags -s\
		-ldflags "-X main.VERSION=1.0.0 -X 'main.GITHASH=`git rev-parse --short HEAD`' -X 'main.BUILT=`date +"%Y/%m/%d %H:%M:%S"`' -X 'main.GOVERSION=`go version | cut -d" " -f 3`'"
endif

DATE=$(shell date --iso=seconds)
beta: build
	mkdir -p /rel/apps/excalibur_beta && cd /rel/apps/excalibur_beta/ && mkdir -p log data conf backup
	cp conf/config.beta.toml /rel/apps/excalibur_beta/conf/
	cp conf/server.crt /rel/apps/excalibur_beta/conf/
	cp conf/server.key /rel/apps/excalibur_beta/conf/
	-cp /rel/apps/excalibur_beta/gate_srv /rel/apps/excalibur_beta/backup/gate_srv.lastest
	-mv /rel/apps/excalibur_beta/gate_srv /rel/apps/excalibur_beta/backup/gate_srv.$(DATE)
	cp gate_srv/gate_srv /rel/apps/excalibur_beta/gate_srv
	supervisorctl restart ex_beta

sup:
	cp sup_excalibur.conf /etc/supervisord.d/
	supervisorctl update

test:
	go test ./... -count=1

rollback_beta:
	rm /rel/apps/excalibur_beta/gate_srv
	cp /rel/apps/excalibur_beta/backup/gate_srv.lastest /rel/apps/excalibur_beta/gate_srv
	supervisorctl restart ex_beta

prod: build
	mkdir -p /rel/apps/excalibur_prod && cd /rel/apps/excalibur_prod/ && mkdir -p log data conf backup
	cp conf/config.prod.toml /rel/apps/excalibur_prod/conf/
	cp conf/server.crt /rel/apps/excalibur_prod/conf/
	cp conf/server.key /rel/apps/excalibur_prod/conf/
	-cp /rel/apps/excalibur_prod/gate_srv /rel/apps/excalibur_prod/backup/gate_srv.lastest
	-mv /rel/apps/excalibur_prod/gate_srv /rel/apps/excalibur_prod/backup/gate_srv.$(DATE)
	cp gate_srv/gate_srv /rel/apps/excalibur_prod/gate_srv
	supervisorctl restart ex_prod

rollback_prod:
	rm /rel/apps/excalibur_prod/gate_srv
	cp /rel/apps/excalibur_prod/backup/gate_srv.lastest /rel/apps/excalibur_prod/gate_srv
	supervisorctl restart ex_prod

approval: build
	mkdir -p /rel/apps/excalibur_approval && cd /rel/apps/excalibur_approval/ && mkdir -p log data conf backup
	cp conf/config.approval.toml /rel/apps/excalibur_approval/conf/
	cp conf/server.crt /rel/apps/excalibur_approval/conf/
	cp conf/server.key /rel/apps/excalibur_approval/conf/
	-cp /rel/apps/excalibur_approval/gate_srv /rel/apps/excalibur_approval/backup/gate_srv.lastest
	-mv /rel/apps/excalibur_approval/gate_srv /rel/apps/excalibur_approval/backup/gate_srv.$(DATE)
	cp gate_srv/gate_srv /rel/apps/excalibur_approval/gate_srv
	supervisorctl restart ex_approval

rollback_approval:
	rm /rel/apps/excalibur_approval/gate_srv
	cp /rel/apps/excalibur_approval/backup/gate_srv.lastest /rel/apps/excalibur_approval/gate_srv
	supervisorctl restart ex_approval

taptap: build
	mkdir -p /rel/apps/excalibur_taptap && cd /rel/apps/excalibur_taptap/ && mkdir -p log data conf backup
	cp conf/config.taptap.toml /rel/apps/excalibur_taptap/conf/
	cp conf/server.crt /rel/apps/excalibur_taptap/conf/
	cp conf/server.key /rel/apps/excalibur_taptap/conf/
	-cp /rel/apps/excalibur_taptap/gate_srv /rel/apps/excalibur_taptap/backup/gate_srv.lastest
	-mv /rel/apps/excalibur_taptap/gate_srv /rel/apps/excalibur_taptap/backup/gate_srv.$(DATE)
	cp gate_srv/gate_srv /rel/apps/excalibur_taptap/gate_srv
	supervisorctl restart ex_taptap

rollback_taptap:
	rm /rel/apps/excalibur_taptap/gate_srv
	cp /rel/apps/excalibur_taptap/backup/gate_srv.lastest /rel/apps/excalibur_taptap/gate_srv
	supervisorctl restart ex_taptap