.PHONY: build
build:
	go build -v ./cmd/server

.PHONY: test
test:
	go test -v -race -timeout 30s ./...
.PHONY: migrate
DATABASE ?= go_market
ACTION ?= up
migrate:
	migrate -path migrations -database "mysql://root:root@tcp(db:3306)/$(DATABASE)" $(ACTION)
.DEFAULT_GOAL := build