BIN := sakura
VERSION := $$(make -s show-version)
CURRENT_REVISION := $(shell git rev-parse --short HEAD)
BUILD_LDFLAG := "-s -w -X main.revision=$(CURRENT_REVISION)"
GOBIN ?= $(shell go env GOPATH)/bin
export GO111MODULE=on

.PHONY: build
build:
	go build -ldflags=$(BUILD_LDFLAG) -o $(BIN) .

.PHONY: show-version
show-version: $(GOBIN)/gobump
	gobump show -r

$(GOBIN)/gobump:
	go install github.com/x-motemen/gobump/cmd/gobump@latest

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -rf $(BIN)
	go clean
