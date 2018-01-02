export GOMETALINTER:=$(GOPATH)/bin/gometalinter

# grab the gometalinter binary and install the actual linters
$(GOMETALINTER): | $(GOPATH)
	$(call PROMPT,Installing $@)
	$(GO) get github.com/alecthomas/gometalinter
	$(GOMETALINTER) --install

tools:: $(GOMETALINTER)

clean-tools::
	rm -f $(GOMETALINTER)

update-tools::
	$(GO) get -u github.com/alecthomas/gometalinter
	$(GOMETALINTER) --install

GOMETALINTER_LINELENGTH?=100
GOMETALINTER_DEADLINE?=30s
GOMETALINTER_CYCLO?=15

GOMETALINTER_OPT_DEFAULT=\
	--exclude='should have comment or be unexported' \
	--exclude='should have comment \(or a comment on this block\) or be unexported' \
	--exclude='error return value not checked \(defer ' \
	--deadline=$(GOMETALINTER_DEADLINE) \
	--line-length=$(GOMETALINTER_LINELENGTH) \
	--cyclo-over=$(GOMETALINTER_CYCLO)

# protobuf-generated code fails "go lint" with a couple of things...
GOMETALINTER_OPT_PROTOBUF=\
	--exclude="don't use underscores in Go names; (func|var|const|type|method|struct field) [^ ]+ should be" \
	--exclude="\.pb.go:[0-9]+:[0-9]+:warning: context.Context should be the first parameter of a function" \
	--exclude="\.pb.go:[0-9]+:[0-9]+:warning: exported method [^ ]+ returns unexported type" \
	--exclude="\.pb.go:[0-9]+:[0-9]+:warning: should replace [^ ]+ \+= 1 with [^ ]+\+\+" \
	--exclude="\.pb.go:[0-9]+::warning: Errors unhandled." \
	--exclude="\.pb.go:[0-9]+::warning: cyclomatic complexity"

GOMETALINTER_OPT_GOA=\
	--exclude="github.com/goadesign/goa/design. imported but not used"

GOMETALINTER_OPT?=\
	$(GOMETALINTER_OPT_DEFAULT) \
	$(GOMETALINTER_OPT_PROTOBUF) \
	$(GOMETALINTER_OPT_GOA) \
	$(GOMETALINTER_OPT_EXTRA)
