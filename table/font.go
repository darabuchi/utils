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

func TextSize(label string, fontSize int32) *Size {
	return NewSize(fontSize*2, fontSize*2*int32(len([]rune(label))))
}

func DrawFont(dst *image.RGBA, src image.Image, x, y int32, label string, fontSize int32) (*Size, error) {
	c := freetype.NewContext()
	c.SetFont(_font)
	c.SetFontSize(float64(fontSize))
	c.SetClip(dst.Bounds())
	c.SetDst(dst)
	c.SetDPI(144)
	c.SetSrc(src)
	c.SetHinting(font.HintingFull)

	size, err := c.DrawString(label, freetype.Pt(int(x), int(y+fontSize)))
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return NewSize(int32(size.X), int32(size.Y)), nil
}
