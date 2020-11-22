GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
BINARY_NAME := satistuffed
MAKE:=make -f app.mk

include help.mk

.PHONY: install build test clean run

install:
	go mod download

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

build/win:
	GOOS=windows GOARCH=amd64 $(MAKE) build

build/mac:
	GOOS=darwin GOARCH=amd64 $(MAKE) build

build/linux:
	GOOS=linux GOARCH=amd64 $(MAKE) build

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
