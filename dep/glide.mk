export GLIDE:=$(GOPATH)/bin/glide

ifndef NO_VENDOR
NO_VENDOR=$(shell $(GLIDE) nv)
install: $(GLIDE)
endif

#
#  Glide v0.12 no longer creates git repo under the vendor directory, which
#  makes things slower IMHO. I also sometimes find it useful to edit files
#  in vendored packages during debug, and it's less easy to see what you edited
#  if you can't just "git diff" in there.
#
GLIDE_VERSION:=v0.11.1

OS:=$(shell uname -s)
ifeq ("$(OS)","Darwin")
GLIDE_ARCH:=darwin-amd64
else
GLIDE_ARCH:=linux-amd64
endif

ifndef GLIDE_OPT_INSTALL
GLIDE_OPT_INSTALL:=--cache --cache-gopath
endif

$(GLIDE): | $(GOPATH)/bin
	$(call PROMPT,Installing $@)
	curl -sL https://github.com/Masterminds/glide/releases/download/$(GLIDE_VERSION)/glide-$(GLIDE_VERSION)-$(GLIDE_ARCH).tar.gz | tar -xz --to-stdout -f- $(GLIDE_ARCH)/glide > $@
	chmod +x $@

vendor glide.lock: $(GLIDE) glide.yaml
	$(call PROMPT,$@)
	$(GLIDE) install $(GLIDE_OPT_INSTALL)
	touch $@

glide.yaml: | $(GLIDE)
	$(GLIDE) init

