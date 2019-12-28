package main

import (
	"fmt"
	"os"

	"github.com/TangoEnSkai/commit-go/commit"
)

// in this main program, we check incoming commit message for:
// 		1. length: commit message must be shorter than maxLength we defined
// 		2. type of commit: we predefined types of commit, all commit must be any of these types
// 		3. style of commit: we follow conventional commits and it's defined as regex
// if any of these checks above fails, it will abort the commit command,
// thus, you should commit again in proper way
func main() {
	origin := initialise()

	commitMsg, err := commit.Read(*origin)
	if err != nil {
		commit.Print(commitMsg, err.Error())
		os.Exit(1)
	}

	header := commit.NewFirstLine(commitMsg)

	// 1. check the commit message is shorter than maximum length
	errMsg, ok := commit.CheckLength(header)
	if !ok {
		commit.PrintAndExit(commitMsg, errMsg)
	}

	// 2. check commit type
	errMsg, ok = commit.CheckCommitType(header)
	if !ok {
		commit.PrintAndExit(commitMsg, errMsg)
	}

	// 3. check style
	errMsg, ok = commit.PatternMatch(header)
	if !ok {
		commit.PrintAndExit(commitMsg, errMsg)
	}
}

func initialise() *string {
	if len(os.Args) < 2 {
		fmt.Println("unexpected error: no arguments passed to this script")
		os.Exit(1)
	}

	pathToCommitMsg := os.Args[1]

	return &pathToCommitMsg
}



