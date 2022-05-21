package table

import (
	"fmt"
	"image"
	"image/color"
	"strings"

	"github.com/darabuchi/log"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

type LineOne struct {
	text string

	fontSize float64

	size *Size

	fgColor, bgColor color.Color
}

func (p *LineOne) DrawImg(x, y float64, img *image.RGBA) error {
	lines := strings.Split(p.text, "\n")

	size := p.Size()

	for idx, line := range lines {
		log.Debugf("line:%v,text:%v,x:%.2f,y:%.2f", idx, line, x, y)
		fs := TextSize(line, p.fontSize)
		s := p.Size()

		DrawRectangleColor(img, p.bgColor, x, y, s.Width, s.Height)

		size, err := DrawFont(img, image.NewUniform(p.fgColor), x+(size.Width-fs.Width)/2, y, line, p.fontSize)
		if err != nil {
			log.Errorf("err:%v", err)
			return err
		}
		x += size.Height
	}

	return nil
}

func (p *LineOne) Cols() int32 {
	return 1
}

func (p *LineOne) MinSize() *Size {
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

func (p *LineOne) setSize(size *Size) {
	p.size = size
}

func (p *LineOne) Size() *Size {
	if p.size != nil {
		return p.size
	}
	return p.MinSize()
}

func NewLine() *LineOne {
	return &LineOne{
		text:     "",
		fontSize: defaultFontSize,
		size:     nil,
		fgColor:  drawing.ColorBlack,
		bgColor:  drawing.ColorWhite,
	}
}

func (p *LineOne) SetText(format string, a ...interface{}) *LineOne {
	p.text = strings.ReplaceAll(fmt.Sprintf(format, a...), "\t", "    ")
	return p
}

func (p *LineOne) SetFontSize(fontSize float64) *LineOne {
	p.fontSize = fontSize
	return p
}

func (p *LineOne) SetFgColor(color color.Color) *LineOne {
	p.fgColor = color
	return p
}

func (p *LineOne) SetBgColor(color color.Color) *LineOne {
	p.bgColor = color
	return p
}
