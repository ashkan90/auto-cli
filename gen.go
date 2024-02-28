package main

import (
	"bytes"
	"os"
)

func genContent(pkg, source, nodeName string) []byte {
	contents, err := os.ReadFile(FILE_NAME)
	if err != nil {
		return []byte{}
	}

	contents = bytes.ReplaceAll(contents, []byte(FLAG_PACKAGE), []byte(pkg))
	contents = bytes.ReplaceAll(contents, []byte(FLAG_SOURCE), []byte(source))
	contents = bytes.ReplaceAll(contents, []byte(FLAG_NODE_NAME), []byte(nodeName))

	_ = gen(contents, nodeName+".go", *out)

	return contents
}
