CMD_LINKFLAGS:=$(GOPATH)/bin/linkflags
$(CMD_LINKFLAGS): | $(GOPATH)
	$(call PROMPT,Installing $@)
	$(GO) get github.com/gravitational/version/cmd/linkflags

clean-tools::
	rm -f $(CMD_LINKFLAGS)

define VERSION_LDFLAGS
$(shell $(CMD_LINKFLAGS) -pkg=$(GOPATH)/src/$(1))
endef

define VERSION_LDFLAGS_VENDOR
$(addprefix -X $(strip $(1))/vendor/,$(filter-out -X,$(call VERSION_LDFLAGS,$(1))))
endef
