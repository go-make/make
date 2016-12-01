export GO_BINDATA:=$(GOPATH)/bin/go-bindata

#
#  Currently using a fork of go.rice as I've just added some enhancements to suit my needs. Hoping to
#  get that upstreamed soon.
#
$(GO_BINDATA): | $(GOPATH)
	$(call PROMPT,Installing $@)
	$(GO) get -u github.com/jteeuwen/go-bindata/...
