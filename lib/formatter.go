package lib

import (
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/jehiah/go-strftime"
)

// Formatter formats time into a specific format.
type Formatter interface {
	Format(time.Time) string
}

type strftimeFormatter struct {
	format string
}

func (f strftimeFormatter) Format(t time.Time) string {
	return strftime.Format(f.format, t)
}

type unixFormatter struct{}

func (f unixFormatter) Format(t time.Time) string {
	return strconv.FormatInt(t.Unix(), 10)
}

type humanizeFormatter struct{}

func (f humanizeFormatter) Format(t time.Time) string {
	return humanize.Time(t)
}

// NewStrftimeFormatter creates a new Formatter for strftime format.
func NewStrftimeFormatter(format string) Formatter {
	return strftimeFormatter{format}
}

// NewUnixFormatter creates a new Formatter for unix time.
func NewUnixFormatter() Formatter {
	return unixFormatter{}
}

func NewHumanizeFOrmatter() Formatter {
	return humanizeFormatter{}
}
