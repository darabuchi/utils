package table_test

import (
	"bytes"
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

func TestTable(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	tb := table.NewTable().
			AddLine(
				table.NewLineText().
					SetText("这行asdfg").
					SetFontSize(35).
					SetAlignment(1).
					SetFgColor(drawing.ColorGreen).
					SetBgColor(drawing.ColorBlue),
				table.NewRows().
						AddCell(table.NewText("流量情况").
							SetFgColor(drawing.ColorGreen).
							SetBgColor(drawing.ColorBlue)).
					AddCell(table.NewText("流量消耗（上传/下载）")),
				table.NewRows().
						AddCell(table.NewFlowRate().
							AddData(10, 20, 10, 20, 10)).
					AddCell(table.NewText("↑100Mbps\n↓100Mbps")),
			)
	
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
	
	utils.FileWrite("test.png", b.String())
	
	// 	// Define a small dataset
	// 	x := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	// 	y := []float64{10, 20, 10, 20, 10, 20, 10, 20}
	//
	// 	// Use the dataset as a series
	// 	series1 := analytics.NewSeriesFrom(x, y)
	//
	// 	// Get a 3rd Order Polynomial fit for the series
	// 	// fit := series1.FitPolynomial(2)
	//
	// 	// // Display two interpolated/extrapolated points
	// 	// fmt.Println(analytics.Extrapolate(fit, 4))
	// 	// fmt.Println(analytics.Extrapolate(fit, 9))
	//
	// 	// // Create smoothed version of the dataset
	// 	// series2 := series1.Smoother(3)
	// 	//
	// 	// // Display the underlying smoothed values
	// 	// x1, y1 := series2.ToArrays()
	// 	// fmt.Println(x1)
	// 	// fmt.Println(y1)
}
