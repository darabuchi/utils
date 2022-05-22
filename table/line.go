package table

import (
	"image"
)

type Line interface {
	DrawImg(x, y float64, img *image.RGBA) error
	MinSize() *Size
	setSize(*Size)
	Size() *Size
}

type CellsLine interface {
	Line
	RangeCell(func(idx int, cell Cell) error) error
	Cols() int32
}

type LineOne interface {
	Line

	IsFull() bool
}
