.PHONY: build

build:
	go build -v ./cmd/apiserver

	test:

.PNONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build