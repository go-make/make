export PROTOC:=$(GOPATH)/bin/protoc
export PROTOC_GEN_GO:=$(GOPATH)/bin/protoc-gen-go

OS:=$(shell uname -s)
ifeq ("$(OS)","Darwin")
PROTOC_ARCH:=osx-x86_64
else
PROTOC_ARCH:=linux-x86_64
endif

PROTOC_VERSION:=3.1.0
PROTOC_ZIP:=protoc-$(PROTOC_VERSION)-$(PROTOC_ARCH).zip

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
	unzip -j /tmp/$(PROTOC_ZIP) bin/protoc -d $(dir $@)

$(PROTOC_GEN_GO):
	$(call PROMPT,Installing $@)
	$(GO) get github.com/golang/protobuf/protoc-gen-go

%.pb.go: %.proto
	$(call PROMPT,Generating $@)
	$(PROTOC) -I$(dir $<) $(PROTOC_FLAGS) $< --go_out=plugins=grpc:$(dir $@)
