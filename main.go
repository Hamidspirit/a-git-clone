package main

import (
	"fmt"
	"os"

	"github.com/Hamidspirit/a-git-clone/cli"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("This is not a Valid command -> ", os.Args)
		os.Exit(1)
	}
	cli.ParseOsArgs(os.Args)
}
