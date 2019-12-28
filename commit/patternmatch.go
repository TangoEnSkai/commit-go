package commit

import (
	"fmt"
	"regexp"
)

// patternMatch checks whether the given msg follows proper style or out
// the pattern is defined under getPattern and
// built in regexp.MatchString check the format of commit message
// if it fails, it let users know which rules they have to follow
func PatternMatch(m CommitMessage) (errMsg string, ok bool) {
	p := getPattern()

	mStr := string(m)
	// style check
	matched, err := regexp.MatchString(p, mStr)
	if err != nil {
		return fmt.Errorf("error whilst checking commit message %v: %w", m, err).Error(), false
	}

	// since this function is final check, print out the commit message check result
	if !matched {
		return fmt.Sprintf("invalid commit message. \n\tmust follow this rule: %v\n\t\t", p), false
	}

	return "", matched
}

// getPattern return a regular expression of conventional commit message
// The regex pattern is highly inspired by convention of Angular and the Conventional Commits
// however, it could be customized by `@ms-home` in the future, depending on the team's situation
// FYI, the types of commits in regex sorted in alphabetical order for better readability
func getPattern() string {
	const pattern = `^(BREAKING CHANGE|build|chore|ci|docs|feat|fix|perf|refactor|style|test)(\([a-z \-]+\))?: [\w \-]+$`

	return pattern
}