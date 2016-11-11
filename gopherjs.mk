export GOPHERJS:=$(GOPATH)/bin/gopherjs
export TEMPLE:=$(GOPATH)/bin/temple

$(GOPHERJS):
	$(GO) get github.com/gopherjs/gopherjs

$(TEMPLE):
	$(GO) get github.com/go-humble/temple

