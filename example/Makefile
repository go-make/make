.PHONY: fast
fast: build coverage-short lint-fast

.PHONY: all
all: build coverage lint-full

export GOPATH:=$(realpath $(shell pwd)/../../../..)
export GOROOT:=/usr/local/go

GOMETALINTER_LINELENGTH:=120
GOMETALINTER_DEADLINE:=300s
GOMETALINTER_CYCLO:=25

#
#  Note that this references the initial release,
#  which may not be current.
#
#  See https://gopkg.in/make.v1 for a list of
#  versions and http://gopkg.in for more info.
#
-include $(GOPATH)/src/gopkg.in/make.v1/batteries.mk
$(GOPATH)/src/gopkg.in/make.v1/batteries.mk:
	go get gopkg.in/make.v1
