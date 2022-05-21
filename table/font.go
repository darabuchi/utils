package table

import (
	_ "embed"
	"image"

	"github.com/darabuchi/log"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed font.ttf
var fontBuf []byte

var _font *truetype.Font

const (
	defaultFontSize = 18
)

func init() {
	var err error
	_font, err = freetype.ParseFont(fontBuf)
	if err != nil {
		log.Panicf("err:%v", err)
		return
	}
}

func SetFont(f *truetype.Font) {
	_font = f
}

func TextSize(label string, fontSize float64) *Size {
	return NewSize(fontSize*2*float64(len([]rune(label))), (fontSize+float64(int64(fontSize)>>6))*2)
}

func DrawFont(dst *image.RGBA, src image.Image, x, y float64, label string, fontSize float64) (*Size, error) {
	c := freetype.NewContext()
	c.SetFont(_font)
	c.SetFontSize(fontSize)
	c.SetClip(dst.Bounds())
	c.SetDst(dst)
	c.SetDPI(144)
	c.SetSrc(src)
	c.SetHinting(font.HintingFull)

	size, err := c.DrawString(label, freetype.Pt(int(x), int(y+fontSize*2)))
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return NewSize(float64(size.X), float64(size.Y)), nil
}
