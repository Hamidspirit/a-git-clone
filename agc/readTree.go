package agc

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Hamidspirit/a-git-clone/util"
)

func ReadTree(treeHash string, purged bool) {
	if !purged {
		emptyCurrentDir()
		purged = true
	}
	objType, treedata := CatFile(treeHash)
	if objType != "tree" {
		log.Fatal("expected tree object got this:", objType)
	}

	entries, err := ParseTree(treedata)
	if err != nil {
		log.Fatal("failed to parse tree data: ", err)
	}

	for _, entry := range entries {
		if entry.Type == "blob" {
			blobType, blobData := CatFile(entry.Hash)
			if blobType != "blob" {
				log.Fatal("expected blob: ", blobType)
			}
			// Write to file
			if err := os.WriteFile(entry.Name, []byte(blobData), 0644); err != nil {
				log.Fatal("failed to write blob file: ", err)
			}
		} else if entry.Type == "tree" {
			wd, err := os.Getwd()
			if err != nil {
				log.Fatal("Failed to get woeking die: ", err)
			}
			p := util.FilePathParser(wd, entry.Name)
			// Ensure directory exists (in case path includes directories)
			if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
				log.Fatal("failed o create directory: ", err)
			}

			// recurse
			ReadTree(entry.Hash, true)
		} else {
			log.Fatal("unknown entry type: ", entry.Type)
		}
	}
}

func ParseTree(data string) ([]TreeEntry, error) {
	var entries []TreeEntry
	lines := strings.Split(strings.TrimSpace(data), "\n")

	for _, line := range lines {
		parts := strings.SplitN(line, " ", 3)
		if len(parts) > 3 {
			return nil, fmt.Errorf("invalid tree entry: %s", line)
		}

		entries = append(entries, TreeEntry{
			Type: parts[0],
			Hash: parts[1],
			Name: parts[2],
		})
	}
	return entries, nil

}

func emptyCurrentDir() error {
	return filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// skip root itself
		if path == "." {
			return nil
		}

		// convert to relative path for consistency
		// relpath , err := filepath.Rel(".", path)
		// if err != nil {
		// 	return nil
		// }

		if isIgnored(d) {
			if d.IsDir() {
				return filepath.SkipDir // dont decent into ignored dir
			}
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return nil
		}

		if !d.IsDir() && info.Mode().IsRegular() {
			return os.Remove(path)
		}

		// if it si a directory
		if d.IsDir() {
			err := os.Remove(path)
			if err != nil && !os.IsNotExist(err) {
				// ignore if dir is non empty or error
				return nil
			}
		}
		return nil

	})
}
