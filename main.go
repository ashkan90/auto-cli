package main

import (
	"context"
	"flag"
	"github.com/cristalhq/acmd"

	"os"
)

const (
	FILE_NAME = "node_template"
)

const (
	FLAG_PACKAGE   = "{{ package }}"
	FLAG_SOURCE    = "{{ core-source }}"
	FLAG_NODE_NAME = "{{ node-name }}"
)

var (
	pkg      *string
	source   *string
	nodeName *string
	out      *string
)

func init() {

	flag.Parse()
}

type generalFlags struct {
	IsVerbose bool
	Package   string
	Core      string
	NodeName  string
	Output    string
}

func (c *generalFlags) Flags() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.BoolVar(&c.IsVerbose, "verbose", false, "should app be verbose")
	fs.StringVar(&c.Package, "package", "main", "where this node will be located in local package. Default is 'main'")
	fs.StringVar(&c.Core, "core-source", "github.com/ashkan90/auto", "")
	fs.StringVar(&c.NodeName, "node-name", "", "whether the name of node")
	fs.StringVar(&c.Output, "out", "/", "where to locate generated file")

	return fs
}

func main() {
	var cfg generalFlags
	cmds := []acmd.Command{
		{
			Name:        "generate",
			Alias:       "gen",
			Description: "Generate blueprint node",
			ExecFunc: func(c context.Context, args []string) error {
				if err := cfg.Flags().Parse(args); err != nil {
					return err
				}

				genContent(cfg.Package, cfg.Core, cfg.NodeName)
				return nil
			},
		},
	}

	r := acmd.RunnerOf(cmds, acmd.Config{
		AppName:        "auto-cli",
		AppDescription: "Generate auto blueprint node",
		Version:        "v0.0.1",
		Output:         os.Stdout,
		Args:           os.Args,
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}
