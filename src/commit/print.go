package commit

import (
	"fmt"
	"os"
)

// print prints error message with the commit message
func Print(m CommitMessage, errMsg string) {
	fmt.Printf("%s\n\nCommit message:\n\n%s", errMsg, m)
}

// printAndExit print error messages in terms of commit types, also terminate the program
func PrintAndExit(m CommitMessage, errMsg string) {
	Print(m, errMsg)
	os.Exit(1)
}
