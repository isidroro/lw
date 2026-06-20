package lib

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func GetFilesInDir(baseRoute string) []File {
	var files []File
	filepath.WalkDir(baseRoute, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		// TODO: implementar recursividad
		if d.IsDir() {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		ext := filepath.Ext(d.Name())

		// Restamos la extension para guardar nombre y ext por separado
		name := strings.TrimSuffix(d.Name(), ext)
		files = append(files, File{Name: name, Extension: ext, Weight: info.Size()})
		return nil
	})
	return files
}
