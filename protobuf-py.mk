# support for accessing microservices from python
# you have to manually include this, it's not part of the default gotools.mk

PIP:=$(shell which pip)
ifeq ($(PIP),)
$(error Needs pip installed)
endif

PYTHON_VERSION:=$(shell python --version 2>&1 | sed 's,^Python \([23]..\).*$$,python\1,g')
PYTHON_PATH:=$(shell which python | sed 's,/bin/python,,g')
GRPC_PYTHON_TOOLS:=$(PYTHON_PATH)/lib/$(PYTHON_VERSION)/site-packages/grpc_tools/

info:
	echo $(PYTHON_VERSION)
	echo $(PYTHON_PATH)

$(GRPC_PYTHON_TOOLS):
	$(PIP) install grpcio-tools

%_pb2.py: %.proto
	$(call PROMPT,Generating $@)
	T=$$(mktemp -d) bash -c ' \
		$(call STRIP_GOGO,$<) > $$T/$(notdir $<) && \
		python -m grpc_tools.protoc \
			--proto_path=$$T \
			-I$(dir $<) \
			-I$(GOPATH)/src \
			$(PROTOC_FLAGS) \
			--python_out=$(dir $<) \
			--grpc_python_out=$(dir $<) \
			$$T/$(notdir $<) ; \
		rm -rf $$T \
	'
