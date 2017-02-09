package nlftime

import (
	"strings"
	"testing"
)

func TestSource(t *testing.T) {
	source := NewSource(strings.NewReader(
		`aaa
hoge`))

	s, err := source.Generate()
	if err != nil {
		t.Errorf("expected nil but got: %v", err)
	}
	expectedStr := "aaa"
	if s != expectedStr {
		t.Errorf("expected %q but got: %q", expectedStr, s)
	}
}
