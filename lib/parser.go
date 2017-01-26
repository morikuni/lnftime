package lib

import (
	"errors"
	"time"

	"github.com/morikuni/nlftime/lib/whenrule"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

var (
	// ErrNoTimeFound indicates that no time found in the text.
	ErrNoTimeFound = errors.New("no time found in the text")
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
	p.Add(en.Weekday(rules.Override))
	p.Add(en.CasualDate(rules.Override))
	p.Add(en.CasualTime(rules.Override))
	p.Add(en.Deadline(rules.Override))
	p.Add(whenrule.YMD(rules.Override))
	p.Add(whenrule.Time(rules.Override))
	p.Add(whenrule.Hour(rules.Override))
	return parser{p}
}
