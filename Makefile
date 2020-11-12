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

up:
	make install
	docker-compose up

ssh/app:
	docker-compose run --rm app bash

db/migrate:
	docker-compose run --rm app migrate -database 'mysql://root@tcp(db)/satistuffed' -path db/migrations up

db/reset:
	docker-compose run --rm app migrate -database 'mysql://root@tcp(db)/satistuffed' -path db/migrations drop -f
	docker-compose run --rm app migrate -database 'mysql://root@tcp(db)/satistuffed' -path db/migrations up

db/seed:
	echo 'db seed'
