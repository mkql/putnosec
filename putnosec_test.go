package putnosec_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mkql/putnosec"
)

func TestGoSecOutput(t *testing.T) {
	gsout := putnosec.MustReadGoSecOutputJSON("./testdata/gosec_output.json")
	expect := putnosec.GoSecOutput{
		Issues: []putnosec.GoSecIssue{
			{
				RuleID: "G204",
				File:   "/path/to/source.go",
				Line:   "35",
			},
			{
				RuleID: "G304",
				File:   "/path/to/source.go",
				Line:   "246",
			},
		},
	}
	if diff := cmp.Diff(expect, gsout); diff != "" {
		t.Errorf("expect %v but got %v, diff:\n%s", expect, gsout, diff)
	}
}
