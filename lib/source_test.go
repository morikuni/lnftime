package lib

import (
	"strings"
	"testing"
)

func TestSource(t *testing.T) {
	source := NewSource(strings.NewReader(
		`aaa
hoge`),
		5)
	s, err := source.Generate()
	if err != nil {
		t.Errorf("expected nil but got: %v", err)
	}
	expectedStr :=
		`aaa
h`
	if s != expectedStr {
		t.Errorf("expected %q but got: %q", expectedStr, s)
	}
}
