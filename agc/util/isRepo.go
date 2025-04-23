package util

import (
	"log"
	"os"

	"github.com/Hamidspirit/a-git-clone/agc"
)

func IsRepo() bool {
	info, err := os.Stat(agc.GitRepo)
	if os.IsNotExist(err) {
		log.Fatal("No repository has been initiated.")
	}

	return info.IsDir()
}
