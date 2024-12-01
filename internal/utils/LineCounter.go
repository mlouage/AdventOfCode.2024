package utils

import (
	"bytes"
	"io"
)

type LineCounter struct {
	Size int    // Size of the buffer
	Sep  string // End of line character
}

func NewLineCounter(size int, sep string) *LineCounter {
	if size == 0 {
		size = 32 * 1024
	}
	if sep == "" {
		sep = "\n"
	}
	return &LineCounter{Size: size, Sep: sep}
}

func (b *LineCounter) Count(r io.Reader) (int, error) {
	defaultSize := 32 * 1024
	defaultEndLine := "\n"

	if b.Size == 0 {
		b.Size = defaultSize
	}

	if b.Sep == "" {
		b.Sep = defaultEndLine
	}

	buf := make([]byte, b.Size)
	var count int

	for {
		n, err := r.Read(buf)
		count += bytes.Count(buf[:n], []byte(b.Sep))

		if err != nil {
			if err == io.EOF {
				return count, nil
			}
			return count, err
		}

	}
}