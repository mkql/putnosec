package putnosec

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

type TransformedFile struct {
	Path    string
	Content *bytes.Buffer
}

func (t *TransformedFile) String() string {
	return fmt.Sprintf("path: %s\ncontent:\n%s", t.Path, t.Content.String())
}

type TransformedFiles []TransformedFile

func (ts *TransformedFiles) Overwrite() error {
	if ts == nil {
		return nil
	}
	for _, t := range *ts {
		f, err := os.Create(t.Path)
		if err != nil {
			return errors.Wrapf(err, "failed to open file %q for overwriting", t.Path)
		}
		if _, err := io.Copy(f, t.Content); err != nil {
			return errors.Wrapf(err, "failed to write to file %q", t.Path)
		}
		if err := f.Close(); err != nil {
			return errors.Wrapf(err, "failed to close file %q", t.Path)
		}
	}
	return nil
}

func (ts *TransformedFiles) PrintPlan(verbose bool) {
	if ts == nil {
		return
	}
	for _, t := range *ts {
		log.Printf("will put #nosec directive to the file %s\n", t.Path)
		if verbose {
			log.Println(t.Content.String())
		}
	}
}

func Transformed(issues []GoSecIssue, suffix ...string) TransformedFiles {
	filesWithIssues := make(map[string][]GoSecIssue)
	for _, issue := range issues {
		filesWithIssues[issue.File] = append(filesWithIssues[issue.File], issue)
	}
	// sort issues by line numbers
	for _, iss := range filesWithIssues {
		sort.SliceStable(iss, func(i, j int) bool {
			return iss[i].Target() < iss[j].Target()
		})
	}

	var result TransformedFiles

	for path, iss := range filesWithIssues {
		for _, is := range iss {
			fmt.Printf("%s:%d block: %t\n", is.File, is.Target(), is.Block())
		}
		// #nosec G304
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		s := bufio.NewScanner(f)
		var lines []string
		for s.Scan() {
			lines = append(lines, s.Text())
		}
		if err := s.Err(); err != nil {
			panic(err)
		}
		var idx int
		var transformedLines []string
		for _, issue := range iss {
			target := issue.Target() - 1 // Must l < len(lines)
			for idx < target {           // idx < len(lines) holds
				transformedLines = append(transformedLines, lines[idx])
				idx++
			}
			// here idx == l
			directive := strings.Join(append([]string{"// #nosec", issue.RuleID}, suffix...), " ")
			// directive := fmt.Sprintf("// #nosec %s", issue.RuleID)
			transformedLines = append(transformedLines, directive)
		}
		for idx < len(lines) {
			transformedLines = append(transformedLines, lines[idx])
			idx++
		}
		buf := new(bytes.Buffer)
		for _, l := range transformedLines {
			if _, err := fmt.Fprintln(buf, l); err != nil {
				panic(err)
			}
		}
		formattedBuf := new(bytes.Buffer)
		formatted, err := format.Source(buf.Bytes())
		if err != nil {
			fmt.Println(buf.String())
			panic(err)
		}
		if _, err := formattedBuf.Write(formatted); err != nil {
			panic(err)
		}
		result = append(result, TransformedFile{
			Path:    path,
			Content: formattedBuf,
		})
		if err := f.Close(); err != nil {
			panic(err)
		}
	}
	return result
}
