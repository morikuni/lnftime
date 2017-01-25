package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/morikuni/nlftime/lib"
)

func main() {
	format := kingpin.Flag("format", "strftime format string").Short('f').Default("%Y-%m-%dT%H:%M:%S%z").String()

	kingpin.Parse()

	source := lib.NewSource(os.Stdin, 2048)
	parser := lib.NewParser()
	formatter := lib.NewFormatter()

	s, err := source.Generate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read a text from the source: %s", err)
		os.Exit(1)
	}

	t, err := parser.Parse(s, time.Now())
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse the text: %s", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, formatter.Format(*format, t))
}
