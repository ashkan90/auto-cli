package main

import (
	"bytes"
	"flag"
	"log"
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
	pkg = flag.String("package", "main", "where this node will be located in local package. Default is 'main'")
	source = flag.String("core-source", "github.com/ashkan90/...", "")
	nodeName = flag.String("node-name", "", "whether the name of node")
	out = flag.String("out", "/", "where to locate generated file")

	flag.Parse()
}

func main() {
	if *nodeName == "" {
		log.Println("Missing node name. You must enter node name")
		os.Exit(1)
	}

	if *out == "" {
		log.Println("Missing output path. You must enter output path")
		os.Exit(1)
	}

	contents, err := os.ReadFile(FILE_NAME)
	if err != nil {
		log.Println("CLI has occurred an problem while generating your node", err)
		os.Exit(1)
	}

	contents = bytes.ReplaceAll(contents, []byte(FLAG_PACKAGE), []byte(*pkg))
	contents = bytes.ReplaceAll(contents, []byte(FLAG_SOURCE), []byte(*source))
	contents = bytes.ReplaceAll(contents, []byte(FLAG_NODE_NAME), []byte(*nodeName))

	err = gen(contents, *nodeName+".go", *out)
	if err != nil {
		log.Println("CLI has occurred an problem while writing your node file", err)
		os.Exit(1)
	}

	os.Exit(0)
}
