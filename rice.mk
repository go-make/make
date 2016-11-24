export RICE:=$(GOPATH)/bin/rice

$(RICE): | $(GOPATH)
	$(call PROMPT,Installing $@)
	git clone -b appendbox https://github.com/boyvinall/go.rice.git $(GOPATH)/src/github.com/GeertJohan/go.rice
	$(GO) install github.com/GeertJohan/go.rice/...
