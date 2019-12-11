export GOCOV:=$(GOPATH)/bin/gocov
export GOCOV_HTML:=$(GOPATH)/bin/gocov-html

$(GOCOV): | $(GOPATH)
	$(call PROMPT,Installing $@)
	$(GO) get github.com/axw/gocov/gocov

$(GOCOV_HTML): | $(GOPATH)
	$(call PROMPT,Installing $@)
	$(GO) get github.com/matm/gocov-html

tools:: $(GOCOV) $(GOCOV_HTML)

clean-tools::
	rm -f $(GOCOV) $(GOCOV_HTML)

update-tools::
	$(GO) get -u github.com/axw/gocov/gocov
	$(GO) get -u github.com/matm/gocov-html
