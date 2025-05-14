package util

import (
	"log"
	"os"
)

func IsRepo(dir string) bool {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		log.Fatal("No repository has been initiated.")
	}

	return info.IsDir()
}
