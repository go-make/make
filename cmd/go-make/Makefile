export GOPATH:=$(shell go env GOPATH)

.PHONY: all
all: install coverage lint-full

install:: templates.bin.go

# templates.bin.go is generated, no point warning about that
GOMAKE:=github.com/go-make/make
-include $(GOPATH)/src/$(GOMAKE)/batteries.mk

templates.bin.go: $(GO_BINDATA) $(wildcard templates/*)
	$(call PROMPT,Generating $@)
	$(GO_BINDATA) -o $@ templates/

clobber::
	rm -f templates.bin.go
