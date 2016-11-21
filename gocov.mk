export GOCOV:=$(GOPATH)/bin/gocov
export GOCOV_HTML:=$(GOPATH)/bin/gocov-html

$(GOCOV): | $(GOPATH)
	$(GO) get github.com/axw/gocov/gocov

$(GOCOV_HTML): | $(GOPATH)
	$(GO) get gopkg.in/matm/v1/gocov-html
