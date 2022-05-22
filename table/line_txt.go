package table

import (
	"fmt"
	"image"
	"image/color"
	"strings"

	"github.com/darabuchi/log"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

type Alignment int

const (
	AlignLeft Alignment = iota
	AlignCenter
	AlignRight
)

type LineTxt struct {
	text string

	fontSize float64

	size *Size

	fgColor, bgColor color.Color
	alignment        Alignment
}

func (p *LineTxt) DrawImg(x, y float64, img *image.RGBA) error {
	lines := strings.Split(p.text, "\n")

	size := p.Size()
	DrawRectangleColor(img, p.bgColor, x, y, size.Width, size.Height)

	for idx, line := range lines {
		log.Debugf("line:%v,text:%v,x:%.2f,y:%.2f", idx, line, x, y)
		fs := TextSize(line, p.fontSize)

		switch p.alignment {
		case AlignCenter:
			_, err := DrawFont(img, image.NewUniform(p.fgColor), x+(size.Width-fs.Width)/2, y, line, p.fontSize)
			if err != nil {
				log.Errorf("err:%v", err)
				return err
			}
		case AlignRight:
			_, err := DrawFont(img, image.NewUniform(p.fgColor), x+(size.Width-fs.Width), y, line, p.fontSize)
			if err != nil {
				log.Errorf("err:%v", err)
				return err
			}
		default:
			_, err := DrawFont(img, image.NewUniform(p.fgColor), x, y, line, p.fontSize)
			if err != nil {
				log.Errorf("err:%v", err)
				return err
			}
		}

		y += fs.Height
	}

	return nil
}

func (p *LineTxt) MinSize() *Size {
	lines := strings.Split(p.text, "\n")
	size := NewSize(0, 0)

	for _, line := range lines {
		s := TextSize(line, p.fontSize)
		size.Height += s.Height
		if s.Width > size.Width {
			size.Width = s.Width
		}
	}

	return size
}

func (p *LineTxt) setSize(size *Size) {
	p.size = size
}

func (p *LineTxt) Size() *Size {
	if p.size != nil {
		return p.size
	}
	return p.MinSize()
}

func NewLineText() *LineTxt {
	return &LineTxt{
		text:     "",
		fontSize: defaultFontSize,
		size:     nil,
		fgColor:  drawing.ColorBlack,
		bgColor:  drawing.ColorWhite,
	}
}

func (p *LineTxt) SetText(format string, a ...interface{}) *LineTxt {
	p.text = strings.ReplaceAll(fmt.Sprintf(format, a...), "\t", "    ")
	return p
}

func (p *LineTxt) SetFontSize(fontSize float64) *LineTxt {
	p.fontSize = fontSize
	return p
}

func (p *LineTxt) SetFgColor(color color.Color) *LineTxt {
	p.fgColor = color
	return p
}

func (p *LineTxt) SetBgColor(color color.Color) *LineTxt {
	p.bgColor = color
	return p
}

func (p *LineTxt) SetAlignment(alignment Alignment) *LineTxt {
	p.alignment = alignment
	return p
}

func (p *LineTxt) IsFull() bool {
	return true
}
