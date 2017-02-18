GO ?= go
GOPATH := $(CURDIR)/../../../..

all: install

build:
	GOPATH=$(GOPATH) $(GO) build

fmt:
	GOPATH=$(GOPATH) find . -name "*.go" | xargs gofmt -w -s

install:
	GOPATH=$(GOPATH) $(GO) install time_wrapper
	# unsupported, but seems to work for us https://bugs.debian.org/cgi-bin/bugreport.cgi?bug=717172
	strip bin/time_wrapper
