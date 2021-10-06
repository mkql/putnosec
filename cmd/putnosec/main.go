package main

import (
	"flag"
	"log"
	"os"

	"github.com/mkql/putnosec"
)

var (
	write   bool
	verbose bool
	suffix  string
)

func init() {
	flag.BoolVar(&write, "w", false, "overwrite target source codes")
	flag.BoolVar(&verbose, "v", false, "verbose mode")
	flag.StringVar(&suffix, "s", "", "suffix to #nosec directives. ex) -s FIXME generates comment like '// #nosec G100 FIXME'.")
}

func main() {
	flag.Parse()
	if err := putnosec.Execute(write, verbose, suffix); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
