package lib

import (
	"github.com/jehiah/go-strftime"
	"time"
)

// Formatter formats time into a specific format.
type Formatter interface {
	Format(string, time.Time) string
}

type formatter struct{}

func (f formatter) Format(format string, t time.Time) string {
	return strftime.Format(format, t)
}

// NewFormatter creates a new Formatter.
func NewFormatter() Formatter {
	return formatter{}
}
