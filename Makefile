BIN_NAME := shipment-calculator
BIN_DIR ?= bin
GOOS ?= linux
GOARCH ?= amd64
IMAGE_NAME := shipment-calculator
IMAGE_TAG := latest

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

## go

.PHONY: go-build
go-build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BIN_DIR)/$(BIN_NAME) main.go

.PHONY: go-run
go-run:
	go run main.go

.PHONY: go-test
go-test:
	go test ./... -v -count=1

.PHONY: run
run: go-build
	./$(BIN_DIR)/$(BIN_NAME)

## Docker

.PHONY: docker-build
docker-build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

.PHONY: docker-run
docker-run: docker-build
	docker run --rm -p 8080:8080 $(IMAGE_NAME):$(IMAGE_TAG)
