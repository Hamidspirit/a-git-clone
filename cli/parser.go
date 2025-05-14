package cli

import (
	"flag"
	"fmt"
	"log"

	"github.com/Hamidspirit/a-git-clone/agc"
	"github.com/Hamidspirit/a-git-clone/util"
)

func ParseOsArgs(args []string) {
	if args[1] != "init" {
		util.IsRepo(agc.GitRepo)
	}

	switch args[1] {
	case "init":
		initCmd := flag.NewFlagSet("init", flag.ExitOnError)
		pathflag := initCmd.String("path", ".", "initialize repo at path")

		initCmd.Parse(args[2:])

		path := util.PathParser(*pathflag)
		agc.Init(path)

	case "commit":
		commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)
		msgFlag := commitCmd.String("m", "commit added", "Add Message to Commits")

		commitCmd.Parse(args[2:])
		fmt.Println(*msgFlag)
	case "hash-object":
		hashObjectCmd := flag.NewFlagSet("hash-object", flag.ExitOnError)
		objectPath := hashObjectCmd.String("p", ".", "Path to object")
		objectType := hashObjectCmd.String("type", "blob", "Type of object hash")

		// Parse the flags starting from args[2:]
		hashObjectCmd.Parse(args[2:])
		path := util.PathParser(*objectPath)
		file := util.ExtractName(args[2:])
		fp, hash := agc.HashObject(path, *objectType, file)
		fmt.Printf("hash of %s:\n %s", fp, hash)

	case "cat-file":
		catFileCmd := flag.NewFlagSet("cat-file", flag.ExitOnError)

		catFileCmd.Parse(args[2:])
		agc.CatFile(args[2])

	case "write-tree":
		writeTreeCmd := flag.NewFlagSet("write-tree", flag.ExitOnError)

		writeTreeCmd.Parse(args[2:])
		agc.WriteTree()
	default:
		log.Println("valid command Not Found")
	}
}
