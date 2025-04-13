include .env
export

TOOLS_PATH=bin/tools
ogen=$(TOOLS_PATH)/ogen
goose=$(TOOLS_PATH)/goose

PROTO_PATH=./proto
PROTO_OUT=./internal/models/gen/proto
PROTO_OUT_MODULE=gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto

include ./proto/proto.mk

POSTGRES_MIGRATIONS_PATH=./internal/store/postgres/migrations
POSTGRES_DSN="host=$(STORE_POSTGRES_HOST) port=$(STORE_POSTGRES_PORT) user=$(STORE_POSTGRES_USER) password=$(STORE_POSTGRES_PASSWORD) dbname=$(STORE_POSTGRES_DATABASE) sslmode=disable"

$(ogen):
	GOBIN=`pwd`/$(TOOLS_PATH) go install github.com/ogen-go/ogen/cmd/ogen@v1.5.0

$(goose):
	GOBIN=`pwd`/$(TOOLS_PATH) go install github.com/pressly/goose/v3/cmd/goose@v3.22.1

setup: $(ogen) $(goose)

.PHONY: run
run:
	go run ./cmd/core

.PHONY: compose.up
compose.up:
	docker compose up -d

.PHONY: compose.down
compose.down:
	docker compose down

.PHONY: generate.api
generate.api: $(ogen)
	$(ogen) --loglevel error --clean --config .ogen.yml --target ./api/gen ./api/openapi.yaml
	docker run --rm -v `pwd`:/spec redocly/cli:1.25.3 bundle ./api/openapi.yaml > ./api/bundle.yaml

.PHONY: migrate.postgres.create
migrate.postgres.create: $(goose)
	$(goose) create $(name) sql -dir $(POSTGRES_MIGRATIONS_PATH)

.PHONY: migrate.postgres.up
migrate.postgres.up: $(goose)
	$(goose) postgres $(POSTGRES_DSN) up -dir $(POSTGRES_MIGRATIONS_PATH)

.PHONY: migrate.postgres.down
migrate.postgres.down: $(goose)
	$(goose) postgres $(POSTGRES_DSN) down-to 0 -dir $(POSTGRES_MIGRATIONS_PATH)