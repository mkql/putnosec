package putnosec

import "os"

func Execute(write, verbose bool, suffix ...string) error {
	gsecOut := MustReadGoSecOutputJSON(os.Stdin)
	transformed := Transformed(gsecOut.Issues, suffix...)
	transformed.PrintPlan(verbose)
	if write {
		return transformed.Overwrite()
	}
	return nil
}
