package lib

import (
	"errors"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"time"
)

var (
	// ErrNoTimeFound indicates that no time found in the context.
	ErrNoTimeFound = errors.New("no time found in the context")
)

// Parser parse string into time object according to specific base time.
type Parser interface {
	Parse(text string, base time.Time) (time.Time, error)
}

type parser struct {
	p *when.Parser
}

func (p parser) Parse(text string, base time.Time) (time.Time, error) {
	r, err := p.p.Parse(text, base)
	if err != nil {
		return time.Time{}, err
	}
	if r == nil {
		return time.Time{}, ErrNoTimeFound
	}
	return r.Time, nil
}

// NewParser creates a new Parser.
func NewParser() Parser {
	p := when.New(nil)
	p.Add(common.All...)
	p.Add(en.All...)
	return parser{p}
}
