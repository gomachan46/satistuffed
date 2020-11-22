include help.mk

APPMAKE:=make -f app.mk

.PHONY: init build up down db/migrate db/reset db/seed install

init:
	make build
	make install
	make db/reset
	make db/seed

build:
	docker-compose build

up: install
	docker-compose up

down:
	docker-compose down

db/migrate:
	docker-compose run --rm migrate make db/migrate

db/reset:
	docker-compose run --rm migrate make db/reset

db/seed:
	docker-compose run --rm migrate make db/seed

install:
	$(APPMAKE) install

clean:
	$(APPMAKE) app.mk clean