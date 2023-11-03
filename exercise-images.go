package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Width, Height int
}

// 画像の幅と高さを定義するimage.Rectangleを返す
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

// 色モデルを返す
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// 指定されたピクセルの色を返す
func (img Image) At(x, y int) color.Color {
	v := uint8((x ^ y*x) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{Width: 100, Height: 100}
	pic.ShowImage(m)
}
