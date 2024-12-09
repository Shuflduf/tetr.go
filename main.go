package main

import (
    "image/color"
    "log"
    "os"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
    testImg *ebiten.Image
}

func (g *Game) Update() error {
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    drawOptions := ebiten.DrawImageOptions{}
    drawOptions.GeoM.Translate(50.0, 100.0)
    screen.DrawImage(g.testImg, &drawOptions)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 320, 240
}

func main() {
    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("Hello, World!")

    // Load the image from a file
    f, err := os.Open("path/to/your/image.png")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    img, _, err := ebitenutil.NewImageFromReader(f)
    if err != nil {
        log.Fatal(err)
    }

    game := &Game{
        testImg: img,
    }

    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
