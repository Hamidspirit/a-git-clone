package cli

import (
	"flag"
	"fmt"
	"log"

	"github.com/Hamidspirit/a-git-clone/agc"
	"github.com/Hamidspirit/a-git-clone/agc/util"
)

func ParseOsArgs(args []string) {
	if args[1] != "init" {
		util.IsRepo()
	}

	switch args[1] {
	case "init":
		initCmd := flag.NewFlagSet("init", flag.ExitOnError)

		initCmd.Parse(args[2:])
		agc.Init()

	case "commit":
		commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)
		msgFlag := commitCmd.String("m", "commit added", "Add Message to Commits")

		commitCmd.Parse(args[2:])
		fmt.Println(*msgFlag)
	case "hash-object":
		hashObjectCmd := flag.NewFlagSet("hash-object", flag.ExitOnError)
		objectPath := hashObjectCmd.String("p", "", "Path to object")

		// Parse the flags starting from args[2:]
		hashObjectCmd.Parse(args[2:])
		fp, hash := util.HashObject(objectPath, args)
		fmt.Printf("hash of %s:\n %x", fp, hash)
	default:
		log.Println("valid command Not Found")
	}
}
