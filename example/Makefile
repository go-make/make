.PHONY: fast
fast: build coverage-short lint-fast

.PHONY: all
all: build coverage lint-full

export GOPATH:=$(realpath $(shell pwd)/../../../..)
export GOROOT:=/usr/local/go

GOMETALINTER_LINELENGTH:=120
GOMETALINTER_DEADLINE:=300s
GOMETALINTER_CYCLO:=25

-include $(GOPATH)/src/github.com/boyvinall/gomake/batteries.mk
$(GOPATH)/src/github.com/boyvinall/gomake/batteries.mk:
	go get github.com/boyvinall/gomake
