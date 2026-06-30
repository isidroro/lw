package main

import (
	"fmt"
	"lw/lib"
	"os"
	"path/filepath"
)

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	var err error
	dir, err = filepath.Abs(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var root *lib.Node
	root, err = lib.GetTree(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	lib.PrintTree(root)
}
