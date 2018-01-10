export GODOC2MD:=$(GOPATH)/bin/godoc2md

$(GODOC2MD): | $(GOPATH)
	$(call PROMPT,Installing $@)
	$(GO) get github.com/davecheney/godoc2md

tools:: $(GODOC2MD)

clean-tools::
	rm -f $(GODOC2MD)

update-tools::
	$(GO) get -u github.com/davecheney/godoc2md
