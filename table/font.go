package table

import (
	"image"
	"io/fs"

	"github.com/AndreKR/multiface"
	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	emoji "github.com/go-xman/go.emoji"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	eastasianwidth "github.com/moznion/go-unicode-east-asian-width"
	"golang.org/x/image/font"
)

var FontList []*truetype.Font

const (
	defaultFontSize = 18
)

func AddFount(f *truetype.Font) {
	FontList = append(FontList, f)
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
	log.Infof("add font %s", path)
	fontBuf, err := utils.FileReadWithFs(path, fsys)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return AddFontWithBuf(fontBuf)
}

func TextSize(label string, fontSize float64) *Size {
	return NewSize(fontSize*FontLen(label), (fontSize+float64(int64(fontSize)>>6))*2)
}

func FontLen(str string) float64 {
	var count float64

	str = emoji.ReplaceEmoji(str, func(emoji string) string {
		count += 4
		return ""
	})

	for _, v := range []rune(str) {
		count++
		if eastasianwidth.IsFullwidth(v) {
			count++
		}
	}
	return count
}

func DrawFont(dst *image.RGBA, src image.Image, x, y float64, label string, fontSize float64) {

	d := font.Drawer{}
	d.Dst = dst
	d.Src = src

	face := new(multiface.Face)
	for _, f := range FontList {
		face.AddTruetypeFace(truetype.NewFace(f, &truetype.Options{Size: fontSize, DPI: 144}), f)
	}
	d.Face = face

	d.Dot = freetype.Pt(int(x), int(y+fontSize*2))
	d.DrawString(label)
}
