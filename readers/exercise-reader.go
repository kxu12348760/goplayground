package main

import (
	"golang.org/x/tour/reader"
	"fmt"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	m := MyReader{}
	reader.Validate(m)
	b := make([]byte, 8)

	for i := 0; i < 10; i++ {
		n, err := m.Read(b)
		fmt.Printf("n = %v err =%v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
	}
}
