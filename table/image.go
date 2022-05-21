package table

import (
	"image"
	"image/color"
)

func DrawRectangleColor(img *image.RGBA, color color.Color, x, y, wight, high float64) {
	for i := x; i <= x+wight; i++ {
		for j := y; j < y+high; j++ {
			img.Set(int(i), int(j), color)
		}
	}
}
