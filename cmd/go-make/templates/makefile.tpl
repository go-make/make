.PHONY: fast
fast: build coverage-short lint-fast

.PHONY: all
all: build coverage lint-full

export GOPATH:=$(realpath $(shell pwd)/{{.GoPathRel}})

-include $(GOPATH)/src/gopkg.in/make.v2/batteries.mk
$(GOPATH)/src/gopkg.in/make.v2/batteries.mk:
	go get gopkg.in/make.v2

.PHONY: build
build: install
