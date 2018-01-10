export RICE:=$(GOPATH)/bin/rice

#
#  Currently using a fork of go.rice as I've just added some enhancements to suit my needs. Hoping to
#  get that upstreamed soon.
#
$(RICE): | $(GOPATH)
	$(call PROMPT,Installing $@)
	[ -d $(GOPATH)/src/github.com/GeertJohan/go.rice ] || git clone -b appendbox https://github.com/boyvinall/go.rice.git $(GOPATH)/src/github.com/GeertJohan/go.rice
	cd $(GOPATH)/src/github.com/GeertJohan/go.rice/rice ; $(GO) get -v .
	$(GO) install github.com/GeertJohan/go.rice/rice

tools:: $(RICE)

clean-tools::
	rm -f $(RICE)

