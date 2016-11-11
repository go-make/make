ifndef GO
export GO:=go
endif

# this allows you to run
#
#     $ eval `make gopath`
#
.PHONY: gopath
gopath:
	@echo export GOPATH=$(GOPATH)
	@echo export PATH=$(subst ::,:,$(GOPATH)/bin:$(subst $(GOPATH)/bin,,$(PATH)))
