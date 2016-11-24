CMD_LINKFLAGS:=$(GOPATH)/bin/linkflags
$(CMD_LINKFLAGS): | $(GOPATH)
	$(call PROMPT,Installing $@)
	$(GO) get github.com/gravitational/version/cmd/linkflags

define VERSION_LDFLAGS
$(shell $(CMD_LINKFLAGS) -pkg=$(GOPATH)/src/$(1))
endef

define VERSION_LDFLAGS_VENDOR
$(addprefix -X $(1)/vendor/,$(filter-out -X,$(shell $(CMD_LINKFLAGS) -pkg=$(GOPATH)/src/$(1))))
endef
