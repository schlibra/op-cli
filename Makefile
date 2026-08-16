APP_NAME	?= op-cli
SRC_DIR		?= .
OUTPUT_DIR	?= dist

ifeq ($(shell go env GOOS 2>/dev/null),windows)
    BIN_EXT := .exe
else
    BIN_EXT :=
endif

VERSION		?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT      ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
BUILD_TIME  ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')

LDFLAGS     := -s -w \
               -X 'op-cli/utils.Version=$(VERSION)' \
               -X 'op-cli/utils.GitCommit=$(COMMIT)' \
               -X 'op-cli/utils.BuildTime=$(BUILD_TIME)'

.PHONY: all

all: build build-all

build:
	@mkdir -p $(OUTPUT_DIR)
	#go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)$(BIN_EXT) $(SRC_DIR)

build-all: linux-amd64 linux-arm64 darwin-amd64 darwin-arm64 windows-amd64

linux-amd64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-linux-amd64 $(SRC_DIR)

linux-arm64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-linux-arm64 $(SRC_DIR)

darwin-amd64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-darwin-amd64 $(SRC_DIR)

darwin-arm64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-darwin-arm64 $(SRC_DIR)

windows-amd64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-windows-amd64.exe $(SRC_DIR)

# 清理产物
clean:
	rm -rf $(OUTPUT_DIR)