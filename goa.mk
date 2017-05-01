export GOAGEN:=$(GOPATH)/bin/goagen

$(GOAGEN):
	$(call PROMPT,Installing $@)
	$(GO) get -u github.com/goadesign/goa/...
