package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var defaultFont text.Face
var controlsHidden = false

func init() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	fontFace, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	defaultFont = text.NewGoXFace(fontFace)
}

func UIUpdate() {
  if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
    x, y := ebiten.CursorPosition()
    if x < 200 && y > 500 {
      controlsHidden = true
    }
  }
}

func UI() *ebiten.Image {
	img := ebiten.NewImage(896, 672)
	drawOptions := text.DrawOptions{}
	drawOptions.GeoM.Translate(10, 10)
	text.Draw(img, "TETR.GO", defaultFont, &drawOptions)
	Controls(img)
	return img
}

func Controls(img *ebiten.Image) {
  if controlsHidden {
    return
  }
	const HEIGHT = 30
	const SCALE = 0.6
	offset := [2]int{30, 860}
	controlsList := []string{
		"A - LEFT",
		"D - RIGHT",
		"L ARROW - CCW",
		"R ARROW - CW",
		"S - SOFT",
		"W - HARD",
		"SHIFT - HOLD",
    "CLICK TO HIDE",
	}
	for i, item := range controlsList {
		drawOptions := text.DrawOptions{}
		drawOptions.GeoM.Translate(float64(offset[0]), float64((i*HEIGHT)+offset[1]))
		drawOptions.GeoM.Scale(SCALE, SCALE)
		drawOptions.ColorScale.Scale(0.3, 0.3, 0.3, 1.0)
		text.Draw(img, item, defaultFont, &drawOptions)
	}
}
