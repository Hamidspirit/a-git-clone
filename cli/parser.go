package cli

import (
	"flag"
	"log"

	"github.com/Hamidspirit/a-git-clone/agc"
)

func ParseOsArgs(args []string) {
	switch args[1] {
	case "init":
		initCmd := flag.NewFlagSet("init", flag.ExitOnError)

		initCmd.Parse(args[2:])
		agc.Init()
	default:
		log.Println("valid command Not Found")
	}
}
