DIR_GOMAKE:=$(dir $(lastword $(MAKEFILE_LIST)))

# output directory for test reports etc
DIR_OUT?=./out

include $(DIR_GOMAKE)/tools.mk

.PHONY: install
install::
	$(call PROMPT,$@)
	$(GO) install $(INSTALL_FLAGS) $(NO_VENDOR)

.PHONY: lint
lint: $(GOLANGCI_LINT)
	$(call PROMPT,$@)
	$(GOLANGCI_LINT) run

.PHONY: lint-full
lint-full: $(GOLANGCI_LINT)
	$(call PROMPT,$@)
	$(GOLANGCI_LINT) run --enable-all

.PHONY: lint-fast
lint-fast: $(GOLANGCI_LINT)
	$(call PROMPT,$@)
	$(GOLANGCI_LINT) run --fast

.PHONY: test test-full
test: test-full
test-full:
	$(call PROMPT,$@)
	GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) $(GO) test -v -race $(NO_VENDOR) $(TESTFLAGS)

.PHONY: test-short
test-short:
	$(call PROMPT,$@)
	GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) $(GO) test -v -short $(NO_VENDOR) $(TESTFLAGS)

.PHONY: coverage coverage-full
coverage: coverage-full
coverage-full: $(GOCOV) $(GOCOV_HTML) | $(DIR_OUT)
	$(call PROMPT,$@)
	GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) $(GOCOV) test -v -race $(NO_VENDOR) $(TESTFLAGS) > $(DIR_OUT)/coverage.json
	$(GOCOV_HTML) $(DIR_OUT)/coverage.json > $(DIR_OUT)/coverage.html
	$(GOCOV) report $(DIR_OUT)/coverage.json

.PHONY: coverage-short
coverage-short: $(GOCOV) $(GOCOV_HTML) | $(DIR_OUT)
	$(call PROMPT,$@)
	GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) $(GOCOV) test -v -short $(NO_VENDOR) $(TESTFLAGS) > $(DIR_OUT)/coverage.json
	$(GOCOV_HTML) $(DIR_OUT)/coverage.json > $(DIR_OUT)/coverage.html
	$(GOCOV) report $(DIR_OUT)/coverage.json

$(DIR_OUT):
	$(call PROMPT,Creating output directory)
	mkdir -p $@

# this is defined as a double-colon rule so that the calling makefile can clean whatever it needs to as well
.PHONY: clean
clean::
	rm -f $(DIR_OUT)/coverage.html $(DIR_OUT)/coverage.json
	$(GO) clean

.PHONY: clobber
clobber:: clean

# each makelet adds to this to install a full set of tools
# it should not be necessary to run this as required tools should be automatically installed
.PHONY: tools
tools::

# each makelet adds to this to clean themselves up
# it can be useful if you update your go version and want to rebuild the same set of tools you had before
.PHONY: clean-tools
clean-tools::

# each makelet adds to this one too.  can you guess what it does?
.PHONY: update-tools
update-tools::
