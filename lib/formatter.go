package lib

import (
	"strconv"
	"time"

	"github.com/jehiah/go-strftime"
)

// Formatter formats time into a specific format.
type Formatter interface {
	Format(string, time.Time) string
}

type strftimeFormatter struct{}

func (f strftimeFormatter) Format(format string, t time.Time) string {
	return strftime.Format(format, t)
}

type unixFormatter struct{}

func (f unixFormatter) Format(_ string, t time.Time) string {
	return strconv.FormatInt(t.Unix(), 10)
}

// NewStrftimeFormatter creates a new Formatter for strftime format.
func NewStrftimeFormatter() Formatter {
	return strftimeFormatter{}
}

// NewUnixFormatter creates a new Formatter for unix time.
// The format arguments are ignored.
func NewUnixFormatter() Formatter {
	return unixFormatter{}
}
