package main

import (
	"log"

	"github.com/fogleman/gg"
	"github.com/goki/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	width    = 2500
	height   = 1686
	boxWidth = width / 3
)

func main() {
	dc := gg.NewContext(width, height)
	dc.SetHexColor("#FFFFFF")
	dc.Clear()

	page2(dc)

	dc.SavePNG("richmenu.png")
}

func text(dc *gg.Context) {
	// if err := dc.LoadFontFace("./FiraSans-Light.ttf", 96); err != nil {
	// 	panic(err)
	// }

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}
	face := truetype.NewFace(font, &truetype.Options{Size: 96})
	dc.SetFontFace(face)

	text := "Hello World 👍"
	textMargin := 60.0
	maxWidth := float64(width) - textMargin

	dc.SetRGB(0, 0, 0)
	dc.DrawStringWrapped(text, float64(width)/2, float64(height)/2, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
}

func page1(dc *gg.Context) {
	dc.SetRGB(255.0, 255.0, 255.0)
	dc.SetLineWidth(5.0)
	dc.DrawLine(0, float64(height)/2, float64(width), float64(height)/2)
	dc.DrawLine(float64(boxWidth), float64(height)/2, float64(boxWidth), float64(height))
	dc.DrawLine(float64(boxWidth*2), 0, float64(boxWidth*2), float64(height))
	dc.Stroke()

	dc.SavePNG("richmenu.png")
}

func page2(dc *gg.Context) {
	dc.SetRGB(255.0, 255.0, 255.0)
	dc.SetLineWidth(5.0)
	dc.DrawLine(0, float64(height)/2, float64(width), float64(height)/2)
	dc.DrawLine(float64(boxWidth), 0, float64(boxWidth), float64(height))
	dc.DrawLine(float64(boxWidth*2), 0, float64(boxWidth*2), float64(height))
	dc.Stroke()
}
