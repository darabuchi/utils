package table

import (
	"image"
	"io/fs"
	"unicode"

	"github.com/AndreKR/multiface"
	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var fontList []*truetype.Font

const (
	defaultFontSize = 18
)

func AddFount(f *truetype.Font) {
	fontList = append(fontList, f)
}

func AddFontWithBuf(fontBuf []byte) error {
	f, err := freetype.ParseFont(fontBuf)
	if err != nil {
		log.Panicf("err:%v", err)
		return err
	}

	AddFount(f)
	return nil
}

func AddFontWithPath(path string) error {
	fontBuf, err := utils.FileRead(path)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return AddFontWithBuf(fontBuf)
}

func AddFontWithPathInFs(path string, fsys fs.FS) error {
	fontBuf, err := utils.FileReadWithFs(path, fsys)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return AddFontWithBuf(fontBuf)
}

func TextSize(label string, fontSize float64) *Size {
	return NewSize(fontSize*2*FontLen(label), (fontSize+float64(int64(fontSize)>>6))*2)
}

func FontLen(str string) float64 {
	var count float64
	for _, v := range str {
		count++
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	return count / 2
	// return float64(len(str) * 2)
}

func DrawFont(dst *image.RGBA, src image.Image, x, y float64, label string, fontSize float64) {

	d := font.Drawer{}
	d.Dst = dst
	d.Src = src

	face := new(multiface.Face)
	for _, f := range fontList {
		face.AddTruetypeFace(truetype.NewFace(f, &truetype.Options{Size: fontSize, DPI: 144}), f)
	}
	d.Face = face

	d.Dot = freetype.Pt(int(x), int(y+fontSize*2))
	d.DrawString(label)
}
