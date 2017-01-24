package lib

import (
	"testing"
	"time"
)

func TestFormatter(t *testing.T) {
	formatter := NewFormatter()

	base := time.Date(1992, time.June, 18, 12, 0, 0, 0, time.Local)
	s := formatter.Format("%Y-%m-%d", base)
	expectedStr := "1992-06-18"
	if s != expectedStr {
		t.Errorf("expected %s but got: %s", expectedStr, s)
	}
}
