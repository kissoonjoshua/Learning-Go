.DEFAULT_GOAL := build

.PHONY: build test fmt vet clean
fmt:
	go fmt ./...
	
vet: fmt
	go vet ./...

build: vet test
	go build ./...

test:
	go test ./...
	
clean:
	go clean