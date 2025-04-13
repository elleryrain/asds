### STAGE 1: Build ###
FROM golang:1.22.1-alpine3.19@sha256:fc5e5848529786cf1136563452b33d713d5c60b2c787f6b2a077fa6eeefd9114 AS builder
WORKDIR /build
ADD ./go.mod .
COPY . .
RUN go build -o server ./cmd/app/main.go


### STAGE 2: Run ###
FROM alpine:3.18@sha256:48d9183eb12a05c99bcc0bf44a003607b8e941e1d4f41f9ad12bdcc4b5672f86

ARG NATS_ADDR

ENV NATS_ADDR=${NATS_ADDR}

COPY --from=builder /build/server server
COPY --from=builder /build/templates/ templates/
ENTRYPOINT ["./server"]
