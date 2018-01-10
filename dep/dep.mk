export DEP:=$(GOPATH)/bin/dep

$(DEP): | $(GOPATH)/bin
	$(call PROMPT,Installing $@)
	$(GO) get github.com/golang/dep/cmd/dep

tools:: $(DEP)

clean-tools::
	rm -f $(DEP)

update-tools:: | $(GOPATH)/bin
	$(GO) get -u github.com/golang/dep/cmd/dep

vendor Gopkg.lock: $(DEP) Gopkg.toml
	$(call PROMPT,$@)
	$(DEP) ensure
	touch $@

.PHONY: dep-update
dep-update: $(DEP) Gopkg.toml
	$(call PROMPT,$@)
	$(DEP) ensure
	$(DEP) ensure -update

Gopkg.toml: | $(DEP)
	[ -f $@ ] || $(DEP) init -no-examples

