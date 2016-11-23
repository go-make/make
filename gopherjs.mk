export GOPHERJS:=$(GOPATH)/bin/gopherjs
export TEMPLE:=$(GOPATH)/bin/temple

$(GOPHERJS):
	$(call PROMPT,Installing $@)
	$(GO) get github.com/gopherjs/gopherjs

$(TEMPLE):
	$(call PROMPT,Installing $@)
	$(GO) get github.com/go-humble/temple

