package nlftime

import "io"

// Source generates a text.
type Source interface {
	Generate() (string, error)
}

type source struct {
	r   io.Reader
	buf []byte
}

func (s source) Generate() (string, error) {
	_, err := s.r.Read(s.buf)
	if err != nil {
		return "", err
	}
	return string(s.buf), nil
}

// NewSource creates a new Source.
func NewSource(r io.Reader, bufferSize int) Source {
	return source{r, make([]byte, bufferSize)}
}
