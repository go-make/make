export GODOC2MD:=$(GOPATH)/bin/godoc2md

$(GOPATH)/src/github.com/davecheney/godoc2md: $(DEP)
	$(call PROMPT,Cloning $@)
	mkdir -p $(GOPATH)/src/github.com/davecheney && \
		cd $(GOPATH)/src/github.com/davecheney && \
		git clone https://github.com/boyvinall/godoc2md

$(GODOC2MD): | $(GOPATH)/src/github.com/davecheney/godoc2md
	$(call PROMPT,Installing $@)
	cd $(GOPATH)/src/github.com/davecheney/godoc2md && \
		go install

tools:: $(GODOC2MD)

clean-tools::
	rm -f $(GODOC2MD)

update-tools::
	$(GO) get -u github.com/davecheney/godoc2md

