DIR_GOMAKE:=$(dir $(lastword $(MAKEFILE_LIST)))

include $(DIR_GOMAKE)/globals.mk
include $(DIR_GOMAKE)/glide.mk
include $(DIR_GOMAKE)/rice.mk
include $(DIR_GOMAKE)/version.mk
include $(DIR_GOMAKE)/gometalinter.mk
include $(DIR_GOMAKE)/gopherjs.mk
include $(DIR_GOMAKE)/gocov.mk
include $(DIR_GOMAKE)/goconvey.mk
include $(DIR_GOMAKE)/bindata.mk
include $(DIR_GOMAKE)/protobuf.mk
include $(DIR_GOMAKE)/protobuf-gogo.mk
include $(DIR_GOMAKE)/goa.mk

