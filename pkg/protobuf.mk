export PROTOC:=$(GOPATH)/bin/protoc
export PROTOC_GEN_GO:=$(GOPATH)/bin/protoc-gen-go

# protoc needs protoc-gen-go to be on the PATH
ifeq ($(filter $(GOPATH)/bin,$(PATH)),)
export PATH:=$(PATH_GOPATHBIN)
endif

# pass this to --go_out to generate client/server hooks
PROTOC_PARAMS_GRPC:=plugins=grpc

OS:=$(shell uname -s)
ifeq ("$(OS)","Darwin")
PROTOC_ARCH:=osx-x86_64
else
PROTOC_ARCH:=linux-x86_64
endif

PROTOC_VERSION:=3.3.0
PROTOC_ZIP:=protoc-$(PROTOC_VERSION)-$(PROTOC_ARCH).zip

define CLEAN_GRPC
	rm -rf vendor/golang.org/x/net/context vendor/google.golang.org/grpc vendor/github.com/golang/protobuf
endef

NET_CONTEXT:=$(GOPATH)/src/golang.org/x/net/context
$(NET_CONTEXT):
	$(call PROMPT,go get $@)
	$(GO) get golang.org/x/net/context

GRPC:=$(GOPATH)/src/google.golang.org/grpc
$(GRPC): $(NET_CONTEXT)
	$(call PROMPT,go get $@)
	$(GO) get google.golang.org/grpc

$(PROTOC): $(GPRC) | $(PROTOC_GEN_GO) $(GOPATH)/bin
	$(call PROMPT,Downloading $@)
	[ -f /tmp/$(PROTOC_ZIP) ] || (curl -sL https://github.com/google/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC_ZIP) > /tmp/$(PROTOC_ZIP))
	unzip -oj /tmp/$(PROTOC_ZIP) bin/protoc -d $(dir $@)

$(PROTOC_GEN_GO):
	$(call PROMPT,Installing $@)
	$(GO) get github.com/golang/protobuf/protoc-gen-go

clean-tools::
	rm -f $(PROTOC) $(PROTOC_GEN_GO)

update-tools::
	$(GO) get -u github.com/golang/protobuf/protoc-gen-go
	$(MAKE) --always-make $(PROTOC)

# you might want to override this if you use gogo
PROTOC_GO_OUT?=--go_out=$(PROTOC_PARAMS):$(GOPATH)/src

%.pb.go: %.proto $(PROTOC)
	$(call PROMPT,Generating $@)
	$(PROTOC) -I$(dir $<) -I$(GOPATH)/src $(PROTOC_FLAGS) $< $(PROTOC_GO_OUT)

# used by C# and python tools to strip gogo-protobuf markup from .proto files
define STRIP_GOGO
grep -v gogo.proto $(1) | sed "s,\[(gogoproto.moretags).*\],,g"
endef

