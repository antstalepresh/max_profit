SHELL=/bin/bash

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: build
build:
	CGO_ENABLED=0 go build -a -o ./build/max_profit cmd/main.go

.PHONY: run
run:
	go run cmd/main.go
