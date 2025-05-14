package agc

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var ignorePatterns = []string{
	".agc",
	".git",
	"agc.exe",
}

func WriteTree() {
	root, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current working direcetory", err)
	}

	fmt.Println("list of subdirectories")
	tree := visit(root)
	if err != nil {
		log.Fatal("Failed to walk directory", err)
	}

	_ = visit(tree)
}

func visit(path string) (treeHash string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("Failed to visit path:", err)
	}

	var te []TreeEntry

	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())

		if isIgnored(entry) {
			continue
		}

		if entry.IsDir() {
			subtree := visit(fullPath)
			if err != nil {
				log.Fatal("Failed to visit sub-tree", err)
			}
			// Hash subtree
			hash := visit(subtree)
			te = append(te, TreeEntry{
				ObjectType: "tree",
				OID:        hash,
				Name:       entry.Name(),
			})
		} else {
			// this abomination is my fault
			p := ""
			ot := "blob"
			fp, oid := HashObject(p, ot, entry.Name())

			te = append(te, TreeEntry{
				ObjectType: "blob",
				OID:        oid,
				Name:       entry.Name(),
			})

			fmt.Println("  fp:", fp, " oid:", oid)
		}

	}
	// fmt.Println("  ", path, entry.IsDir())
	return treeHash
}

func isIgnored(d fs.DirEntry) bool {
	name := d.Name()

	for _, pattern := range ignorePatterns {
		if strings.HasPrefix(pattern, "*.") {
			// Match file extension like *.log
			if strings.HasSuffix(name, pattern[1:]) {
				return true
			}
		} else if pattern == name {
			// Exact file name match or dir
			return true
		}
	}

	return false
}
