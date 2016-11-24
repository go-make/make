CMD_LINKFLAGS:=$(GOPATH)/bin/linkflags
$(CMD_LINKFLAGS): | $(GOPATH)
	$(call PROMPT,Installing $@)
	$(GO) get github.com/gravitational/version/cmd/linkflags

#  Use this define something like:
#
#  $(BIN): | $(CMD_LINKFLAGS)
#      $(GO) build -ldflags='$(call VERSION_LDFLAGS,github.com/example/myrepo)'
#
define VERSION_LDFLAGS
$(shell $(CMD_LINKFLAGS) -pkg=$(GOPATH)/src/$(1))
endef

define VERSION_LDFLAGS_VENDOR
$(addprefix -X $(1)/vendor/,$(filter-out -X,$(shell $(CMD_LINKFLAGS) -pkg=$(GOPATH)/src/$(1))))
endef
