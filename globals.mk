ifndef GO
export GO:=go
endif

#
# From any make rule, you can use this to let you know what's
# running .. helps to give more understandable console output.
#
# foo:
#     $(call PROMPT,$@)
#
define PROMPT
	@echo
	@echo "**********************************************************"
	@echo "*"
	@echo "*     $(1)"
	@echo "*"
	@echo "**********************************************************"
	@echo
endef

# for ldflags, see "go tool link -h"
STRIP_DEBUG:=-s -w

$(GOPATH)/bin:
	$(call PROMPT,mkdir $@)
	mkdir -p $@

# this allows you to run
#
#     $ eval `make gopath`
#
.PHONY: gopath
gopath:
	@echo export GOPATH=$(GOPATH)
	@echo export PATH=$(subst ::,:,$(GOPATH)/bin:$(subst $(GOPATH)/bin,,$(PATH)))
