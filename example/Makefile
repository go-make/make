# The first target defined is the default if no target is
# specified on the command line.  Make sure this doesn't
# take too long to run, so that people will run it on every
# build.
.PHONY: fast
fast: build coverage-short lint-fast

# Also define the "full fat" rule that does everything
.PHONY: all
all: build coverage lint-full

# This assumes that this makefile is copied into something
# like $GOPATH/src/github.com/example/package
export GOPATH:=$(realpath $(shell pwd)/../../../..)

# probably not needed
export GOROOT:=/usr/local/go

# set some override options for gometalinter
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
