package main

import (
	"fmt"
	"image"
)

func main() {
	// Images come defined in the "image" package
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
