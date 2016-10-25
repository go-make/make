ifndef GO
GO:=go
endif

DIR_GOMAKE:=$(dir $(lastword $(MAKEFILE_LIST)))
include $(DIR_GOMAKE)/glide.mk
include $(DIR_GOMAKE)/rice.mk
include $(DIR_GOMAKE)/version.mk
include $(DIR_GOMAKE)/gometalinter.mk

# this allows you to run
#
#     $ eval `make gopath`
#
.PHONY: gopath
gopath:
	@echo export GOPATH=$(GOPATH)
	@echo export PATH=$(subst ::,:,$(GOPATH)/bin:$(subst $(GOPATH)/bin,,$(PATH)))
