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
		agc.Commit(*msgFlag)
	case "hash-object":
		hashObjectCmd := flag.NewFlagSet("hash-object", flag.ExitOnError)
		objectPath := hashObjectCmd.String("p", ".", "Path to object")
		objectType := hashObjectCmd.String("type", "blob", "Type of object hash")

		// Parse the flags starting from args[2:]
		hashObjectCmd.Parse(args[2:])
		path := util.PathParser(*objectPath)
		files := util.ExtractName(args[2:])
		hashobj := agc.HashObject(path, *objectType, files)
		for _, obj := range hashobj {
			fmt.Printf("hash of %s:\n %s\n", obj.FPath, obj.ObjectID)
		}

	case "cat-file":
		catFileCmd := flag.NewFlagSet("cat-file", flag.ExitOnError)
		fid := catFileCmd.String("p", "", "print content of file with this hash")

		catFileCmd.Parse(args[2:])
		agc.CatFile(*fid)

	case "write-tree":
		writeTreeCmd := flag.NewFlagSet("write-tree", flag.ExitOnError)

		writeTreeCmd.Parse(args[2:])
		agc.WriteTree()
	case "read-tree":
		readTreeCmd := flag.NewFlagSet("read-tree", flag.ExitOnError)
		treeHash := readTreeCmd.String("h", "", "Hash of the tree you want to write to dir")

		readTreeCmd.Parse(args[2:])
		agc.ReadTree(*treeHash, false)
	default:
		log.Println("valid command Not Found")
	}
}
