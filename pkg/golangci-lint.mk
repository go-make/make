export GOLANGCI_LINT:=$(GOPATH)/bin/golangci-lint
export _SELF:=$(lastword $(MAKEFILE_LIST))
GOLANGCI_LINT_VERSION:=v1.46.2

# grab the golangci binary and install the actual linters
$(GOLANGCI_LINT): | $(GOPATH)
	$(call PROMPT,Installing $@)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin $(GOLANGCI_LINT_VERSION)

tools:: $(GOLANGCI_LINT)

clean-tools::
	rm -f $(GOLANGCI_LINT)

update-tools:: update-golangci-lint

update-golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin $(GOLANGCI_LINT_VERSION)
