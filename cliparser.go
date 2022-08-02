package main

import (
	g "lojas.io/repo"

	"github.com/alecthomas/kong"
)

type Context struct {
	Debug bool
}

type LsCmd struct {
	Paths []string `arg:"" optional:"" name:"path" help:"Paths to list." type:"path"`
}

type RmCmd struct {
	Force     bool `help:"Force removal."`
	Recursive bool `help:"Recursively remove files."`

	Paths []string `arg:"" name:"path" help:"Paths to remove." type:"path"`
}

type RepoCmd struct {
	Readme bool   `help:"Creates the initial readme."`
	Branch string `help:"The default branch to use. If ommited will be 'main'." type:"string" default:"main"`
	Name   string `arg:"" name:"name" help:"The repo name to initialize." type:"string"`
}

func (r *RepoCmd) Run(ctx *Context) error {
	g.Git(r.Name, r.Branch, r.Readme)
	return nil
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Repo RepoCmd `cmd:"" help:"Manages repos."`
}

func cliparser() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
