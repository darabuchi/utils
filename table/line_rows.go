package table

import (
	"image"

	"github.com/darabuchi/log"
)

type Rows struct {
	cells []Cell

	size *Size
}

func (p *Rows) RangeCell(f func(idx int, cell Cell) error) error {
	for i, cell := range p.cells {
		err := f(i, cell)
		if err != nil {
			log.Errorf("err:%v", err)
			return err
		}
	}

	return nil
}

func (p *Rows) Cols() int32 {
	return int32(len(p.cells))
}

func (p *Rows) DrawImg(x, y float64, img *image.RGBA) error {
	for idx, cell := range p.cells {
		s := cell.Size()
		log.Debugf("cell:%v,x:%.2f,y:%.2f,%v", idx, x, y, s)
		err := cell.DrawImg(x, y, img)
		if err != nil {
			log.Errorf("err:%v", err)
			return err
		}
		x += s.Width
		DrawRectangleColor(img, borderColor, x, y, borderSize, s.Height)
		x += borderSize
	}

	return nil
}

func (p *Rows) MinSize() *Size {
	var w, h float64
	for _, cell := range p.cells {
		s := cell.MinSize()
		w += s.Width + borderSize
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

func (p *Rows) AddCell(cell ...Cell) *Rows {
	p.cells = append(p.cells, cell...)
	return p
}
