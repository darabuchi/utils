package table

import (
	"image"
)

type Cell interface {
	DrawImg(x, y float64, img *image.RGBA) error
	MinSize() *Size
	setSize(*Size)
	Size() *Size
}
