package table

type Size struct {
	Width  int32
	Height int32
}

func NewSize(width, height int32) *Size {
	return &Size{
		Width:  width,
		Height: height,
	}
}

func (p *Size) AddWidth(width int32) *Size {
	p.Width += width
	return p
}

func (p *Size) AddHeight(height int32) *Size {
	p.Height += height
	return p
}
