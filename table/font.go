package table

import (
	"embed"
	_ "embed"
	"image"
	"unicode"
	
	"github.com/AndreKR/multiface"
	"github.com/darabuchi/log"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed font
var fontFs embed.FS

var fontList []*truetype.Font

const (
	defaultFontSize = 18
)

func init() {
	add := func(path string) {
		log.Infof("load font %s", path)
		fontBuf, err := fontFs.ReadFile("font/" + path)
		if err != nil {
			log.Panicf("err:%v", err)
			return
		}
		f, err := freetype.ParseFont(fontBuf)
		if err != nil {
			log.Panicf("err:%v", err)
			return
		}
		fontList = append(fontList, f)
	}
	
	dirs, err := fontFs.ReadDir("font")
	if err != nil {
		log.Panicf("err:%v", err)
		return
	}
	
	for _, dir := range dirs {
		if dir.IsDir() {
			continue
		}
		add(dir.Name())
	}
}

// func SetFont(f *truetype.Font) {
// 	_font = f
// }

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
	face := new(multiface.Face)
	for _, f := range fontList {
		face.AddTruetypeFace(truetype.NewFace(f, &truetype.Options{Size: fontSize, DPI: 144}), f)
	}
	
	d := font.Drawer{}
	d.Dst = dst
	d.Src = src
	d.Face = face
	d.Dot = freetype.Pt(int(x), int(y+fontSize*2))
	d.DrawString(label)
}
