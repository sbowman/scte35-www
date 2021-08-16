PROJECT = scte35
VERSION ?= 0.0.1

REPO = github.com/sbowman/scte35-www
BUILD_TIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
UNAME := $(shell uname -s)
GOROOT := $(shell go env GOROOT)

GO_FILES = $(shell find . -path ./.idea -prune -o -type f -name '*.go' -print)

VERBOSITY ?= 2

default: $(PROJECT)

$(PROJECT): export GOOS = js
$(PROJECT): export GOARCH = wasm
$(PROJECT): $(GO_FILES)
	@go build -o html/${PROJECT}.wasm

html/wasm_exec.js:
	@cp "$(GOROOT)/misc/wasm/wasm_exec.js" html

.PHONY: docker
docker:
	@docker-compose up

run: $(PROJECT) html/wasm_exec.js docker
