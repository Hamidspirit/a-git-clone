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
		objectType := hashObjectCmd.String("type", "blob", "Type of object hash")

		// Parse the flags starting from args[2:]
		hashObjectCmd.Parse(args[2:])
		fp, hash := agc.HashObject(objectPath, objectType, args)
		fmt.Printf("hash of %s:\n %s", fp, hash)

	case "cat-file":
		catFileCmd := flag.NewFlagSet("cat-file", flag.ExitOnError)

		catFileCmd.Parse(args[2:])
		agc.CatFile(args[2])

	default:
		log.Println("valid command Not Found")
	}
}
