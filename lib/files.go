package lib

import (
	"fmt"
	"strconv"
)

type File struct {
	Name      string
	Extension string
	Weight    int64 //bytes?
}

func PrintFile(f *File) {
	weightStr := strconv.FormatInt(f.Weight, 10)
	fmt.Printf("%s", f.Name+f.Extension+" "+weightStr)
}
