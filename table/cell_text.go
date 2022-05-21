package table

import (
	"fmt"
	"image"
	"image/color"
	"strings"

	"github.com/darabuchi/log"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

type Text struct {
	text string

	size     *Size
	fontSize float64

	fgColor, bgColor color.Color
}

func (p *Text) setSize(size *Size) {
	p.size = size
}

func (p *Text) Size() *Size {
	if p.size != nil {
		return p.size
	}
	return p.MinSize()
}

func (p *Text) DrawImg(x, y float64, img *image.RGBA) error {
	lines := strings.Split(p.text, "\n")

	s := p.Size()
	DrawRectangleColor(img, p.bgColor, x, y, s.Width, s.Height)

	for idx, line := range lines {
		s := TextSize(line, p.fontSize)
		log.Debugf("line:%v,text:%v,x:%.2f,y:%.2f,%v", idx, line, x, y, s)
		_, err := DrawFont(img, image.NewUniform(p.fgColor), x, y, line, p.fontSize)
		if err != nil {
			log.Errorf("err:%v", err)
			return err
		}
		y += s.Height + borderSize
	}

	return nil
}

func (p *Text) MinSize() *Size {
	lines := strings.Split(p.text, "\n")
	size := NewSize(0, 0)

	for _, line := range lines {
		s := TextSize(line, p.fontSize)
		size.Height += s.Height + borderSize
		if s.Width > size.Width {
			size.Width = s.Width
		}
	}

	return size
}

func NewText(format string, a ...interface{}) *Text {
	return &Text{
		text:     strings.ReplaceAll(fmt.Sprintf(format, a...), "\t", "    "),
		fontSize: defaultFontSize,
		fgColor:  drawing.ColorBlack,
		bgColor:  drawing.ColorWhite,
	}
}

func (p *Text) SetText(format string, a ...interface{}) *Text {
	p.text = strings.ReplaceAll(fmt.Sprintf(format, a...), "\t", "    ")
	return p
}

func (p *Text) SetFontSize(fontSize float64) *Text {
	p.fontSize = fontSize
	return p
}

func (p *Text) SetFgColor(color color.Color) *Text {
	p.fgColor = color
	return p
}

func (p *Text) SetBgColor(color color.Color) *Text {
	p.bgColor = color
	return p
}
