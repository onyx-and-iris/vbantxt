PROGRAM = vbantxt

GO = go
BIN_DIR := bin

WINDOWS=$(BIN_DIR)/$(PROGRAM)_windows_amd64.exe
LINUX=$(BIN_DIR)/$(PROGRAM)_linux_amd64
VERSION=$(shell git describe --tags $(shell git rev-list --tags --max-count=1))

.DEFAULT_GOAL := build

.PHONY: fmt vet build windows linux test clean
fmt:        
	$(GO) fmt ./...

vet: fmt        
	$(GO) vet ./...

build: vet windows linux | $(BIN_DIR)
	@echo version: $(VERSION)

windows: $(WINDOWS)

linux: $(LINUX)


$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/$(PROGRAM)/

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/$(PROGRAM)/

test:
	$(GO) test ./...

$(BIN_DIR):
	@mkdir -p $@

clean:
	@rm -rv $(BIN_DIR)