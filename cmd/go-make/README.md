# go-make CLI tool

This is a tool to auto-create a simple Makefile (and a few other useful files)
that you can use as a starting point.  Most likely, you'll want to customise
things after that point.

## Usage

Assuming you have a go project in a folder somewhere underneath your `GOPATH`,
and that `$GOPATH/bin` is on your `PATH`, just go to the root of that project
and run this:

```
go get github.com/go-make/make/...
go-make init
```

By default, no existing files will be overwritten.

Of course, help is available from the tool (**Note**: please refer to the
actual command output to ensure you have the latest version):

```
$ go-make help
NAME:
   go-make - Create Makefile scaffolding for golang projects

USAGE:
   go-make [global options] command [command options] [arguments...]

COMMANDS:
     init, i      initialise scaffolding for this project
     manifest, m  list all files that can be generated
     help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

and for individual commands:

```
$ go-make help init
NAME:
   go-make init - initialise scaffolding for this project

USAGE:
   go-make init [command options] [files...]

DESCRIPTION:
   [files...] is a list of one/more files to generate. If not specified,
   all known files will be generated. By default, no existing files will
   be overwritten - but you can force generation of files with -f

OPTIONS:
   --force, -f  overwrites existing files - use with care!
```
