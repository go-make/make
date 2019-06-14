export GOLANGCI_LINT:=$(GOPATH)/bin/golangci-lint
export _SELF:=$(lastword $(MAKEFILE_LIST))
GOLANGCI_LINT_VERSION:=v1.16.0

# grab the gometalinter binary and install the actual linters
$(GOLANGCI_LINT): | $(GOPATH)
	$(call PROMPT,Installing $@)
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin $(GOLANGCI_LINT_VERSION)

tools:: $(GOLANGCI_LINT_VERSION)

clean-tools::
	rm -f $(GOLANGCI_LINT_VERSION)

update-tools::
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin $(GOLANGCI_LINT_VERSION)
