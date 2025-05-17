BIN_NAME := shipment-calculator
BIN_DIR ?= bin
GOOS ?= linux
GOARCH ?= amd64
IMAGE_NAME := shipment-calculator
IMAGE_SERVER_NAME := $(IMAGE_NAME)-server
IMAGE_UI_NAME := $(IMAGE_NAME)-ui
IMAGE_TAG := latest

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

## go

.PHONY: go-build
go-build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BIN_DIR)/$(BIN_NAME) main.go

.PHONY: go-run-server
go-run-server:
	go run main.go

.PHONY: go-test
go-test:
	go test ./... -v -count=1

.PHONY: run-server
run-server: go-build
	./$(BIN_DIR)/$(BIN_NAME)

## Docker

.PHONY: docker-build-server
docker-build-server:
	docker build -t $(IMAGE_SERVER_NAME):$(IMAGE_TAG) .

.PHONY: docker-run-server
docker-run-server: docker-build-server
	docker run --rm -p 8080:8080 $(IMAGE_SERVER_NAME):$(IMAGE_TAG)

.PHONY: docker-run
docker-run:
	docker compose build --no-cache && \
	docker compose up

.PHONY: docker-deploy
docker-deploy:
	docker compose build --no-cache && \
	docker compose up -d
