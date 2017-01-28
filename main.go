package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/morikuni/nlftime/lib"
	"github.com/spf13/pflag"
)

var (
	Version string
)

func main() {
	os.Exit(Run(os.Args, os.Stdin, os.Stdout, os.Stderr))
}

func Run(args []string, in io.Reader, out io.Writer, errW io.Writer) int {
	flag := pflag.NewFlagSet("nlftime", pflag.ContinueOnError)

	format := flag.String("strftime", "%Y-%m-%dT%H:%M:%S%z", "output in strftime format.")
	unix := flag.Bool("unix", false, "output in unix time.")
	humanize := flag.Bool("humanize", false, "output in relative time.")
	help := flag.BoolP("help", "h", false, "print this help.")
	version := flag.Bool("version", false, "print version of nlftime")
	flag.Usage = func() {
		fmt.Fprintln(errW)
		fmt.Fprintln(errW, "Usage: nlftime [flags] <text>")
		fmt.Fprintln(errW)
		fmt.Fprintln(errW, flag.FlagUsages())
	}

	err := flag.Parse(args[1:])
	if err != nil {
		return 1
	}

	if *help {
		flag.Usage()
		return 0
	}

	if *version {
		fmt.Fprintln(out, "nlftime version", Version)
		return 0
	}

	ags := flag.Args()
	if len(ags) != 0 {
		in = strings.NewReader(strings.Join(ags, " "))
	}

	source := lib.NewSource(in, 2048)
	parser := lib.NewParser()
	var formatter lib.Formatter
	switch {
	case *unix:
		formatter = lib.NewUnixFormatter()
	case *humanize:
		formatter = lib.NewHumanizeFOrmatter()
	default:
		formatter = lib.NewStrftimeFormatter(*format)
	}

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

	fmt.Fprintln(out, formatter.Format(t))
	return 0
}
