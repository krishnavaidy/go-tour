package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	Min := image.Point{0, 0}
	Max := image.Point{100, 100}
	return image.Rectangle{Min, Max}
}

func (i Image) At(x, y int) color.Color {
	index := uint8(x*x + y*y)
	return color.Gray{index}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
