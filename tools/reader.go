package tools

import (
	"bufio"
	"bytes"
	"os"
)

type LineReader struct {
	*bufio.Scanner
}

func New(input string) *LineReader {
	s := bufio.NewScanner(bytes.NewBufferString(input))
	s.Split(bufio.ScanLines)
	return &LineReader{Scanner: s}
}

func Open(fileName string) (*LineReader, func(), error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	return &LineReader{Scanner: s}, func() { f.Close() }, nil
}

func (r *LineReader) ForEach(f func(string)) {
	for r.Scan() {
		f(r.Text())
	}
}
