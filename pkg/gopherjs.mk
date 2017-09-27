export GOPHERJS:=$(GOPATH)/bin/gopherjs
export TEMPLE:=$(GOPATH)/bin/temple

$(GOPHERJS):
	$(call PROMPT,Installing $@)
	$(GO) get github.com/gopherjs/gopherjs

$(TEMPLE):
	$(call PROMPT,Installing $@)
	$(GO) get github.com/go-humble/temple

clean-tools::
	rm -f $(GOPHERJS) $(TEMPLE)

#
# Some generic JS tools below, not specifically gopherjs but useful in that context
#
ifeq ("Linux","$(OS)")
JS_PREFIX:=/usr
endif
ifeq ("Darwin","$(OS)")
JS_PREFIX:=/usr/local
endif

ifndef YARN
YARN:=$(JS_PREFIX)/bin/yarn
endif

ifndef GULP
GULP:=./node_modules/gulp/bin/gulp.js
endif

ifndef NPM
NPM:=$(JS_PREFIX)/bin/npm
endif

# either/both of these targets can be generated from yarn
yarn.lock node_modules: package.json $(YARN)
	$(call PROMPT,$@)
	$(YARN)
	touch $@

# make sure updated yarn.lock triggers re-install
node_modules: yarn.lock $(NPM)

package.json:
	$(error You'll need to manually write $@)

# if you need gulp, it should be run from local
$(GULP): node_modules

$(NPM) $(YARN):
	$(error Please install $@)
