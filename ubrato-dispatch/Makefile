include .env
include ./proto/proto.mk

PROTO_PATH = ./proto
PROTO_OUT = ./internal/pb
PROTO_OUT_MODULE = gitlab.com/vrchub/event-api/internal/pb

.PHONY: install run docker_compose_up docker_compose_down docker_clean

setup:
	cp ./configs/.env.example .env
	go mod tody

docker_build:
	docker build -f ./deployments/Dockerfile -t ubrato-dispatch:dev .

docker_compose_up:
	docker compose up

docker_compose_down:
	docker compose down
	docker rmi effectivemobile-service-api

docker_clean:
	docker stop ubrato-dispatch
	docker rm ubrato-dispatch
	docker rmi ubrato-dispatch:dev

run:
	@EMAIL_LOGIN=$(EMAIL_LOGIN) EMAIL_FROM=$(EMAIL_FROM) EMAIL_ADDR=$(EMAIL_ADDR) EMAIL_PASS=$(EMAIL_PASS) EMAIL_HOST=$(EMAIL_HOST) go run cmd/app/main.go
