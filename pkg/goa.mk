export GOAGEN:=$(GOPATH)/bin/goagen

$(GOAGEN):
	$(call PROMPT,Installing $@)
	$(GO) get -u github.com/goadesign/goa/...

#------------------------------------------------
#
#	generate
#
#------------------------------------------------

define GOAGEN_COMMAND
.PHONY: goagen-$(1)
goagen-$(1): $(GOAGEN)
	$(call PROMPT,goagen $(1))
	$(GOAGEN) $(1) \
		--design=$(GOAGEN_DESIGN) \
		--out=$(GOAGEN_OUTPUT)
endef

GOAGEN_COMMAND_LIST:=\
	app \
	bootstrap \
	client \
	controller \
	gen \
	js \
	main \
	schema \
	swagger

# commands \
# help \
# version

$(foreach command,$(GOAGEN_COMMAND_LIST),$(eval $(call GOAGEN_COMMAND,$(command))))
