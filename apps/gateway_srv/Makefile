.PHONY: build all deploy sup

build:
	go test ./... -count=1
	go build

all: build

sup:
	cp sup_gateway.conf /etc/supervisord.d/
	supervisorctl update

deploy:
	git pull
	go build
	mkdir -p /rel/apps/gateway
	mkdir -p /rel/apps/gateway/log
	-rm /rel/apps/gateway/gateway_srv
	mv gateway_srv /rel/apps/gateway/
	cp config.toml /rel/apps/gateway/
	supervisorctl restart gateway

useast:
	sed -i 's/excalibur.53site.com/excalibur.useast.53site.com/g' /rel/apps/gateway/config.toml
	supervisorctl restart gateway
