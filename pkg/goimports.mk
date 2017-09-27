GOIMPORTS:=$(GOPATH)/bin/goimports

$(GOIMPORTS):
	go get golang.org/x/tools/cmd/goimports

clean-tools::
	rm -f $(GOIMPORTS)
