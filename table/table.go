package table

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"

	"github.com/darabuchi/log"
)

type Table struct {
	lines []Line
}

func NewTable() *Table {
	return &Table{}
}

func (p *Table) AddLine(lines ...Line) *Table {
	p.lines = append(p.lines, lines...)
	return p
}

const borderSize = 3

func (p *Table) ToImg() (*bytes.Buffer, error) {

	size := NewSize(0, 0)

	for _, line := range p.lines {
		s := line.MinSize()
		if s.Width > size.Width {
			size.Width = s.Width
		}

		size.Height += s.Height + borderSize
	}

	img := image.NewRGBA(image.Rect(0, 0, int(size.Width), int(size.Height)))
	draw.Draw(img, img.Bounds(), image.White, img.Bounds().Min, draw.Src)

	var y int32
	for _, line := range p.lines {
		s := line.Size()
		err := line.DrawImg(0, y, img)
		if err != nil {
			log.Errorf("err:%v", err)
			return nil, err
		}
		y += s.Height
	}

	var b bytes.Buffer
	err := png.Encode(&b, img)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &b, nil
}
