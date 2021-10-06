package putnosec

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type GoSecOutput struct {
	Issues []GoSecIssue `json:"issues"`
}

type GoSecIssue struct {
	RuleID string `json:"rule_id"`
	File   string `json:"file"`
	Line   string `json:"line"`
}

func (i *GoSecIssue) Target() int {
	hyphen := strings.Index(i.Line, "-")
	var num string
	if hyphen == -1 {
		num = i.Line
	} else {
		num = i.Line[:hyphen]
	}
	result, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	return result
}

func (i *GoSecIssue) Block() bool {
	hyphen := strings.Index(i.Line, "-")
	return hyphen != -1
}

func MustReadGoSecOutputJSON(r io.Reader) GoSecOutput {
	d := json.NewDecoder(r)
	var result GoSecOutput
	if err := d.Decode(&result); err != nil {
		panic(err)
	}
	return result
}
