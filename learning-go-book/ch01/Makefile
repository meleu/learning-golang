.DEFAULT_GOAL := build

.PHONY:fmt vet build

test:
	go test ./...

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

clean:
	go clean
