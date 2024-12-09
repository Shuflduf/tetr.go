package main

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	texture   *ebiten.Image
	testValue int
}

func (g *Game) Update() error {
	g.testValue += 1
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawOptions := ebiten.DrawImageOptions{}
	drawOptions.GeoM.Translate(50.0, float64(g.testValue))
	screen.DrawImage(g.texture.SubImage(image.Rect(0, 0, 30, 0)).(*ebiten.Image), &drawOptions)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	// Load the image from a file
	f, err := os.Open("Tetr-Skin.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := ebitenutil.NewImageFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		texture: img,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
