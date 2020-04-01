GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet

all: clean test build

.PHONY: build
build:
	$(GOBUILD) -v ./...

.PHONY: test
test:
	$(GOTEST) -count=1 ./...

clean:
	$(GOCLEAN) ./...

create-service:
	docker run -d -p 8082:8080 --name flying-saucer-service-gofss rvillablanca/flying-saucer-service:1.0

stop-service:
	docker stop flying-saucer-service-gofss || true

remove-service: stop-service
	docker rm flying-saucer-service-gofss || true

recreate-service: remove-service create-service

start-service:
	docker start flying-saucer-service-gofss
