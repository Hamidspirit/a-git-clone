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
		os.Exit(1)
	}

	// On Windows, set hidden attr
	if runtime.GOOS == "windows" {
		cmd := exec.Command("attrib", "+H", GitRepo)
		cmd.Run()
	}
	fmt.Println("Git repo initialized at ", GitRepo)
}
