.PHONY: build all deploy

build:
	go test ./... -count=1
	go build

all: build

deploy:
	git pull
	go build
	mkdir -p /mega/apps/gateway_http
	rm /mega/apps/gateway_http/gateway_http
	mv gateway_http /mega/apps/gateway_http/
	cp config.toml /mega/apps/gateway_http/
	supervisorctl restart gateway
