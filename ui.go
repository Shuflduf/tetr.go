package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var defaultFont text.Face

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

// func DrawText(screen *ebiten.Image) {
// 	ebitenutil.DebugPrint(screen, "NO")
// }

func UI() *ebiten.Image {
	img := ebiten.NewImage(200, 200)
	text.Draw(img, "AAAA", defaultFont, &text.DrawOptions{})
	return img
}
