Makefile snippets (aka "makelets" :smile:) that I use for building golang projects.
Everyone likes their own, I like these :bowtie:  Just wanted a public
place to put them so I don't end up with scattered copies everywhere .. so this repo
is kind of intended just for my benefit, but maybe you'll find a use too.

For full batteries-included simplicity, use in a makefile as follows
(make sure you replace the spaces with a tab).

(See [gopkg.in/make.v1](https://gopkg.in/make.v1) for a list of versions)

````
-include $(GOPATH)/src/gopkg.in/make.v1/batteries.mk
$(GOPATH)/src/gopkg.in/make.v1/batteries.mk:
	go get gopkg.in/make.v1
````

Included in this repo is complete working [Makefile](example/Makefile), just
drop that file in the root of your go project.

Take a look in that file to see the make rules included.  If you want more
control over things then you can instead include just the tool definitions:

````
-include $(GOPATH)/src/gopkg.in/make.v1/gotools.mk
$(GOPATH)/src/gopkg.in/make.v1/gotools.mk:
	go get gopkg.in/make.v1
````

# Features

## Make targets

It's definitely recommended that you use [autocomplete](docs/bash-completion.md) for make.

These targets are supported by `batteries.mk`:

| Target(s) | Description |
|----|----|
| `build` | Does a `go install` .. might get renamed to `install` |
| `vendor` | Uses glide to install versioned packages to your `vendor` directory |
| `lint-fast`, `lint`, `lint-full` | Uses gometalinter to scan for possible nasties. The different targets just enable different options, either `--fast`, default, or `--enable-all` |
| `test-short`, `test` | invokes `go test` on all local packages (it skips the vendored ones). `test-short` will pass the `--short` option in case you have some tests that take a long time to run (e.g. testing timeouts etc) |
| `coverage-short`, `coverage` | uses `gocov` to invoke `go test` on all your local packages, generating code coverage info. Also builds an HTML file that shows which source lines are not covered by your tests |
| `clean` | it's a double-colon make rule so that it can be extended by your own makefile |
| `clobber` | gets things really really clean, it will invoke the `clean` target first |

and these are defined elsewhere, included above:

| Target(s) | Description |
|-----------|-------------|
| `gopath` | Allows you to run `eval $(make gopath)` from your bash shell, sets up `$GOPATH` in your environment and adds `$GOPATH/bin` to your `$PATH` |

## Macros etc

| Variable             | Description                                                                     |
|----------------------|---------------------------------------------------------------------------------|
| `$(call PROMPT, ...)` | Use in a make rule, just prints a little banner heading with some text |
| `STRIP_DEBUG` | Define options to pass to `go build` via `-ldflags` to strip debug for smaller binaries |

# Supported Tools

## glide

There are a plethora of package managers / vendoring tools for go at the moment.
I currently use [glide](https://github.com/Masterminds/glide).

Defined variables:

| Variable             | Description                                                                     |
|----------------------|---------------------------------------------------------------------------------|
| `$(GLIDE)`             | Refers to the `glide` executable. You can depend on this to get glide installed |
| `$(GLIDE_OPT_INSTALL)` | options to be passed to `glide install`                                         |

## gometalinter

Superb wrapper to simplify running a whole bunch of linters on your code.
See [gometalinter](https://github.com/alecthomas/gometalinter)

This makelet defines a few settings which you can override before you include
gomake.

Defined variables:

| Variable             | Description                                                                     |
|----------------------|---------------------------------------------------------------------------------|
| `$(GOMETALINTER)` | Refers to the `gometalinter` executable. You can depend on this to get gometalinter installed |
| `$(GOMETALINTER_OPT)` | Sets up some default options for `gometalinter` (including the settings just below) |
| `$(GOMETALINTER_LINELENGTH)` | Sets the text line length used during linting |
| `$(GOMETALINTER_DEADLINE)` | Sets the a linting time deadline (some linters can take quite a while...) |
| `$(GOMETALINTER_CYCLO)` | Sets the [cyclomatic complexity](https://en.wikipedia.org/wiki/Cyclomatic_complexity) used during linting |

## goconvey

The [goconvey](https://github.com/smartystreets/goconvey) makelet needs no other
integration in your main makefile. Just run as follows (from your repo):

````
$ make goconvey
````

And it should install itself and fire up a webpage on http://127.0.0.1:8080.
All your `go test` tests will be run on any code change, and you can
get browser notifications when a test fails.

## gocov

[gocov](https://github.com/axw/gocov) provides a nice wrapper around `go test`
by simplifying the process of getting code coverage on a per-package basis.
The [gocov-html](https://github.com/matm/gocov-html) generates a nice HTML file
where you can easily see which source lines are not tested.

## rice
Useful to embed directories of static files into binaries.
See [go.rice](https://github.com/GeertJohan/go.rice)

| Variable             | Description                                                                     |
|----------------------|---------------------------------------------------------------------------------|
| `$(RICE)`            | Refers to the `rice` executable. You can depend on this to get rice installed |

## gopherjs
golang in the browser #FTW. See [gopherjs](https://github.com/gopherjs/gopherjs)

| Variable             | Description                                                                     |
|----------------------|---------------------------------------------------------------------------------|
| `$(GOPHERJS)`        | Refers to the `gopherjs` executable. You can depend on this to get gopherjs installed |
| `$(TEMPLE)`          | Refers to the `temple` executable (used to enable go-templates for client-side markup generation. You can depend on this to get temple installed |

## gravitational/version
Allows you to print the git tag/hash of the repo when your app runs. See
[github.com/gravitational/version](https://github.com/gravitational/version) or the [godocs](https://godoc.org/github.com/gravitational/version#pkg-index) for more info.

Use this define something like:

    $(BIN): | $(CMD_LINKFLAGS)
        $(GO) build -ldflags='$(call VERSION_LDFLAGS_VENDOR,github.com/example/myrepo)'

NB, this assumes you've vendored this package. If not, then omit the `_VENDOR` in the line above.
