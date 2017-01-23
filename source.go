package nlftime

import (
	"bytes"
	"io"
)

// Source generates a text.
type Source interface {
	Generate() (string, error)
}

type source struct {
	r   io.Reader
	buf *bytes.Buffer
}

func (s source) Generate() (string, error) {
	_, err := io.CopyN(s.buf, s.r, int64(s.buf.Cap()))
	if err != nil {
		return "", err
	}
	return s.buf.String(), nil
}

// NewSource creates a new Source.
func NewSource(r io.Reader, bufferSize int) Source {
	return source{r, bytes.NewBuffer(make([]byte, bufferSize))}
}
