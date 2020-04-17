export GO_BINDATA:=$(GOPATH)/bin/go-bindata

$(GO_BINDATA): | $(GOPATH)
	$(call PROMPT,Installing $@)
	$(GO) get github.com/go-bindata/go-bindata/...

tools:: $(GO_BINDATA)

clean-tools::
	rm -f $(GO_BINDATA)

update-tools::
	$(GO) get -u github.com/go-bindata/go-bindata/...
