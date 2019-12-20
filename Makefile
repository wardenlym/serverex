.PHONY: all dev infra status deploy

all: dev

dev:
	cd apps/excalibur && make
	./apps/excalibur/gate_srv/gate_srv -c=./apps/excalibur/conf/config.dev.toml

infra:
ifeq ($(PLATFORM), windows)
	cmd /C start /B ./infra/mongodb/bin/mongod.exe --dbpath ./infra/mongodb/db/ --logpath ./infra/mongodb/log/mongod.log
else
	./infra/mongodb/bin/mongod --dbpath ./infra/mongodb/db/ --logpath ./infra/mongodb/log/mongod.log &
endif

deploy:
	git pull && cd apps/excalibur && make beta