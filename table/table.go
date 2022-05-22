package table

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"

	"github.com/auyer/steganography"
	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

type Table struct {
	lines []Line

	steganography interface{}
}

func NewTable() *Table {
	return &Table{}
}

func (p *Table) SetSteganography(data interface{}) *Table {
	p.steganography = data
	return p
}

func (p *Table) AddLine(lines ...Line) *Table {
	p.lines = append(p.lines, lines...)
	return p
}

const (
	borderSize = 4
)

var (
	borderColor = drawing.Color{R: 234, G: 229, B: 227, A: 255}
)

func (p *Table) ToImg() (*bytes.Buffer, error) {

	size := NewSize(0, 0)

	colWidthMap := map[int]float64{}
	rowHightMap := map[int]float64{}
	for rowIdx, line := range p.lines {
		s := line.MinSize()
		if s.Width > size.Width {
			size.Width = s.Width
		}

		switch x := line.(type) {
		case CellsLine:
			err := x.RangeCell(func(colIdx int, cell Cell) error {
				size := cell.MinSize()
				if colWidthMap[colIdx] < size.Width {
					colWidthMap[colIdx] = size.Width
				}
				if rowHightMap[rowIdx] < size.Height {
					rowHightMap[rowIdx] = size.Height
				}
				return nil
			})
			if err != nil {
				log.Errorf("err:%v", err)
				return nil, err
			}
		}

		size.Height += s.Height + borderSize
	}

	{
		var w float64
		for _, i := range colWidthMap {
			w += i + borderSize
		}

		if w > size.Width {
			size.Width = w
		}
	}

	{
		var h float64
		for _, i := range rowHightMap {
			h += i + borderSize
		}

		if h > size.Height {
			size.Height = h
		}
	}

	log.Debugf("img size:%v", size)

	img := image.NewRGBA(image.Rect(0, 0, int(size.Width), int(size.Height)))
	draw.Draw(img, img.Bounds(), image.White, img.Bounds().Min, draw.Src)

	for rowIdx, line := range p.lines {
		switch x := line.(type) {
		case LineOne:
			s := line.MinSize()
			s.Width = size.Width
			line.setSize(NewSize(s.Width, s.Height))

		case CellsLine:
			err := x.RangeCell(func(colIdx int, cell Cell) error {
				cell.setSize(NewSize(colWidthMap[colIdx], rowHightMap[rowIdx]))
				return nil
			})
			if err != nil {
				log.Errorf("err:%v", err)
				return nil, err
			}
		}
	}

	var y float64
	for idx, line := range p.lines {
		s := line.Size()
		log.Debugf("line:%v,size:%v", idx, s)

		DrawRectangleColor(img, borderColor, 0, y+s.Height, size.Width, borderSize)

		err := line.DrawImg(0, y, img)
		if err != nil {
			log.Errorf("err:%v", err)
			return nil, err
		}
		y += s.Height + borderSize
	}

	var b bytes.Buffer

	if p.steganography != nil {
		err := steganography.Encode(&b, img, []byte(utils.ToString(p.steganography)))
		if err != nil {
			log.Errorf("err:%v", err)
			return nil, err
		}
	} else {
		err := png.Encode(&b, img)
		if err != nil {
			log.Errorf("err:%v", err)
			return nil, err
		}
	}

	return &b, nil
}
