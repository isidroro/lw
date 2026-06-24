package lib

import (
	"fmt"
	"strings"
)

type File struct {
	Name      string
	Extension string
	Weight    int64
}

// actualmente usado
// Children es array de punteros a nodos
type Node struct {
	Name     string
	IsDir    bool
	Size     int64
	Children []*Node
}

func PrintFile(f *File) string {
	weight := formatSize(f.Weight)
	return fmt.Sprintf("%s%s%s%s", f.Name, f.Extension, "  ", weight)
}

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
		return fmt.Sprintf("%.1fM", float64(bytes)/mb)
	case bytes >= kb:
		return fmt.Sprintf("%.1fK", float64(bytes)/kb)
	default:
		return fmt.Sprintf("%dB", bytes)
	}
}

func PrintTree(root *Node) {
	printNode(root, 0)
}

func printNode(n *Node, prof int) {
	indent := strings.Repeat("  ", prof)

	if n.IsDir {
		fmt.Printf("%s%s/\n", indent, n.Name)
	} else {
		fmt.Printf("%s%-20s %s\n", indent, n.Name, formatSize(n.Size))
	}

	for _, child := range n.Children {
		printNode(child, prof+1)
	}
}
