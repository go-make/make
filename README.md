Makefile snippets (aka "makelets" :smile:) that I use for building golang projects.
Everyone likes their own, I like these :bowtie:  Just wanted a public
place to put them so I don't end up with scattered copies everywhere .. so this repo
is kind of intended just for my benefit, but maybe you'll find a use too.  If you'd
like something else included, then fork away and send a PR if you like.

There are a few different example projects which demonstrate usage.  I didn't include
them in here for two reasons: (1) so that you can easily see what's needed in your
own projects, (2) to keep this repo slightly simpler.  The examples:

- [github.com/go-make/example-simple](https://github.com/go-make/example-simple)

# Scaffolding generation

Along with the makelets, this repo also contains a [simple CLI tool](cmd/go-make)
to generate various scaffolding.  Further info on the CLI tool is available there.

# Features

## Make targets

It's definitely recommended that you use [autocomplete](docs/bash-completion.md) for make.

These targets are supported by [`batteries.mk`](batteries.mk):

| Target(s)                        | Description                                                                                                                                                                                                  |
|----------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `install`                        | Does a `go install`                                                                                                                                                                                          |
| `vendor`                         | Uses glide to install versioned packages to your `vendor` directory                                                                                                                                          |
| `lint-fast`, `lint`, `lint-full` | Uses golinter-ci to scan for possible nasties.<br>The different targets just enable different options, either `--fast`,<br>default, or `--enable-all`                                                       |
| `test-short`, `test`             | invokes `go test` on all local packages (it skips the vendored ones).<br>`test-short` will pass the `--short` option in case you have some tests<br>that take a long time to run (e.g. testing timeouts etc) |
| `coverage-short`, `coverage`     | uses `gocov` to invoke `go test` on all your local packages, generating code <br> coverage info. Also builds an HTML file that shows which source lines <br> are not covered by your tests                   |
| `clean::`                        | it's a [double-colon](https://www.gnu.org/software/make/manual/html_node/Double_002dColon.html) make rule so that it can be extended by your own makefile                                                    |
| `clobber::`                      | gets things really really clean, it will invoke the `clean` target first                                                                                                                                     |

and these are defined elsewhere, included above:

| Target(s) | Description                                                                                                                                |
|-----------|--------------------------------------------------------------------------------------------------------------------------------------------|
| `gopath`  | Allows you to run `eval $(make gopath)` from your bash shell, sets up `$GOPATH` in your environment and adds `$GOPATH/bin` to your `$PATH` |

## Macros etc

| Variable              | Description                                                                             |
|-----------------------|-----------------------------------------------------------------------------------------|
| `$(call PROMPT, ...)` | Use in a make rule, just prints a little banner heading with some text                  |
| `STRIP_DEBUG`         | Define options to pass to `go build` via `-ldflags` to strip debug for smaller binaries |

# Supported Tools

Most of the tool support is quite simple, in many cases not much more than
something to install the tool and a make variable to refer to the binary.
Others are more fully featured.  Brief info follows below.

- [Scaffolding generation](#scaffolding-generation)
- [Features](#features)
	- [Make targets](#make-targets)
	- [Macros etc](#macros-etc)
- [Supported Tools](#supported-tools)
	- [glide](#glide)
	- [goconvey](#goconvey)
	- [gocov](#gocov)
	- [goa](#goa)
	- [gopherjs](#gopherjs)
	- [gravitational/version](#gravitationalversion)
	- [docker](#docker)
	- [protobuf/grpc](#protobufgrpc)

## [glide](glide.mk)

There are a plethora of package managers / vendoring tools for go at the moment.
I currently use [glide](https://github.com/Masterminds/glide).

Defined variables:

| Variable               | Description                                                                     |
|------------------------|---------------------------------------------------------------------------------|
| `$(GLIDE)`             | Refers to the `glide` executable. You can depend on this to get glide installed |
| `$(GLIDE_OPT_INSTALL)` | options to be passed to `glide install` (used to install package dependencies)  |

## [goconvey](goconvey.mk)

The [goconvey](https://github.com/smartystreets/goconvey) makelet needs no other
integration in your main makefile. Just run as follows (from your repo):

````
$ make goconvey
````

And it should install itself and fire up a webpage on http://127.0.0.1:8080.
All your `go test` tests will be run on any code change, and you can
get browser notifications when a test fails.

## [gocov](gocov.mk)

[gocov](https://github.com/axw/gocov) provides a nice wrapper around `go test`
by simplifying the process of getting code coverage on a per-package basis.
The [gocov-html](https://github.com/matm/gocov-html) generates a nice HTML file
where you can easily see which source lines are not tested.

## [goa](goa.mk)

See [github.com/goadesign/goa](https://github.com/goadesign/goa)

| Variable    | Description                                                                   |
|-------------|-------------------------------------------------------------------------------|
| `$(GOAGEN)` | Refers to the `goagen` executable. You can depend on this to get it installed |

There are make targets defined for each goa generate command, see table below.  This can
be useful as part of a make process to ensure things are always up-to-date.

| Target(s)         | Description                          |
|-------------------|--------------------------------------|
| goagen-app        | Runs the goagen "app" command        |
| goagen-bootstrap  | Runs the goagen "bootstrap" command  |
| goagen-client     | Runs the goagen "client" command     |
| goagen-controller | Runs the goagen "controller" command |
| goagen-gen        | Runs the goagen "gen" command        |
| goagen-js         | Runs the goagen "js" command         |
| goagen-main       | Runs the goagen "main" command       |
| goagen-schema     | Runs the goagen "schema" command     |
| goagen-swagger    | Runs the goagen "swagger" command    |

## [gopherjs](gopherjs.mk)

golang in the browser #FTW. See [gopherjs](https://github.com/gopherjs/gopherjs)

| Variable      | Description                                                                                                                                       |
|---------------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| `$(GOPHERJS)` | Refers to the `gopherjs` executable. You can depend on this to get gopherjs installed                                                             |
| `$(TEMPLE)`   | Refers to the `temple` executable (used to enable go-templates for client-side markup generation). You can depend on this to get temple installed |

## [gravitational/version](version.mk)

Allows you to print the git tag/hash of the repo when your app runs. See
[github.com/gravitational/version](https://github.com/gravitational/version) or the [godocs](https://godoc.org/github.com/gravitational/version#pkg-index) for more info.

Use this define something like:

    $(BIN): | $(CMD_LINKFLAGS)
        $(GO) build -ldflags='$(call VERSION_LDFLAGS_VENDOR,github.com/example/myrepo)'

NB, this assumes you've vendored this package. If not, then omit the `_VENDOR` in the line above.

## [docker](docker.mk)

Part of a separate repo, so that you can use the docker makelets in non-go projects.
See [github.com/go-make/docker](https://github.com/go-make/docker)

## [protobuf/grpc](protobuf.mk)
