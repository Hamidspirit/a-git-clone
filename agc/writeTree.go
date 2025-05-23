package agc

import (
	"crypto/sha1"
	"encoding/hex"
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

func WriteTree() (treeHash string) {
	root, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current working directory:", err)
	}

	treeOid := visit(root)
	fmt.Println("Root tree OID:", treeOid)
	return treeOid
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
	treeHash := SaveTreeObject(treeContent)

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

func SaveTreeObject(data string) (treehash string) {
	final := []byte("tree\x00" + data)
	hasher := sha1.New()

	hasher.Write(final)

	treehash = hex.EncodeToString(hasher.Sum(nil))

	treepath := filepath.Join(GitRepo, "objects", treehash)
	file, err := os.Create(treepath)
	if err != nil {
		log.Fatal("failed to create tree file:", err)
	}
	defer file.Close()

	n, err := file.Write(final)
	if err != nil {
		log.Fatal("failed to write tree file", err)
	}

	if n != len(final) {
		log.Fatalf("Incomplete write: expected %d, wrote %d", len(final), n)
	}
	return treehash
}
