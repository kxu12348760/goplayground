package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// Reader method on readers interface
	// func (T) Read(b []byte) (n int, err error)
	r := strings.NewReader("Hello Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err =%v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}