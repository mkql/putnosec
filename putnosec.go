package putnosec

import (
	"encoding/json"
	"os"
)

type GoSecOutput struct {
	Issues []GoSecIssue `json:"issues"`
}

type GoSecIssue struct {
	RuleID string      `json:"rule_id"`
	File   string      `json:"file"`
	Line   json.Number `json:"line"`
}

func MustReadGoSecOutputJSON(path string) GoSecOutput {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	d := json.NewDecoder(f)
	var result GoSecOutput
	if err := d.Decode(&result); err != nil {
		panic(err)
	}
	return result
}
