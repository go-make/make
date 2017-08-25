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
-include $(GOPATH)/src/$(GOMAKE)/batteries.mk
$(GOPATH)/src/$(GOMAKE)/%:
	$(GO) get $(GOMAKE)/...

#------------------------------------------------
#
#	now to actually build stuff...
#
#------------------------------------------------

# The 'build' target is where you customise for your project.
# In this simple example, we'll just install

.PHONY: build
build: install

install: $(GLIDE)

