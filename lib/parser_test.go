package lib

import (
	"testing"
	"time"
)

func TestParser(t *testing.T) {
	parser := NewParser()

	base := time.Date(1992, time.June, 18, 12, 0, 0, 0, time.Local)
	resultTime, err := parser.Parse("tomorrow 8am", base)
	if err != nil {
		t.Errorf("expected nil but got: %v", err)
	}
	expectedTime := time.Date(1992, time.June, 19, 8, 0, 0, 0, time.Local)
	if resultTime != expectedTime {
		t.Errorf("expected %s but got: %s", expectedTime, resultTime)
	}
}
