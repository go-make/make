export GOPATH:=$(realpath $(shell pwd)/../../../../..)

.PHONY: all
all: install coverage lint-full

install:: templates.bin.go vendor

# templates.bin.go is generated, no point warning about that
GOMETALINTER_OPT_EXTRA:=--exclude=^templates.bin.go:
GOMAKE:=gopkg.in/make.v4
-include $(GOPATH)/src/$(GOMAKE)/batteries.mk
$(GOPATH)/src/$(GOMAKE)/%:
	go get $(GOMAKE)/...

templates.bin.go: $(GO_BINDATA) $(wildcard templates/*)
	$(call PROMPT,Generating $@)
	$(GO_BINDATA) -o $@ templates/

clobber::
	rm -f templates.bin.go
