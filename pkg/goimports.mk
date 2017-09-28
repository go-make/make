GOIMPORTS:=$(GOPATH)/bin/goimports

$(GOIMPORTS):
	$(GO) get golang.org/x/tools/cmd/goimports

clean-tools::
	rm -f $(GOIMPORTS)

update-tools::
	$(GO) get -u golang.org/x/tools/cmd/goimports
