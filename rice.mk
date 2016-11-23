export RICE:=$(GOPATH)/bin/rice

$(RICE): | $(GOPATH)
	$(call PROMPT,Installing $@)
	$(GO) get github.com/boyvinall/go.rice/...
