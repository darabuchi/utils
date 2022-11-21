package table_test

import (
	"bytes"
	"embed"
	_ "embed"
	"image"
	"image/draw"
	"image/png"
	"testing"

	"github.com/AndreKR/multiface"
	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	"github.com/darabuchi/utils/table"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2/drawing"
	"golang.org/x/image/font"
)

//go:embed chinese.ttf
//go:embed emoji.ttf
//go:embed symbola.ttf
var fontBuf embed.FS

func TestTable(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	err := table.AddFontWithPathInFs("chinese.ttf", fontBuf)
	if err != nil {
		log.Panicf("err:%v", err)
		return
	}
	err = table.AddFontWithPathInFs("symbola.ttf", fontBuf)
	if err != nil {
		log.Panicf("err:%v", err)
		return
	}
	err = table.AddFontWithPathInFs("emoji.ttf", fontBuf)
	if err != nil {
		log.Panicf("err:%v", err)
		return
	}

	tb := table.NewTable().
		AddLine(
			table.NewLineText().
				SetText("💰 🇭🇰 Hong Kong 16").
				SetFontSize(35).
				SetAlignment(table.AlignCenter).
				SetFgColor(drawing.ColorGreen).
				SetBgColor(drawing.ColorBlue),
			table.NewFlowCurve().
				// table.NewFlowRate().
				// table.NewFlowHistogram().
				AddData(50, 200, 150).
				SetFull(true).
				SetFgColor(table.Purple),
		).
		SetWmk("水印")

	b, err := tb.ToImg()
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	err = utils.FileWrite("test.png", b.String())
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	b, err = tb.ToWebp()
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	err = utils.FileWrite("test.webp", b.String())
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	b, err = tb.ToJpg()
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	err = utils.FileWrite("test.jpg", b.String())
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	b, err = tb.ToTiff()
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	err = utils.FileWrite("test.tiff", b.String())
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	b, err = tb.ToBmp()
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	err = utils.FileWrite("test.bmp", b.String())
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	log.Info("finish")
}

func TestCurve(t *testing.T) {
	var err error

	fontBuf1, err := utils.FileRead("chinese.ttf")
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}
	font1, err := freetype.ParseFont(fontBuf1)
	if err != nil {
		log.Panicf("err:%v", err)
		return
	}

	fontBuf2, err := utils.FileRead("NotoEmoji-VariableFont_wght.ttf")
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}
	font2, err := freetype.ParseFont(fontBuf2)
	if err != nil {
		log.Panicf("err:%v", err)
		return
	}

	face := new(multiface.Face)
	opts := &truetype.Options{Size: 12.5, DPI: 72}
	face.AddTruetypeFace(truetype.NewFace(font1, opts), font1)
	face.AddTruetypeFace(truetype.NewFace(font2, opts), font2)

	img := image.NewRGBA(image.Rect(0, 0, 50, 50))
	draw.Draw(img, img.Rect, image.White, image.ZP, draw.Src)

	d := font.Drawer{}
	d.Dst = img
	d.Src = image.Black
	d.Face = face
	d.Dot = freetype.Pt(10, 25)
	d.DrawString("中文⚓")

	var b bytes.Buffer
	err = png.Encode(&b, img)
	if err != nil {
		log.Panicf("err:%v", err)
		return
	}
}

func TestFontSize(t *testing.T) {
	// const s = "🇭🇰"
	//
	// t.Log(emoji.FilterEmoji(s))
	//
	// return
	log.SetLevel(log.DebugLevel)

	err := table.AddFontWithPathInFs("chinese.ttf", fontBuf)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	err = table.AddFontWithPathInFs("symbola.ttf", fontBuf)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	err = table.AddFontWithPathInFs("NotoEmoji.ttf", fontBuf)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}

	t.Log(table.FontLen("🇭🇰 001"))
	// t.Log(emoji.HasEmoji("🇭🇰"))
}
