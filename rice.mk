export RICE:=$(GOPATH)/bin/rice

$(RICE): | $(GOPATH)
	$(GO) get github.com/GeertJohan/go.rice/...
