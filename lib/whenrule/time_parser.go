package whenrule

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
)

// 11:00:00 pm
// 11:00AM
// 11:00:00
// 11:00
func Time(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile(`(?:\W|^)` +
			`([0-2]?[0-9])` +
			`:` +
			`([0-5]?[0-9])` +
			`(?::([0-5]?[0-9]))?` +
			`(?:\s?([apAP]\.?[mM]\.?))` +
			`(?:\W|$)`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Hour != nil || c.Minute != nil || c.Second != nil) && s != rules.Override {
				return false, nil
			}

			hour, _ := strconv.Atoi(m.Captures[0])
			minute, _ := strconv.Atoi(m.Captures[1])
			second, _ := strconv.Atoi(m.Captures[2])
			ampm := strings.ToLower(strings.Replace(m.Captures[3], ".", "", -1))

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

func FixAMPM(ampm string, hour int) (int, bool) {
	if ampm != "" && hour > 12 {
		return 0, false
	}

	if ampm == "pm" {
		hour += 12
	}

	return hour, true
}
