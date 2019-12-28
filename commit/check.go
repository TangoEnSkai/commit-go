package commit

import "fmt"

const (
	maxLength = 60 // maximum length of a commit message
	minLength = 10 // minimum length of a commit message
)

// CheckLength is checking the length of commit messages
// for the commit that has proper length, it will NOOP
// for longer commit messages, it prints the error message and stop the program
func CheckLength(m message) (errStr string, ok bool) {
	// need to check length
	l := len(m)

	// message only commit message length is shorter than the limit
	if l < minLength {
		return fmt.Sprintf("\tthe commit `%s` is too short: got(%d), need >= (%d)", m, l, minLength), false
	}

	// message only commit message length is longer than the limit
	if l > maxLength {
		return fmt.Sprintf(
			"\tthe commit `%s` is too long: got(%d), need <= (%d)", m, l, maxLength,
		), false
	}

	return "", true
}
