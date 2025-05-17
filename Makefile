BIN_NAME := shipmentcalc
BIN_DIR ?= bin
GOOS ?= linux
GOARCH ?= amd64
ADDRESS ?= ":8080"

.PHONY: build
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BIN_DIR)/$(BIN_NAME) main.go

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

.PHONY: run
run:
	go run main.go $(ADDRESS)

.PHONY: test
test:
	go test ./... -v -count=1

.PHONY: start-server
start-server: build
	./$(BIN_DIR)/$(BIN_NAME) $(ADDRESS)
