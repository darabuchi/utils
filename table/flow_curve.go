package table

import (
	"image"
	"image/color"

	"github.com/crystal-construct/analytics"
	pie "github.com/elliotchance/pie/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

const (
	flowCurveWight = 20
)

type FlowCurve struct {
	size *Size

	bgColor, fgColor color.Color

	data []float64

	baseData float64
	isFull   bool
}

func (p *FlowCurve) DrawImg(x, y float64, img *image.RGBA) error {
	s := p.Size()

	// 长度为0，表示不需要绘制，直接填充背景色
	if len(p.data) == 0 {
		DrawRectangleColor(img, p.bgColor, x, y, s.Width, s.Height)
		return nil
	} else if len(p.data) == 1 {
		// 长度只有1，就直接一条直线
		DrawRectangleColor(img, p.fgColor, x, y, s.Width, s.Height/2)
		DrawRectangleColor(img, p.bgColor, x, y+s.Height/2, s.Width, 3)
		return nil
	}

	DrawRectangleColor(img, p.bgColor, x, y, s.Width, s.Height)

	max := pie.Max(p.data)
	width := s.Width / float64(len(p.data)-1) / 1000

	series := analytics.NewSeries()
	for i, datum := range p.data {
		series.Add(float64(i), datum)
	}

	fit := series.FitPolynomial(len(p.data))

	// for i := float64(0); i <= float64(len(p.data)); i += 0.001 {
	for i := float64(0); i <= float64(len(p.data))-1; i += 0.001 {
		h := ((max - analytics.Extrapolate(fit, i)) / max) * s.Height
		// log.Debugf("draw at (%f,%f)", i, y+h)
		DrawRectangleColor(img, p.fgColor, x, y+h, width, 6)
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

func (p *FlowCurve) MinSize() *Size {
	return NewSize(float64(flowCurveWight*(len(p.data))), 60)
}

func (p *FlowCurve) setSize(size *Size) {
	p.size = size
}

func (p *FlowCurve) Size() *Size {
	if p.size != nil {
		return p.size
	}
	return p.MinSize()
}

func (p *FlowCurve) SetFgColor(color color.Color) *FlowCurve {
	p.fgColor = color
	return p
}

func (p *FlowCurve) SetBgColor(color color.Color) *FlowCurve {
	p.bgColor = color
	return p
}

func (p *FlowCurve) AddData(data ...float64) *FlowCurve {
	p.data = append(p.data, data...)
	return p
}

func (p *FlowCurve) SetFull(i bool) *FlowCurve {
	p.isFull = i
	return p
}

func (p *FlowCurve) IsFull() bool {
	return p.isFull
}

func NewFlowCurve() *FlowCurve {
	return &FlowCurve{
		fgColor:  drawing.Color{R: 206, G: 228, B: 174, A: 255},
		bgColor:  drawing.ColorWhite,
		data:     nil,
		baseData: 0,
	}
}
