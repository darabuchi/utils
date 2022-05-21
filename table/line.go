package table

import (
	"image"
)

type Line interface {
	DrawImg(x, y float64, img *image.RGBA) error
	MinSize() *Size
	setSize(*Size)
	Size() *Size
	Cols() int32
}

type CellsLine interface {
	Line
	RangeCell(func(idx int, cell Cell) error) error
}
