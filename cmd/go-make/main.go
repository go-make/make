package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func die(err error) {
	fmt.Print(err)
	os.Exit(1)
}

func main() {
	app := cli.NewApp()
	app.Name = "go-make"
	app.Version = ""
	app.Usage = "Create Makefile scaffolding for golang projects"
	app.Commands = []cli.Command{
		commandInit,
		commandManifest,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
