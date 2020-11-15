include help.mk

init:
	make build
	make install
	make db/reset
	make db/seed

build:
	docker-compose build

install:
	echo 'install'

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
