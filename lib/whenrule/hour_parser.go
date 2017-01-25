package whenrule

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
)

// 11 pm
// 11AM
func Hour(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile(`(?:\W|^)` +
			`([0-2]?[0-9])` +
			`(?:\s?([apAP][mM]))` +
			`(?:\W|$)`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Hour != nil || c.Minute != nil || c.Second != nil) && s != rules.Override {
				return false, nil
			}

			hour, _ := strconv.Atoi(m.Captures[0])
			minute := 0
			second := 0
			ampm := strings.ToLower(m.Captures[1])

			hour, ok := FixAMPM(ampm, hour)
			if !ok {
				return false, nil
			}

			c.Hour = &hour
			c.Minute = &minute
			c.Second = &second

			return true, nil
		},
	}
}
