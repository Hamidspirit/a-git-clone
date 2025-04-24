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
	err = filepath.WalkDir(root, visit)
	if err != nil {
		log.Fatal("Failed to walk directory", err)
	}
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if isIgnored(d) {
		if d.IsDir() {
			return fs.SkipDir
		} else {
			return nil
		}
	}

	fmt.Println("  ", path, d.IsDir())
	return nil
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
