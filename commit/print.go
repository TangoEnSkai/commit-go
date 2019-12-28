package commit

import (
	"fmt"
	"os"
)

// Print prints error message with the commit message
func Print(m message, errMsg string) {
	fmt.Printf("%s\n\nCommit message:\n\n%s", errMsg, m)
}

// PrintAndExit print error messages in terms of commit types, also terminate the program
func PrintAndExit(m message, errMsg string) {
	Print(m, errMsg)
	os.Exit(1)
}
