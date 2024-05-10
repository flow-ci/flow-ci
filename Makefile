WEB_APP = web
BUILD_DIR = $(PWD)/bin

.PHONY: test bench run-web

web:
	@go build -o $(BUILD_DIR)/$(WEB_APP) cmd/$(WEB_APP)/main.go

run-web: web
	@$(BUILD_DIR)/$(WEB_APP)

test:
	@go test -v -race ./...

bench:
	@go test -bench=. -benchmem ./...