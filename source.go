package nlftime

import (
	"bufio"
	"io"
)

// Source generates a text.
type Source interface {
	Generate() (string, error)
}

type source struct {
	sc *bufio.Scanner
}

func (s source) Generate() (string, error) {
	if !s.sc.Scan() {
		if err := s.sc.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}
	return s.sc.Text(), nil
}

// NewSource creates a new Source.
func NewSource(r io.Reader) Source {
	return source{bufio.NewScanner(r)}
}
