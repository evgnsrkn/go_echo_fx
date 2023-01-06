BINARY_NAME=main
DIR=$(shell pwd)
GOPATH=$(shell go env GOPATH)

build: 
	go build -o bin/${BINARY_NAME} cmd/main.go
compile:
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows main.go

run:
	./bin/${BINARY_NAME}

live-reload:
	$(GOPATH)/bin/air

build_and_run: build run

clean:
	go clean
	rm -rf bin

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

lint:
	golangci-lint run --enable-all
