.PHONY: test all

test:
	@go test -v ./...

cover:
	@go test -cover ./...

all:
	@go build .

