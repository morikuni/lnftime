package whenrule

import (
	"regexp"
	"strconv"
	"time"

	"github.com/olebedev/when/rules"
)

// YYYY-MM-DD
// YYYY/MM/DD
func YMD(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile(`(?:\W|^)` +
			`([12][0-9]{3})` +
			`[\/-]` +
			`([01]?[0-9])` +
			`[\/-]` +
			`([0-3]?[0-9])` +
			`(?:\W|$)`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Day != nil || c.Month != nil || c.Year != nil) && s != rules.Override {
				return false, nil
			}

			year, _ := strconv.Atoi(m.Captures[0])
			month, _ := strconv.Atoi(m.Captures[1])
			day, _ := strconv.Atoi(m.Captures[2])

			if day == 0 {
				return false, nil
			}

			if month > 12 {
				return false, nil
			}

			if day > 31 {
				return false, nil
			}

			c.Day = &day
			c.Month = &month
			c.Year = &year

			return true, nil
		},
	}
}
