export GOPATH:=$(realpath $(shell pwd)/../../../../..)

.PHONY: all
all: install coverage lint-full

install: bindata.go

-include $(GOPATH)/src/gopkg.in/make.v3/batteries.mk
$(GOPATH)/src/gopkg.in/make.v3/batteries.mk:
	go get gopkg.in/make.v3

bindata.go: $(GO_BINDATA) $(wildcard templates/*)
	$(GO_BINDATA) templates/

