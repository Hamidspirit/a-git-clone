package agc

import (
	"crypto/sha1"
	"fmt"
	"log"
	"os"
)

func Commit(msg string) {
	treehash := WriteTree()
	commit := fmt.Sprintf("tree %s\n\n%s", treehash, msg)
	commithash := fmt.Sprintf("%x", sha1.Sum([]byte(commit)))
	fmt.Printf("commit hash: %s\ncommit:\n\t %s", commithash, commit)

	err := SaveCommitObj(commit)
	if err != nil {
		log.Fatal("Failed to create commit object")
	}

}

func SaveCommitObj(data string) error {
	file, err := os.Create(GitRepo + "/Head")
	if err != nil {
		return err
	}
	file.Write([]byte(data))
	return nil
}
