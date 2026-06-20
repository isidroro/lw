package lib

import (
	"fmt"
)

type File struct {
	Name      string
	Extension string
	Weight    int64
}

func PrintFile(f *File) string {
	return fmt.Sprintf("%s%s %d", f.Name, f.Extension, f.Weight)
}

// TODO: aplicar formateo de unidades
func formatSize(bytes int64) string {
	const (
		kb = 1024
		mb = 1024 * kb
		gb = 1024 * mb
	)

	switch {
	case bytes >= gb:
		return fmt.Sprintf("%.1fG", float64(bytes)/gb)
	case bytes >= mb:
		return fmt.Sprintf("%.1fG", float64(bytes)/mb)
	case bytes >= kb:
		return fmt.Sprintf("%.1fG", float64(bytes)/kb)
	default:
		return fmt.Sprintf("%dB", bytes)
	}
}
