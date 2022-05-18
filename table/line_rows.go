package table

type Rows struct {
	cells []Cell

	size *Size
}

func (p *Rows) DrawImg() {
	// TODO implement me
	panic("implement me")
}

func (p *Rows) MinSize() *Size {
	var w, h int32
	for _, cell := range p.cells {
		s := cell.MinSize()
		w += s.Width
		if h < s.Height {
			h = s.Height
		}
	}

	return NewSize(w, h)
}

func (p *Rows) setSize(size *Size) {
	p.size = size
}

func (p *Rows) Size() *Size {
	if p.size != nil {
		return p.size
	}
	return p.MinSize()
}

func NewRows() *Rows {
	return &Rows{cells: []Cell{}}
}

func (p *Rows) AddCell(cell Cell) *Rows {
	p.cells = append(p.cells, cell)
	return p
}
