// Package tape contains an io.Writer that replaces the file on write
package tape

import (
	"io"
	"os"
)

// Tape primarily serves as an io.Writer that completely
// replaces the file on write
type Tape struct {
	file *os.File
}

// NewTape returns a Tape for a file
func NewTape(file *os.File) *Tape {
	t := new(Tape)
	t.file = file
	return t
}

// Write truncates the file and then writes p to the file,
// effectively completely replacing it
func (t *Tape) Write(p []byte) (int, error) {
	t.file.Truncate(0)
	t.file.Seek(0, io.SeekStart)
	return t.file.Write(p)
}
