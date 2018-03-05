
# Go parameters
DOCKER_USERNAME=mchmarny
GCP_PROJECT_NAME=mchmarny-lab
BINARY_NAME=github-teams-utils

all: test
build:
	go build -o ./bin/$(BINARY_NAME) -v

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/$(BINARY_NAME) 

test:
	go test -v ./...

clean:
	go clean
	rm -f ./bin/$(BINARY_NAME)

run: build
	bin/$(BINARY_NAME) -org mchmarny

deps:
	go get -u github.com/tools/godep
	godep restore

gcr:
	gcloud container builds submit --tag gcr.io/$(GCP_PROJECT_NAME)/$(BINARY_NAME):latest .

run-docker:
	docker build -t $(BINARY_NAME) .
	docker run -p 8080:8080 $(BINARY_NAME):latest

dockerhub:
	docker build -t $(BINARY_NAME) .
	docker tag $(BINARY_NAME):latest $(DOCKER_USERNAME)/$(BINARY_NAME):latest
	docker push $(DOCKER_USERNAME)/$(BINARY_NAME):latest