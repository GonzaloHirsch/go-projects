package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// The reader interface represents the read end of a stream of data
	r := strings.NewReader("Hello, Reader!")

	// Buffer of 8 bytes
	b := make([]byte, 8)
	for {
		// Populates the byte slice
		n, err := r.Read(b)
		// Contents of the slice
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		// Message of the slice
		fmt.Printf("b[:n] = %q\n", b[:n])
		// If the error is EOF, cut the execution
		if err == io.EOF {
			break
		}
	}
}
