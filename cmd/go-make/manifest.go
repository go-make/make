package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var commandManifest = cli.Command{
	Name:    "manifest",
	Aliases: []string{"m"},
	Usage:   "list all files that can be generated",
	Action: func(c *cli.Context) error {
		for outfile := range fileTemplates {
			fmt.Println(outfile)
		}
		return nil
	},
}
