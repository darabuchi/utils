package table

import (
	"image"
	"image/color"

	"github.com/elliotchance/pie/pie"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

const (
	flowHistogramWight = 20
)

type FlowHistogram struct {
	size *Size

	bgColor, fgColor color.Color

	data pie.Float64s

	baseData float64
	isFull   bool
}

func (p *FlowHistogram) DrawImg(x, y float64, img *image.RGBA) error {
	s := p.Size()

	// 长度为0，表示不需要绘制，直接填充背景色
	if len(p.data) == 0 {
		DrawRectangleColor(img, p.bgColor, x, y, s.Width, s.Height)
		return nil
	} else if len(p.data) == 1 {
		// 长度只有1，就直接一条直线
		DrawRectangleColor(img, p.bgColor, x, y, s.Width, s.Height/2)
		DrawRectangleColor(img, p.fgColor, x, y+s.Height/2, s.Width, s.Height/2)
		return nil
	}

	DrawRectangleColor(img, p.fgColor, x, y, s.Width, s.Height)

	max := p.data.Max()
	width := s.Width / float64(len(p.data))

	for i := 0; i < len(p.data); i++ {
		h := ((max - p.data[i]) / max) * s.Height
		DrawRectangleColor(img, p.bgColor, x+1, y, width, h)
		x += width
	}

	// for idx, datum := range p.data {
	// 	h := ((max - datum) / max) * s.Height
	// 	log.Debugf("flow rate %d(%.2f): %.2f", idx, datum, h)
	// 	DrawRectangleColor(img, p.bgColor, x, y, width, h)
	// 	x += width
	// }

	return nil
}

func (p *FlowHistogram) MinSize() *Size {
	return NewSize(float64(flowHistogramWight*(len(p.data))), 60)
}

func (p *FlowHistogram) setSize(size *Size) {
	p.size = size
}

func (p *FlowHistogram) Size() *Size {
	if p.size != nil {
		return p.size
	}
	return p.MinSize()
}

func (p *FlowHistogram) SetFgColor(color color.Color) *FlowHistogram {
	p.fgColor = color
	return p
}

func (p *FlowHistogram) SetBgColor(color color.Color) *FlowHistogram {
	p.bgColor = color
	return p
}

func (p *FlowHistogram) AddData(data ...float64) *FlowHistogram {
	p.data = append(p.data, data...)
	return p
}

func (p *FlowHistogram) SetFull(i bool) *FlowHistogram {
	p.isFull = i
	return p
}

func (p *FlowHistogram) IsFull() bool {
	return p.isFull
}

func NewFlowHistogram() *FlowHistogram {
	return &FlowHistogram{
		fgColor:  drawing.Color{R: 206, G: 228, B: 174, A: 255},
		bgColor:  drawing.ColorWhite,
		data:     nil,
		baseData: 0,
	}
}
