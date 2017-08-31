export DEP:=$(GOPATH)/bin/dep

OS:=$(shell uname -s)
ifeq ("$(OS)","Darwin")
GLIDE_ARCH:=darwin-amd64
else
GLIDE_ARCH:=linux-amd64
endif

ifndef GLIDE_OPT_INSTALL
GLIDE_OPT_INSTALL:=--cache --cache-gopath
endif

$(DEP): | $(GOPATH)/bin
	$(call PROMPT,Installing $@)
	go get -u github.com/golang/dep/cmd/dep

vendor Gopkg.lock: $(DEP) Gopkg.toml
	$(call PROMPT,$@)
	$(DEP) ensure
	touch $@

Gopkg.toml: | $(DEP)
	$(DEP) init

