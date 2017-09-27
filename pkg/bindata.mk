export GO_BINDATA:=$(GOPATH)/bin/go-bindata

$(GO_BINDATA): | $(GOPATH)
	$(call PROMPT,Installing $@)
	$(GO) get -u github.com/jteeuwen/go-bindata/...

clean-tools::
	rm -f $(GO_BINDATA)
