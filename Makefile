develop: clean build 

run: migrate.up
	docker compose up backend

build:
	docker compose build backend	

clean:
	docker compose rm -vf 

db.start:
	docker compose up -d postgres-db
	docker compose exec postgres-db \
		 sh -c 'while ! pg_isready -q; do sleep 0.1; done'

migrate.up:db.start
	docker compose --profile migrations run --rm migrate up
	