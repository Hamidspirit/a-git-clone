package agc

import (
	"crypto/sha1"
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
		log.Fatal("Failed to get current working directory:", err)
	}

	treeOid := visit(root)
	fmt.Println("Root tree OID:", treeOid)
}

func visit(path string) string {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("Failed to read directory:", err)
	}

	var treeEntries []string

	for _, entry := range entries {
		if isIgnored(entry) {
			continue
		}

		fullPath := filepath.Join(path, entry.Name())
		// fmt.Println("full path: ", fullPath)

		if entry.IsDir() {
			subtreeOid := visit(fullPath)
			treeEntries = append(treeEntries, fmt.Sprintf("tree %s %s", subtreeOid, entry.Name()))
			fmt.Printf("\t tree %s %s \n", subtreeOid, entry.Name())
			continue
		}

		hashobjs := HashObject(fullPath, "blob", []string{})
		for _, item := range hashobjs {
			fmt.Printf("\t blob %s %s\n", item.ObjectID, entry.Name())
			treeEntries = append(treeEntries, fmt.Sprintf("blob %s %s", item.ObjectID, entry.Name()))

		}
	}

	// Join tree entries and simulate tree object hash
	treeContent := strings.Join(treeEntries, "\n")
	treeHash := fmt.Sprintf("%x", sha1.Sum([]byte(treeContent)))

	// Optional: Write tree content to a file
	treePath := filepath.Join(".agc", "objects", treeHash)
	err = os.WriteFile(treePath, []byte(treeContent), 0644)
	if err != nil {
		log.Fatal("Failed to write tree object:", err)
	}

	return treeHash
}

func isIgnored(d fs.DirEntry) bool {
	name := d.Name()

	for _, pattern := range ignorePatterns {
		if strings.HasPrefix(pattern, "*.") {
			// Match file extlnsion like *.log
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
