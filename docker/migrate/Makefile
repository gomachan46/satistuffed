include help.mk

db/migrate:
	 wait-for-it db:3306 -- migrate -database 'mysql://root@tcp(db)/satistuffed' -path ./migrations up

db/reset:
	wait-for-it db:3306 -- migrate -database 'mysql://root@tcp(db)/satistuffed' -path ./migrations drop -f
	make db/migrate

db/seed:
	echo 'db seed'
