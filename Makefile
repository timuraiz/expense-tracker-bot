.PHONY:
include .env
export

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

run-db:
	docker run -d --name expense-tracker-bot \
                  -e POSTGRES_HOST=$(HOST) \
                  -e POSTGRES_PORT=$(PORT) \
                  -e POSTGRES_USER=$(USER) \
                  -e POSTGRES_PASSWORD=$(PASSWORD) \
                  -e POSTGRES_DB=$(DBNAME) \
                  -p $(PORT):$(PORT) \
                  postgres:latest

open-psql:
	docker exec -it expense-tracker-bot psql -U $(USER) -d $(DBNAME)

stop-db:
	docker stop expense-tracker-bot && docker rm expense-tracker-bot