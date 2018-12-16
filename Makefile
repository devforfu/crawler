GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
SOURCE_DIR=.
BINARY_DIR=./bin

sources := $(wildcard $(SOURCE_DIR)/*.go)
executables := $(patsubst $(SOURCE_DIR)/%.go, $(BINARY_DIR)/%, $(sources))

.PHONY: all

all: test build

$(BINARY_DIR)/%: $(SOURCE_DIR)/%.go
	$(GOBUILD) -o $@ $<

$(SOURCE_DIR)/src/**/*_test.go:
	$(GOTEST) -v $@

build: $(executables)

test:
	$(GOTEST) -v ./src/fetch
	$(GOTEST) -v ./src/findlinks

