# sre-bootcamp
One2N SRE Bootcamp 

## Prerequisite
- Go 1.24.2 or higher installed
- Install Goose for DB schema migrations 
`go install github.com/pressly/goose/v3/cmd/goose@latest`
- Docker
- make

## Local Setup
- Run Postgres in a docker container

```bash
docker run --name sre-postgres -e POSTGRES_USER=myuser -e POSTGRES_PASSWORD=mypassword -e POSTGRES_DB=students_db -p 5432:5432 -d postgres:latest
```
- Set Goose Environment variables
```bash
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgres://myuser:mypassword@localhost:5432/students_db?sslmode=disable
export GOOSE_MIGRATION_DIR=./internal/db/migrations
```
- Run DB schema migrations using goose `goose up`
- Start Application `make run`
