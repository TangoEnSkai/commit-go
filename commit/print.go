package commit

import (
	"fmt"
	"os"
)

// Print prints error message with the commit message
func Print(m Message, errMsg string) {
	fmt.Printf("%s\n\ncommit message:\n\n%s", errMsg, m)
}

// PrintAndExit print error messages in terms of commit types, also terminate the program
func PrintAndExit(m Message, errMsg string) {
	Print(m, errMsg)
	os.Exit(1)
}
