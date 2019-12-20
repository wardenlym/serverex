.PHONY: all test build beta approval prod

PLATFORM=$(shell go env GOOS)
ifeq ($(PLATFORM), windows)
	ext=.exe
endif

all: build

codegen:
	go generate ./...
	go test ./... -count=1 -run "_codegen_"
	go test ./... -v -count=1 -run "_codegen_"

build:
ifeq ($(PLATFORM), windows)
	go build ./cmd/identity_srv
else
	cd cmd/identity_srv && go build -o ../../identity_srv -tags=jsoniter -ldflags -s\
		-ldflags "-X main.VERSION=1.0.0 -X 'main.GITHASH=`git rev-parse --short HEAD`' -X 'main.BUILT=`date +"%Y/%m/%d %H:%M:%S"`' -X 'main.GOVERSION=`go version | cut -d" " -f 3`'"
endif

DATE=$(shell date --iso=seconds)
beta: build
	mkdir -p /rel/apps/identity_beta && cd /rel/apps/identity_beta/ && mkdir -p log
	cp config.beta.toml /rel/apps/identity_beta/config.toml
	-rm -f /rel/apps/identity_beta/identity_srv
	cp identity_srv /rel/apps/identity_beta/identity_srv
	supervisorctl restart identity_beta

sup:
	cp sup_identity.conf /etc/supervisord.d/
	supervisorctl update

test:
	go test ./... -count=1


prod: build
	mkdir -p /rel/apps/identity_prod && cd /rel/apps/identity_prod/ && mkdir -p log
	cp config.prod.toml /rel/apps/identity_prod/config.toml
	-rm -f /rel/apps/identity_prod/identity_srv
	cp identity_srv /rel/apps/identity_prod/identity_srv
	supervisorctl restart identity_prod
