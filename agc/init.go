package agc

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func Init() {

	err := os.Mkdir(GitRepo, 0750)
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

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Init -> Could not get current directory.")
	}

	fmt.Println("Git repo initialized at ", wd, GitRepo)
}
