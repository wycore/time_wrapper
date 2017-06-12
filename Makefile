GO ?= go
GOPATH := $(CURDIR)/../../../..
PACKAGES := $(shell GOPATH=$(GOPATH) go list ./... | grep -v /vendor/)

all: build

build:
	GOPATH=$(GOPATH) $(GO) build -ldflags "-X main.version=`cat VERSION`"

fmt:
	GOPATH=$(GOPATH) find . -name "*.go" | xargs gofmt -w -s

test:
	GOPATH=$(GOPATH) $(GO) test -cover $(PACKAGES)
	GOPATH=$(GOPATH) $(GO) vet $(PACKAGES)
