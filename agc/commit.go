package agc

import (
	"crypto/sha1"
	"fmt"
	"log"
	"os"
	"strings"
)

func Commit(msg string) {
	treehash := WriteTree()
	var commit string

	commit += fmt.Sprintf("tree %s\n", treehash)

	head, b := getHead()
	if b {
		for _, h := range head {
			parenthash := strings.Split(h, " ")
			commit += "parent" + parenthash[1] + "\n"
		}
	}
	commit += msg

	commithash := fmt.Sprintf("%x", sha1.Sum([]byte(commit)))
	fmt.Printf("commit hash: %s\ncommit:\n %s", commithash, commit)
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

func getHead() ([]string, bool) {
	head, err := os.ReadFile(GitRepo + "/Head")
	if err != nil {
		return []string{}, false
	}

	li := strings.Split(string(head), "\n")

	return li[:len(li)-1], true
}
