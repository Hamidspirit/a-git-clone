package util

import (
	"log"
	"os"
	"path/filepath"
)

// returns a string containing path that is command is run
func PathParser(pathflag string) (path string) {

	if path == "." {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal("Failed to get working directory", err)
		}
		return wd
	}

	return pathflag
}

// returns path plus a name might be file name or another dir
func FilePathParser(path, name string) (fpath string) {
	var fp string
	if name != "" {
		fp = filepath.Join(path, name)
	} else {
		fp = path
	}

	return fp
}

// returns the name of a file from list os OS args if exits
func ExtractName(args []string) (name []string) {
	var filenames []string

	for _, item := range args {

		info, err := os.Stat(item)
		if err != nil {
			// skip if file doesn't exist or can't be accessed
			continue
		}

		if !info.IsDir() {
			filenames = append(filenames, item)
		}
	}

	return filenames
}
