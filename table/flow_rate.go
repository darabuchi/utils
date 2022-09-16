package table

import (
	"image"
	"image/color"

	"github.com/crystal-construct/analytics"
	"github.com/elliotchance/pie/pie"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

const (
	flowRateWight = 15
)

type FlowRate struct {
	size *Size

	bgColor, fgColor color.Color

	data pie.Float64s

	baseData float64
	isFull   bool
}

func (p *FlowRate) DrawImg(x, y float64, img *image.RGBA) error {
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
	width := s.Width / float64(len(p.data)-1) / 1000

	series := analytics.NewSeries()
	for i, datum := range p.data {
		series.Add(float64(i), datum)
	}
	fit := series.FitPolynomial(len(p.data))
	for i := float64(0); i <= float64(len(p.data)-1); i += 0.001 {
		h := ((max - analytics.Extrapolate(fit, i)) / max) * s.Height
		DrawRectangleColor(img, p.bgColor, x, y, width, h)
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

func (p *FlowRate) MinSize() *Size {
	return NewSize(float64(flowRateWight*(len(p.data))), 60)
}

func (p *FlowRate) setSize(size *Size) {
	p.size = size
}

func (p *FlowRate) Size() *Size {
	if p.size != nil {
		return p.size
	}
	return p.MinSize()
}

func (p *FlowRate) SetFgColor(color color.Color) *FlowRate {
	p.fgColor = color
	return p
}

func (p *FlowRate) SetBgColor(color color.Color) *FlowRate {
	p.bgColor = color
	return p
}

func (p *FlowRate) AddData(data ...float64) *FlowRate {
	p.data = append(p.data, data...)
	return p
}

func (p *FlowRate) SetFull(i bool) *FlowRate {
	p.isFull = i
	return p
}

func (p *FlowRate) IsFull() bool {
	return p.isFull
}

func NewFlowRate() *FlowRate {
	return &FlowRate{
		fgColor:  drawing.Color{R: 206, G: 228, B: 174, A: 255},
		bgColor:  drawing.ColorWhite,
		data:     nil,
		baseData: 0,
	}
}
