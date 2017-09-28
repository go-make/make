export GOCOV:=$(GOPATH)/bin/gocov
export GOCOV_HTML:=$(GOPATH)/bin/gocov-html

$(GOCOV): | $(GOPATH)
	$(call PROMPT,Installing $@)
	$(GO) get github.com/axw/gocov/gocov

$(GOCOV_HTML): | $(GOPATH)
	$(call PROMPT,Installing $@)
	$(GO) get gopkg.in/matm/v1/gocov-html

clean-tools::
	rm -f $(GOCOV) $(GOCOV_HTML)

update-tools::
	$(GO) get -u github.com/axw/gocov/gocov
	$(GO) get -u gopkg.in/matm/v1/gocov-html
