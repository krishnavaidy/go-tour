package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(s []byte) (int, error) {
	n, err := r.r.Read(s)
	s = s[:cap(s)]
	for i := range s {
		s[i] = s[i] + 13
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
