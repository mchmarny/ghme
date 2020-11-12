APP_NAME         =ghme
RELEASE_VERSION  =v0.5.1
LINTER_VERSION   =v1.31.0

all: help

devenv: ## Sets up development envirnment
	go get -u github.com/go-bindata/go-bindata/...
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b $(go env GOPATH)/bin $(LINTER_VERSION)
	golangci-lint --version
.PHONY: devenv

tidy: ## Updates the go modules and vendors all dependancies 
	go mod tidy
	go mod vendor
.PHONY: tidy

load: clean tidy ## Runs uncompiled code for load
	go run *.go load --token $(GITHUB_TOKEN) --file test.json
.PHONY: load

query: clean tidy ## Runs uncompiled code for query
	go run *.go query --file test.json --value demo
.PHONY: query

build: clean tidy ## Builds binaries
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build \
		-ldflags "-X main.AppVersion=$(RELEASE_VERSION) -X main.AppName=$(APP_NAME)" \
		-mod vendor -o ../bin/$(APP_NAME) .
.PHONY: build

lint: clean ## Lints the entire project 
	golangci-lint run --timeout=3m
.PHONY: lint

tag: ## Creates release tag 
	git tag $(RELEASE_VERSION)
	git push origin $(RELEASE_VERSION)
.PHONY: tag

clean: ## Cleans bin and temp directories
	go clean
	rm -fr ./vendor
	rm -fr ./test
.PHONY: goclean

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk \
		'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help