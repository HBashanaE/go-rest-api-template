# Simple Makefile for a Go project

build:
	@go build -o bin/api cmd/api/main.go

test:
	@go test -v ./...

run: build
	@./bin/api

docker-build:
	docker build -t go-api-template .