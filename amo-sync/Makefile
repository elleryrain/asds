include .env
export

TOOLS_PATH=bin/tools
ogen=$(TOOLS_PATH)/ogen
goose=$(TOOLS_PATH)/goose

PROTO_PATH=./proto
PROTO_OUT=./internal/gen/pb
PROTO_OUT_MODULE=gitlab.ubrato.ru/ubrato/amo-sync/internal/gen/pb

include ./proto/proto.mk

POSTGRES_MIGRATIONS_PATH=./internal/store/postgres/migrations
POSTGRES_DSN="host=$(STORE_POSTGRES_HOST) port=$(STORE_POSTGRES_PORT) user=$(STORE_POSTGRES_USER) password=$(STORE_POSTGRES_PASSWORD) dbname=$(STORE_POSTGRES_DATABASE) sslmode=disable"

$(goose):
	GOBIN=`pwd`/$(TOOLS_PATH) go install github.com/pressly/goose/v3/cmd/goose@v3.22.1

.PHONY: setup
setup: $(goose)

.PHONY: setup
	protoc 

.PHONY: run
run:
	go run ./cmd/amo-sync

.PHONY: compose.up
compose.up:
	docker compose up -d

.PHONY: compose.down
compose.down:
	docker compose down

.PHONY: migrate.postgres.create
migrate.postgres.create:
	$(goose) create $(name) sql -dir $(POSTGRES_MIGRATIONS_PATH)

.PHONY: migrate.postgres.up
migrate.postgres.up:
	$(goose) postgres $(POSTGRES_DSN) up -dir $(POSTGRES_MIGRATIONS_PATH)

.PHONY: migrate.postgres.down
migrate.postgres.down:
	$(goose) postgres $(POSTGRES_DSN) down-to 0 -dir $(POSTGRES_MIGRATIONS_PATH)