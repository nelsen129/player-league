package tape

import (
	"io"
	"os"
)

type Tape struct {
	file *os.File
}

func NewTape(file *os.File) *Tape {
	t := new(Tape)
	t.file = file
	return t
}

func (t *Tape) Write(p []byte) (int, error) {
	t.file.Truncate(0)
	t.file.Seek(0, io.SeekStart)
	return t.file.Write(p)
}
