#------------------------------------------------
#
#	setup environment
#
#------------------------------------------------

export GOPATH:=$(realpath $(shell pwd)/{{.GoPathRel}})

#------------------------------------------------
#
#	standard rules
#
#------------------------------------------------

# The first target defined is the default if no target is
# specified on the command line.  Make sure this doesn't
# take too long to run, so that people will run it on every
# build.
.PHONY: fast
fast: build coverage-short lint-fast

# Also define the "full fat" rule that does everything
.PHONY: all build
all: build coverage lint-full

#
#  See https://gopkg.in/make.v4 for a list of
#  versions and http://gopkg.in for more info.
#
GOMAKE:=gopkg.in/make.v4
{{- if .Dep}}
GOMAKE_VENDOR:=dep
{{- end}}
-include $(GOPATH)/src/$(GOMAKE)/batteries.mk
$(GOPATH)/src/$(GOMAKE)/%:
	go get $(GOMAKE)/...

#------------------------------------------------
#
#	now to actually build stuff...
#
#------------------------------------------------

# The 'build' target is where you customise for your project.
# In this simple example, we'll just install

.PHONY: build
build: install

{{- if .Docker}}
#------------------------------------------------
#
#	docker support
#
#------------------------------------------------

GOMAKE_DOCKER:=gopkg.in/go-make/docker.v0
-include $(GOPATH)/src/$(GOMAKE_DOCKER)/Makefile
$(GOPATH)/src/$(GOMAKE_DOCKER)/%:
	go get $(GOMAKE_DOCKER)/...
{{- end}}
