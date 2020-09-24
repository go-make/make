package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

type fileDetails struct {
	templatefile string
}

//go:generate go-bindata -o templates.bin.go templates/

var fileTemplates = map[string]*fileDetails{
	".editorconfig": {
		templatefile: "templates/editorconfig.tpl",
	},
	".golangci.yml": {
		templatefile: "templates/golangci.yml.tpl",
	},
	".gitignore": {
		templatefile: "templates/gitignore.tpl",
	},
	".vscode/settings.json": {
		templatefile: "templates/vscode-settings.tpl",
	},
	"Makefile": {
		templatefile: "templates/makefile.tpl",
	},
}

type templateContext struct {
	GoPath    string
	GoPathRel string
	Docker    bool
}

func getTemplateContext(c *cli.Context) templateContext {
	ctx := templateContext{}
	ctx.GoPath = os.Getenv("GOPATH")

	pwd, err := os.Getwd()
	if err != nil {
		die(err)
	}

	rel, err := filepath.Rel(pwd, ctx.GoPath)
	if err != nil {
		die(err)
	}

	ctx.GoPathRel = rel
	ctx.Docker = c.Bool("docker")
	return ctx
}

func createFile(c *cli.Context, outfile string, f *fileDetails) {
	if _, err := os.Stat(outfile); err == nil {
		if !c.Bool("force") {
			// file exists, don't overwrite
			fmt.Fprintln(os.Stderr, outfile, " - already exists, not overwriting")
			return
		}
	}

	b, err := Asset(f.templatefile)
	if err != nil {
		die(err)
	}

	d := filepath.Dir(outfile)
	if d != "." {
		err = os.MkdirAll(d, 0755)
		if err != nil {
			die(err)
		}
	}

	tpl, err := template.New("t").Parse(string(b))
	if err != nil {
		die(err)
	}

	fmt.Fprintln(os.Stdout, outfile, " - creating")
	w, err := os.Create(outfile)
	if err != nil {
		die(err)
	}

	err = tpl.Execute(w, getTemplateContext(c))
	if err != nil {
		die(err)
	}
}

var commandInit = cli.Command{
	Name:      "init",
	Aliases:   []string{"i"},
	Usage:     "initialise scaffolding for this project",
	UsageText: "go-make init [command options] [files...]",
	Description: `[files...] is a list of one/more files to generate. If not specified,
   all known files will be generated. By default, no existing files will
   be overwritten - but you can force generation of files with -f`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "force, f",
			Usage: "overwrites existing files - use with care!",
		},
		cli.BoolFlag{
			Name:  "docker, d",
			Usage: "include support for docker containers",
		},
	},
	Action: func(c *cli.Context) error {
		args := c.Args()
		if len(args) > 0 {
			for _, outfile := range args {
				f := fileTemplates[outfile]
				if f == nil {
					fmt.Fprintln(os.Stderr, "unknown output file:", outfile)
					os.Exit(1)
				}
				createFile(c, outfile, f)
			}
		} else {
			for outfile, f := range fileTemplates {
				createFile(c, outfile, f)
			}
		}
		return nil
	},
}
