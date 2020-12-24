package main

/*
Taken from: https://tour.golang.org/moretypes/18

Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)
*/

// To make it available offline run "go get golang.org/x/tour/gotour"
import "golang.org/x/tour/pic"

func Fill(x, y int) uint8 {
	return uint8(x ^ y)
}

func Pic(dx, dy int) [][]uint8 {
	// Allocating the slice
	im := make([][]uint8, dy)
	for i := range im {
		im[i] = make([]uint8, dx)
		// Filling the slice
		for j := range im[i] {
			im[i][j] = Fill(i, j)
		}
	}
	return im
}

func main() {
	pic.Show(Pic)
}
