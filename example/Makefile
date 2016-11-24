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
#  Note that this references gopkg.in/make.v0
#  which always equates to the master branch.
#  To make sure future updates don't break your
#  build, you should lock this to the current
#  version.  See http://gopkg.in for more info,
#  and https://gopkg.in/make.v0 for a list of
#  versions.
#
-include $(GOPATH)/src/gopkg.in/make.v0/batteries.mk
$(GOPATH)/src/gopkg.in/make.v0/batteries.mk:
	go get gopkg.in/make.v0
