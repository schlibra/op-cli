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

all: current all-platform

all-platform: linux darwin windows android bsd other

windows: windows-amd64 windows-386 windows-arm64

linux: linux-amd64 linux-386 linux-arm64 linux-arm linux-riscv64 linux-ppc64 linux-ppc64le linux-s390x linux-loong64

darwin: darwin-amd64 darwin-arm64

bsd: freebsd openbsd netbsd

freebsd: freebsd-amd64 freebsd-386 freebsd-arm64 freebsd-arm freebsd-riscv64

openbsd: openbsd-amd64 openbsd-386 openbsd-arm64 openbsd-arm

netbsd: netbsd-amd64 netbsd-386 netbsd-arm64 netbsd-arm

other: dragonfly solaris illumos

dragonfly: dragonfly-amd64

solaris: solaris-amd64

illumos: illumos-amd64

android: android-amd64 android-386 android-arm android-arm64

current:
	@mkdir -p $(OUTPUT_DIR)
	go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-current$(BIN_EXT) $(SRC_DIR)

linux-amd64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-linux-amd64 $(SRC_DIR)

linux-386:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-linux-386 $(SRC_DIR)

linux-arm64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-linux-arm64 $(SRC_DIR)

linux-arm:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-linux-arm $(SRC_DIR)

linux-riscv64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=riscv64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-linux-riscv64 $(SRC_DIR)

linux-ppc64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=ppc64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-linux-ppc64 $(SRC_DIR)

linux-ppc64le:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=ppc64le go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-linux-ppc64le $(SRC_DIR)

linux-s390x:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=s390x go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-linux-s390x $(SRC_DIR)

linux-loong64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=loong64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-linux-loong64 $(SRC_DIR)

darwin-amd64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-darwin-amd64 $(SRC_DIR)

darwin-arm64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-darwin-arm64 $(SRC_DIR)

windows-amd64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-windows-amd64.exe $(SRC_DIR)

windows-386:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-windows-386.exe $(SRC_DIR)

windows-arm64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-windows-arm64.exe $(SRC_DIR)

freebsd-amd64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-freebsd-amd64 $(SRC_DIR)

freebsd-386:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-freebsd-386 $(SRC_DIR)

freebsd-arm64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=freebsd GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-freebsd-arm64 $(SRC_DIR)

freebsd-arm:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=freebsd GOARCH=arm go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-freebsd-arm $(SRC_DIR)

freebsd-riscv64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=freebsd GOARCH=riscv64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-freebsd-riscv64 $(SRC_DIR)

openbsd-amd64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=openbsd GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-openbsd-amd64 $(SRC_DIR)

openbsd-386:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=openbsd GOARCH=386 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-openbsd-386 $(SRC_DIR)

openbsd-arm64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=openbsd GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-openbsd-arm64 $(SRC_DIR)

openbsd-arm:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=openbsd GOARCH=arm go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-openbsd-arm $(SRC_DIR)

netbsd-amd64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=netbsd GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-netbsd-amd64 $(SRC_DIR)

netbsd-386:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=netbsd GOARCH=386 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-netbsd-386 $(SRC_DIR)

netbsd-arm64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=netbsd GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-netbsd-arm64 $(SRC_DIR)

netbsd-arm:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=netbsd GOARCH=arm go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-netbsd-arm $(SRC_DIR)

dragonfly-amd64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=dragonfly GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-dragonfly-amd64 $(SRC_DIR)

solaris-amd64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=solaris GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-solaris-amd64 $(SRC_DIR)

illumos-amd64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=illumos GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-illumos-amd64 $(SRC_DIR)

android-amd64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-android-amd64 $(SRC_DIR)

android-386:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-android-386 $(SRC_DIR)

android-arm64:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-android-arm64 $(SRC_DIR)

android-arm:
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(APP_NAME)-android-arm $(SRC_DIR)

clean:
	rm -rf $(OUTPUT_DIR)