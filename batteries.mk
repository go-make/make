DIR_GOMAKE:=$(dir $(lastword $(MAKEFILE_LIST)))
include $(DIR_GOMAKE)/gotools.mk

.PHONY: install
install: vendor
	$(call PROMPT,$@)
	$(GO) install $$($(GLIDE) nv)

vendor glide.lock: $(GLIDE) glide.yaml
	$(call PROMPT,$@)
	$(GLIDE) install $(GLIDE_OPT_INSTALL)
	touch $@

.PHONY: lint
lint: $(GOMETALINTER)
	$(call PROMPT,$@)
	$(GOMETALINTER) $(GOMETALINTER_OPT) $$($(GLIDE) nv)

.PHONY: lint-full
lint-full: $(GOMETALINTER)
	$(call PROMPT,$@)
	$(GOMETALINTER) $(GOMETALINTER_OPT) --enable-all $$($(GLIDE) nv)

.PHONY: lint-fast
lint-fast: $(GOMETALINTER)
	$(call PROMPT,$@)
	$(GOMETALINTER) $(GOMETALINTER_OPT) --fast $$($(GLIDE) nv)

.PHONY: test test-full
test: test-full
test-full:
	$(call PROMPT,$@)
	$(GO) test -v -race $$($(GLIDE) nv)

.PHONY: test-short
test-short:
	$(call PROMPT,$@)
	$(GO) test -v -short $$($(GLIDE) nv)

.PHONY: coverage coverage-full
coverage: coverage-full
coverage-full: $(GOCOV) $(GOCOV_HTML)
	$(call PROMPT,$@)
	$(GOCOV) test -v -race $$($(GLIDE) nv) > coverage.json
	$(GOCOV_HTML) coverage.json > coverage.html

.PHONY: coverage-short
coverage-short: $(GOCOV) $(GOCOV_HTML)
	$(call PROMPT,$@)
	$(GOCOV) test -v -short $$($(GLIDE) nv) > coverage.json
	$(GOCOV_HTML) coverage.json > coverage.html

.PHONY: clean
clean::

.PHONY: clobber
clobber:: clean
