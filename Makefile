POSTGRESQL_URL="postgres://root:root@localhost:5432/mysportapp?sslmode=disable"

build:
	go build .

local_env_up:
	docker-compose up -d

local_env_stop:
	docker-compose stop

local_env_destroy:
	docker-compose down

clean:
	rm -f mysportapp

run: build
	./mysportapp

start: clean build local_env_up db_migrate_up

db_migrate_up: local_env_up
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

db_migrate_down: local_env_up
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

db_migrate_force: local_env_up
	migrate -database ${POSTGRESQL_URL} -path db/migrations force 1