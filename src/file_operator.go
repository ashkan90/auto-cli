package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
)

func gen(src []byte, fName, out string) error {
	wd, _ := os.Getwd()
	dst, err := os.Create(filepath.Join(wd, filepath.Base(out+fName)))
	if err != nil {
		return err
	}

	srcReader := bytes.NewReader(src)
	defer dst.Close()
	if _, err = io.Copy(dst, srcReader); err != nil {
		return err
	}

	return nil
}
