package commit

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// readMessage takes the commit message from the commit path
// it returns error when we cannot read the file from the given path
// when the file successfully read, it returns the message in string
func Read(path string) (CommitMessage, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error whilst reading file %s: %w", path, err)
	}

	return CommitMessage(string(bytes)), nil
}

// newFirstLine extracts a commit's header from entire commit messages
// entire commit message means header, body, and footer
// in this program, we want to check only the style of header
func NewFirstLine(m string) string {
	lines := strings.Split(m, "\n")

	if len(lines) == 0 {
		panic(fmt.Errorf("no line break in the commit message"))
	}

	return lines[0]
}