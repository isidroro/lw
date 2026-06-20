package main

import (
	"fmt"
	"lw/lib"
	"os"
)

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	files := lib.GetFilesInDir(dir)

	for _, f := range files {
		fmt.Println(lib.PrintFile(&f))
	}
}
