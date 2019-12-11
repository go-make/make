DIR_GOMAKE:=$(dir $(lastword $(MAKEFILE_LIST)))

include $(DIR_GOMAKE)/globals.mk

ifneq ("$(GOPROXY)","")
GOMAKE_VENDOR:=modules
endif

GOMAKE_VENDOR?=dep
-include $(DIR_GOMAKE)/dep/$(GOMAKE_VENDOR).mk

define INCLUDE
ifndef GOMAKE_DISABLE_$(1)
# $$(info including $(1))
include $(DIR_GOMAKE)/pkg/$(1)
endif
endef

$(foreach pkg,$(wildcard $(DIR_GOMAKE)/pkg/*.mk),$(eval $(call INCLUDE,$(notdir $(pkg)))))
