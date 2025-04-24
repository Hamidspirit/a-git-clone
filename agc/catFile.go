package agc

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CatFile(objectHash string) {
	fp := filepath.Join(GitRepo, "objects", objectHash)
	file, err := os.ReadFile(fp)
	if err != nil {
		log.Fatal("Failed to read file: \n", err)
	}
	contents := string(file)
	parts := strings.SplitN(contents, "\x00", 2)
	// fmt.Println("did Null byte exist or nah: ", parts[0])

	fmt.Printf("Contents of file %s: \n %s", objectHash, parts[1])
}
