package agc

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/Hamidspirit/a-git-clone/util"
)

func Init(path string) {
	fp := util.FilePathParser(path, GitRepo)

	err := os.Mkdir(fp, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	// Create object directory
	err = os.Mkdir(GitRepo+"/objects", 0750)
	if err != nil {
		log.Fatal("Failed to make ./agc/objects")
	}

	// On Windows, set hidden attr
	if runtime.GOOS == "windows" {
		cmd := exec.Command("attrib", "+H", GitRepo)
		cmd.Run()
	}

	fmt.Println("Git repo initialized at ", fp)
}
