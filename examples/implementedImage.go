package main

/*
Taken from: https://tour.golang.org/methods/25

Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.

Define your own Image type, implement the necessary methods, and call pic.ShowImage.

Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

ColorModel should return color.RGBAModel.

At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
*/
import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Width  int
	Height int
}

func Fill(x, y int) uint8 {
	return uint8(x ^ y)
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.Width, im.Height)
}

func (im Image) At(x, y int) color.Color {
	v := Fill(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{1080, 720}
	pic.ShowImage(m)
}
