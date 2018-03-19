
# Go parameters
BINARY_NAME=ghutils

all: test

build:
	go build -v -o ./bin/$(BINARY_NAME)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/$(BINARY_NAME) 

test:
	go test -v ./...

clean:
	go clean
	rm -f ./bin/$(BINARY_NAME)

deps:
	go get -u github.com/tools/godep
	godep restore