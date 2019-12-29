package committer_test

import (
	"github.com/TangoEnSkai/committer-go/committer"
	"testing"
)

func TestCheckCommitType(t *testing.T) {
	const(
		validCommit = "test: check commit type"
	)
	type args struct {
		m committer.Message
	}

	tests := []struct {
		name       string
		args       args
		wantErrMsg string
		wantOk     bool
	}{
		{
			name: "success/valid commit",
			args: args{
				m: validCommit,
			},
			wantErrMsg: "",
			wantOk: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			gotErrMsg, gotOk := committer.CheckCommitType(tc.args.m)
			if gotErrMsg != tc.wantErrMsg {
				t.Errorf("CheckCommitType() gotErrMsg = %v, want %v", gotErrMsg, tc.wantErrMsg)
			}
			if gotOk != tc.wantOk {
				t.Errorf("CheckCommitType() gotOk = %v, want %v", gotOk, tc.wantOk)
			}
		})
	}
}

