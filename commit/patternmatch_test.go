package commit_test

import (
	"fmt"
	"github.com/TangoEnSkai/commit-go/commit"
	"testing"
)

func TestPatternMatch(t *testing.T) {
	const (
		pattern = `^(BREAKING CHANGE|build|chore|ci|docs|feat|fix|perf|refactor|style|test)(\([a-z \-]+\))?: [\w \-]+$`
		validPattern = "perf: optimise pattern matching"
		invalidTypePattern = "performance: is not valid type"
	)

	errorMessage := fmt.Sprintf("invalid commit message. \n\tmust follow this rule: %v\n\t\t", pattern)

	type args struct {
		m commit.Message
	}
	tests := []struct {
		name       string
		args       args
		wantErrMsg string
		wantOk     bool
	}{
		{
			name: "success/pattern match ok",
			args: args{
				m: validPattern,
			},
			wantErrMsg: "",
			wantOk: true,
		},
		{
			name: "failed/pattern match failed by type",
			args: args{
				m: invalidTypePattern,
			},
			wantErrMsg: errorMessage,
			wantOk: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotErrMsg, gotOk := commit.PatternMatch(tc.args.m)
			if gotErrMsg != tc.wantErrMsg {
				t.Errorf("PatternMatch() gotErrMsg = %v, want %v", gotErrMsg, tc.wantErrMsg)
			}
			if gotOk != tc.wantOk {
				t.Errorf("PatternMatch() gotOk = %v, want %v", gotOk, tc.wantOk)
			}
		})
	}
}