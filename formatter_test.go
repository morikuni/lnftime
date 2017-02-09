package nlftime

import (
	"testing"
	"time"
)

func TestFormatter(t *testing.T) {
	formatter := NewStrftimeFormatter("%Y-%m-%d")

	base := time.Date(1992, time.June, 18, 12, 0, 0, 0, time.Local)
	s := formatter.Format(base)
	expectedStr := "1992-06-18"
	if s != expectedStr {
		t.Errorf("expected %s but got: %s", expectedStr, s)
	}
}
