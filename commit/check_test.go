package commit_test

import (
	"fmt"
	"github.com/TangoEnSkai/commit-go/commit"
	"testing"
)

func TestCheckLength(t *testing.T) {
	const (
		minLength         = 10
		maxLength         = 60
		shortCommit       = "short"
		validLengthCommit = "the commit length is good enough"
		longCommit        = "this commit is toooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo long"
	)
	type args struct {
		m commit.Message
	}
	tests := []struct {
		name       string
		args       args
		wantErrStr string
		wantOk     bool
	}{
		{
			name: "success/valid commit message",
			args: args{
				m: validLengthCommit,
			},
			wantErrStr: "",
			wantOk:     true,
		},
		{
			name: "failed/too short",
			args: args{
				m: shortCommit,
			},
			wantErrStr: fmt.Sprintf("\tthe commit `%s` is too short: got(%d), need >= (%d)", shortCommit, len(shortCommit), minLength),
			wantOk:     false,
		},
		{
			name: "failed/too long",
			args: args{
				m: longCommit,
			},
			wantErrStr: fmt.Sprintf("\tthe commit `%s` is too long: got(%d), need <= (%d)", longCommit, len(longCommit), maxLength),
			wantOk:     false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			gotErrStr, gotOk := commit.CheckLength(tc.args.m)
			if gotErrStr != tc.wantErrStr {
				t.Errorf("CheckLength() gotErrStr = %v, want %v", gotErrStr, tc.wantErrStr)
			}
			if gotOk != tc.wantOk {
				t.Errorf("CheckLength() gotOk = %v, want %v", gotOk, tc.wantOk)
			}
		})
	}
}
