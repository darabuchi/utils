package table

import (
	"bytes"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"

	"github.com/auyer/steganography"
	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	"github.com/nickalie/go-webpbin"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

type Table struct {
	lines []Line

	steganography interface{}

	wmk string
}

func NewTable() *Table {
	return &Table{}
}

// 添加图片隐写
func (p *Table) SetSteganography(data interface{}) *Table {
	p.steganography = data
	return p
}

// 添加水印
func (p *Table) SetWmk(data string) *Table {
	p.wmk = data
	return p
}

func (p *Table) AddLine(lines ...Line) *Table {
	p.lines = append(p.lines, lines...)
	return p
}

const (
	borderSize = 4
)

var (
	borderColor = drawing.Color{R: 234, G: 229, B: 227, A: 255}
)

func (p *Table) ToImg() (*bytes.Buffer, error) {
	return p.ToPng()
}

func (p *Table) ToJpg() (*bytes.Buffer, error) {
	img, err := p.toImg()
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	var b bytes.Buffer
	err = jpeg.Encode(&b, img, &jpeg.Options{
		Quality: 100,
	})
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &b, nil
}

func (p *Table) ToWebp() (*bytes.Buffer, error) {
	img, err := p.toImg()
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	var b bytes.Buffer
	err = webpbin.Encode(&b, img)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &b, nil
}

func (p *Table) ToPng() (*bytes.Buffer, error) {
	img, err := p.toImg()
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	var b bytes.Buffer
	err = png.Encode(&b, img)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return &b, nil
}

func (p *Table) toImg() (image.Image, error) {
	imgSize := NewSize(0, 0)

	// 找到最大的长宽
	colWidthMap := map[int]float64{}
	rowHighMap := map[int]float64{}
	for rowIdx, line := range p.lines {
		s := line.MinSize()
		if s.Width > imgSize.Width {
			imgSize.Width = s.Width
		}

		switch x := line.(type) {
		case CellsLine:
			err := x.RangeCell(func(colIdx int, cell Cell) error {
				size := cell.Size()
				if colWidthMap[colIdx] < size.Width {
					colWidthMap[colIdx] = size.Width
				}
				if rowHighMap[rowIdx] < size.Height {
					rowHighMap[rowIdx] = size.Height
				}
				return nil
			})
			if err != nil {
				log.Errorf("err:%v", err)
				return nil, err
			}
		}

		imgSize.Height += s.Height + borderSize
	}

	{
		var w float64
		for _, i := range colWidthMap {
			w += i + borderSize
		}

		if w > imgSize.Width {
			imgSize.Width = w
		}
	}

	{
		var h float64
		for _, i := range rowHighMap {
			h += i + borderSize
		}

		if h > imgSize.Height {
			imgSize.Height = h
		}
	}

	log.Debugf("img size:%v", imgSize)

	img := image.NewRGBA(image.Rect(0, 0, int(imgSize.Width), int(imgSize.Height)))
	draw.Draw(img, img.Bounds(), image.White, img.Bounds().Min, draw.Src)

	for rowIdx, line := range p.lines {
		switch x := line.(type) {
		case LineOne:
			log.Debugf("%d is line one", rowIdx)
			s := line.MinSize()
			if x.IsFull() {
				log.Debugf("set width is full %v", imgSize.Width)
				s.Width = imgSize.Width
			}
			line.setSize(NewSize(s.Width, s.Height))
		case CellsLine:
			err := x.RangeCell(func(colIdx int, cell Cell) error {
				cell.setSize(NewSize(colWidthMap[colIdx], rowHighMap[rowIdx]))
				return nil
			})
			if err != nil {
				log.Errorf("err:%v", err)
				return nil, err
			}
		}
	}

	// 对每一个cell进行绘图
	var y float64
	for idx, line := range p.lines {
		s := line.Size()
		log.Debugf("line:%v,size:%v", idx, s)

		DrawRectangleColor(img, borderColor, 0, y+s.Height, imgSize.Width, borderSize)

		err := line.DrawImg(0, y, img)
		if err != nil {
			log.Errorf("err:%v", err)
			return nil, err
		}
		y += s.Height + borderSize
	}

	{
		// 绘制水印
		if p.wmk != "" {
			const wmkFontSize = 60
			// TODO 多行的处理
			wmkSize := TextSize(p.wmk, wmkFontSize)

			for y := float64(0); y < imgSize.Height; y += wmkSize.Height + wmkFontSize*2 {
				for x := float64(0); x < imgSize.Width; x += wmkSize.Width + wmkFontSize*2 {
					DrawFont(img, image.NewUniform(drawing.Color{R: 228, G: 221, B: 184, A: 65}), x, y, p.wmk, wmkFontSize)
				}
			}
		}
	}

	if p.steganography != nil {
		var b bytes.Buffer
		err := steganography.Encode(&b, img, []byte(utils.ToString(p.steganography)))
		if err != nil {
			log.Errorf("err:%v", err)
			return nil, err
		}

		pngImg, err := png.Decode(&b)
		if err != nil {
			log.Errorf("err:%v", err)
			return nil, err
		}

		return pngImg, nil
	}

	return img, nil
}
