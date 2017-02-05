ifndef GO
export GO:=go
endif

INDENT_0:=""
INDENT_1:="    "
INDENT_2:="$(INDENT_1)$(INDENT_1)"
INDENT:="$(or $(INDENT_$(MAKELEVEL)),"... $(INDENT_2)")"

#
# From any make rule, you can use this to let you know what's
# running .. helps to give more understandable console output.
#
# foo:
#     $(call PROMPT,$@)
#
define PROMPT
	@echo
	@echo "$(INDENT)**********************************************************"
	@echo "$(INDENT)*"
	@echo "$(INDENT)*   ($(notdir $(shell pwd)))    $(1)"
	@echo "$(INDENT)*"
	@echo "$(INDENT)**********************************************************"
	@echo
endef

# for ldflags, see "go tool link -h"
STRIP_DEBUG:=-s -w

$(GOPATH)/bin:
	$(call PROMPT,mkdir $@)
	mkdir -p $@

PATH_GOPATHBIN:=$(subst ::,:,$(GOPATH)/bin:$(subst $(GOPATH)/bin,,$(PATH)))

# this allows you to run
#
#     $ eval `make gopath`
#
.PHONY: gopath
gopath:
	@echo export GOPATH=$(GOPATH)
	@echo export PATH=$(PATH_GOPATHBIN)


OS:=$(shell uname -s)
ifeq ("Linux","$(OS)")
PREFIX:=/usr
endif
ifeq ("Darwin","$(OS)")
PREFIX:=/usr/local
endif
