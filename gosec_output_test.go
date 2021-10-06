package putnosec_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mkql/putnosec"
)

func TestGoSecOutput(t *testing.T) {
	path := "./testdata/gosec_output.json"
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	gsout := putnosec.MustReadGoSecOutputJSON(f)
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
				Line:   "246-250",
			},
		},
	}
	if diff := cmp.Diff(expect, gsout); diff != "" {
		t.Errorf("expect %v but got %v, diff:\n%s", expect, gsout, diff)
	}

	if gsout.Issues[0].Block() {
		t.Error()
	}
	if !gsout.Issues[1].Block() {
		t.Error()
	}
	if gsout.Issues[0].Target() != 35 {
		t.Error()
	}
	if gsout.Issues[1].Target() != 246 {
		t.Error()
	}
}
