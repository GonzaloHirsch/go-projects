package main

/*
Taken from: https://tour.golang.org/methods/23

A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
*/
import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func Rotate(c byte) byte {
	switch {
	case 'A' <= c && c <= 'Z':
		return 'A' + (((c - 'A') + 13) % 26)
	case 'a' <= c && c <= 'z':
		return 'a' + (((c - 'a') + 13) % 26)
	}
	return c
}

func (r13 *rot13Reader) Read(b []byte) (n int, e error) {
	// Reading from the composed reader
	n, e = r13.r.Read(b)

	for i := 0; i < len(b); i++ {
		b[i] = Rotate(b[i])
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
