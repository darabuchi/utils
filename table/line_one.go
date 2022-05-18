package table

import (
	"fmt"
	"image"
	"strings"

	"github.com/darabuchi/log"
	"github.com/wcharczuk/go-chart/v2"
)

type LineOne struct {
	text string

	fontSize int32

	size *Size
}

func (p *LineOne) DrawImg(x, y int32, img *image.RGBA) error {
	lines := strings.Split(p.text, "\n")

	for _, line := range lines {
		size, err := DrawFont(img, image.NewUniform(chart.ColorBlack), x, y, line, p.fontSize)
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
		fontSize: 18,
	}
}

func (p *LineOne) SetText(format string, a ...interface{}) *LineOne {
	p.text = strings.ReplaceAll(fmt.Sprintf(format, a...), "\t", "    ")
	return p
}
