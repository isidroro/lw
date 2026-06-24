package lib

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// en desuso
func GetFilesInDir(baseRoute string) []File {
	var files []File
	filepath.WalkDir(baseRoute, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}
		ext := filepath.Ext(d.Name())
		name := strings.TrimSuffix(d.Name(), ext)
		files = append(files, File{Name: name, Extension: ext, Weight: info.Size()})
		return nil
	})
	return files
}

// en uso
func GetTree(baseRoute string) (*Node, error) {
	// shorthand para declarar y tomar referencia en una linea
	root := &Node{Name: filepath.Base(baseRoute), IsDir: true}
	nodes := map[string]*Node{baseRoute: root}

	err := filepath.WalkDir(baseRoute, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// primera iter
		if path == baseRoute {
			return nil
		}

		node := &Node{
			Name:  d.Name(),
			IsDir: d.IsDir(),
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			node.Size = info.Size()
		}

		parentPath := filepath.Dir(path)
		parent := nodes[parentPath]
		// sumar al padre sus nodos hijos
		parent.Children = append(parent.Children, node)

		if d.IsDir() {
			nodes[path] = node
		}
		return nil
	})
	return root, err
}
