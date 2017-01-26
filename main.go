package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/morikuni/nlftime/lib"
)

func main() {
	os.Exit(Run(os.Args, os.Stdin, os.Stdout, os.Stderr))
}

func Run(args []string, in io.Reader, out io.Writer, errW io.Writer) int {
	format := kingpin.Flag("format", "strftime format string").Short('f').Default("%Y-%m-%dT%H:%M:%S%z").String()

	kingpin.Parse()

	source := lib.NewSource(in, 2048)
	parser := lib.NewParser()
	formatter := lib.NewFormatter()

	s, err := source.Generate()
	if err != nil {
		fmt.Fprintf(errW, "failed to read a text from the source: %s\n", err)
		return 1
	}

	t, err := parser.Parse(s, time.Now())
	if err != nil {
		fmt.Fprintf(errW, "failed to parse the text: %s\n", err)
		return 1
	}

	fmt.Fprintln(out, formatter.Format(*format, t))
	return 0
}
