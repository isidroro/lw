package lib

import (
	"strconv"
)

type File struct {
	Name      string
	Extension string
	Weight    int64 //bytes?
}

func PrintFile(f *File) string {
	weightStr := strconv.FormatInt(f.Weight, 10)
	return f.Name + f.Extension + " " + weightStr
}
