package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) At(x, y int) color.Color {
	c := uint8((x + y) / 2)
	return color.RGBA{c, c, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
