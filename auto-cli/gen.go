package auto_cli

import (
	"bytes"
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

func GenContent(pkg, source, out, nodeName string) []byte {
	contents, err := os.ReadFile(FILE_NAME)
	if err != nil {
		return []byte{}
	}

	contents = bytes.ReplaceAll(contents, []byte(FLAG_PACKAGE), []byte(pkg))
	contents = bytes.ReplaceAll(contents, []byte(FLAG_SOURCE), []byte(source))
	contents = bytes.ReplaceAll(contents, []byte(FLAG_NODE_NAME), []byte(nodeName))

	_ = gen(contents, nodeName+".go", out)

	return contents
}
