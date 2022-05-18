package table

import (
	"image"
)

type Line interface {
	DrawImg(x, y int32, img *image.RGBA) error
	MinSize() *Size
	setSize(*Size)
	Size() *Size
	Cols() int32
}
