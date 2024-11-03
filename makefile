program = vbantxt

GO = go
BIN_DIR := bin

WINDOWS=$(BIN_DIR)/$(program)_windows_amd64.exe
LINUX=$(BIN_DIR)/$(program)_linux_amd64
VERSION=$(shell git describe --tags --always --long --dirty)

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
	env GOOS=windows GOARCH=amd64 go build -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/vbantxt/

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/vbantxt/

test:
	$(GO) test ./...

$(BIN_DIR):
	@mkdir -p $@

clean:
	@rm -rv $(BIN_DIR)