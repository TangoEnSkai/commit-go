package info

import (
	"fmt"
	"os"

	"github.com/TangoEnSkai/commit-go/commit"
)

// print prints error message with the commit message
func Print(commitMsg, errMsg string) {
	fmt.Printf("%s\n\nCommit message:\n\n%s", errMsg, commitMsg)
}

// printAndExit print error messages in terms of commit types, also terminate the program
func PrintAndExit(m commit.CommitMessage, errMsg string) {
	Print(commitMsg, errMsg)
	os.Exit(1)
}