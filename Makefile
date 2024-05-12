WEB_APP = web
CLI_APP = cli
BUILD_DIR = $(PWD)/bin

.PHONY: test bench run-web run-cli all

all: web cli

web:
	go build -o $(BUILD_DIR)/$(WEB_APP) cmd/$(WEB_APP)/main.go

run-web: web
	$(BUILD_DIR)/$(WEB_APP)

cli:
	go build -o $(BUILD_DIR)/$(CLI_APP) cmd/$(CLI_APP)/main.go

run-cli: cli
	$(BUILD_DIR)/$(CLI_APP)	

test:
	go test -v -race ./...

bench:
	go test -bench=. -benchmem ./...