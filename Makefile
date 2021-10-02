MAIN_DIR := chat-websocket
BIN_NAME := $(MAIN_DIR)

.PHONY: build
build:
	GO111MODULE=on go build -o ./bin/$(BIN_NAME) ./cmd/$(MAIN_DIR)

.PHONY: run
run:
	GO111MODULE=on go run ./cmd/$(MAIN_DIR)

.PHONY: start
start:
	./bin/$(BIN_NAME)
