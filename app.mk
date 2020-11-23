GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
BINARY_NAME = satistuffed
MAKE:=make -f app.mk

include help.mk

.PHONY: install build test clean run

install:
	go mod download

build:
	$(GOBUILD) -o dist/$(BINARY_NAME) -v

build/win:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o dist/$(BINARY_NAME).exe -v

build/mac:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o dist/$(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
