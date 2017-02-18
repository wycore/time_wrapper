GO ?= go
GOPATH := $(CURDIR)/../../../..
PACKAGES := $(shell GOPATH=$(GOPATH) go list ./... | grep -v /vendor/)

all: install

build:
	GOPATH=$(GOPATH) $(GO) build -ldflags "-X main.version=`cat VERSION`"

fmt:
	GOPATH=$(GOPATH) find . -name "*.go" | xargs gofmt -w -s

test:
	GOPATH=$(GOPATH) $(GO) test -cover $(PACKAGES)
	GOPATH=$(GOPATH) $(GO) vet $(PACKAGES)

install:
	GOPATH=$(GOPATH) $(GO) install time_wrapper
	# unsupported, but seems to work for us https://bugs.debian.org/cgi-bin/bugreport.cgi?bug=717172
	strip bin/time_wrapper
