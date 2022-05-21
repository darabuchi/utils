package table

import (
	"fmt"
)

type Size struct {
	Width, Height float64
	int32
}

func NewSize(width, height float64) *Size {
	return &Size{
		Width:  width,
		Height: height,
	}
}

func (p *Size) AddWidth(width float64) *Size {
	p.Width += width
	return p
}

func (p *Size) AddHeight(height float64) *Size {
	p.Height += height
	return p
}

func (p Size) String() string {
	return fmt.Sprintf("width:%.2f,height:%.2f", p.Width, p.Height)
}
