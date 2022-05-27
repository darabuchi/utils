package table_test

import (
	_ "embed"
	"testing"
	
	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	"github.com/darabuchi/utils/table"
	"github.com/wcharczuk/go-chart/v2/drawing"
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
					AddCell(table.NewText("↑100Mbps\n↓100Mbps").SetAlignment(table.AlignCenter)),
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

// func TestCurve(t *testing.T) {
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
// }
