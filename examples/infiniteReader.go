package main

/*
Taken from: https://tour.golang.org/methods/22

Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

// To make it available offline run "go get golang.org/x/tour/gotour"
import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
