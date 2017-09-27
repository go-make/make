export DEP:=$(GOPATH)/bin/dep

$(DEP): | $(GOPATH)/bin
	$(call PROMPT,Installing $@)
	go get -u github.com/golang/dep/cmd/dep

clean-tools::
	rm -f $(DEP)

vendor Gopkg.lock: $(DEP) Gopkg.toml
	$(call PROMPT,$@)
	$(DEP) ensure
	touch $@

.PHONY: dep-update
dep-update: $(DEP) Gopkg.toml
	$(call PROMPT,$@)
	$(DEP) ensure -update

Gopkg.toml: | $(DEP)
	$(DEP) init -no-examples

