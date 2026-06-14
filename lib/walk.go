package lib

import (
	"io/fs"
	"path/filepath"
)

func GetFilesInDir(baseRoute string) []File {
	filepath.WalkDir(baseRoute, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		return nil
	})
	test := []File{}
	return test
}
