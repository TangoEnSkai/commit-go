package commit

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Read takes the commit message from the commit path
// it returns error when we cannot read the file from the given path
// when the file successfully read, it returns the message in string
func Read(path string) (message, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error whilst reading file %s: %w", path, err)
	}

	return message(string(bytes)), nil
}

// NewFirstLine extracts a commit's header from entire commit messages
// entire commit message means header, body, and footer
// in this program, we want to check only the style of header
func NewFirstLine(m message) message {
	lines := strings.Split(string(m), "\n")

	if len(lines) == 0 {
		panic(fmt.Errorf("no line break in the commit message"))
	}

	ret := message(lines[0])

	return ret
}
