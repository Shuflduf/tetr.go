package main

import (
	"image"
	"log"
	"os"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const BOARD_WIDTH = 10
const BOARD_HEIGHT = 20
const HALF_WIDTH = BOARD_WIDTH / 2
const HALF_HEIGHT = BOARD_HEIGHT / 2

type Game struct{}
type Piece struct {
	colourIndex   int
	rotationIndex int
	position      [2]int
}
type CollisionBlock struct {
	colourIndex int
	position    [2]int
}

// Atlas texture
var texture *ebiten.Image
var collision []CollisionBlock
var currentPiece Piece = Piece{
	0,
	0,
	[2]int{0, 0},
}

func (p *Piece) CanMove(dir [2]int) bool {
	var positions [][2]int
	for _, block := range collision {
		positions = append(positions, block.position)
	}
	for _, block_offset := range PIECES[p.colourIndex][p.rotationIndex] {
		block_position := [2]int{block_offset[0] + p.position[0] + dir[0], block_offset[1] + p.position[1] + dir[1]}
		if slices.Contains(positions, block_position) {
			return false
		}
	}
	return true
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if currentPiece.CanMove([2]int{-1, 0}) {
			currentPiece.position[0] -= 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if currentPiece.CanMove([2]int{1, 0}) {
			currentPiece.position[0] += 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		currentPiece.position[1] += 1
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	offsetGrid := [2]int{14, 10}
	for _, block := range collision {
		drawOptions := ebiten.DrawImageOptions{}
		drawOptions.GeoM.Translate(float64(block.position[0]+offsetGrid[0])*32, float64(block.position[1]+offsetGrid[1])*32)
		pieceIndex := block.colourIndex
		cropRect := image.Rect(32*pieceIndex, 0, 32*(pieceIndex+1), 32)
		screen.DrawImage(texture.SubImage(cropRect).(*ebiten.Image), &drawOptions)
	}
	for _, pos := range PIECES[currentPiece.colourIndex][currentPiece.rotationIndex] {
		drawOptions := ebiten.DrawImageOptions{}
		drawOptions.GeoM.Translate(
			float64(pos[0]+offsetGrid[0]+currentPiece.position[0])*32,
			float64(pos[1]+offsetGrid[1]+currentPiece.position[1])*32,
		)
		pieceIndex := currentPiece.colourIndex
		cropRect := image.Rect(32*pieceIndex, 0, 32*(pieceIndex+1), 32)
		screen.DrawImage(texture.SubImage(cropRect).(*ebiten.Image), &drawOptions)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 896, 672
}

func gameInit() {
	for i := -HALF_WIDTH - 1; i < HALF_WIDTH+1; i++ {
		collision = append(collision, CollisionBlock{7, [2]int{i, HALF_HEIGHT}})
	}
	for i := -HALF_HEIGHT; i < HALF_HEIGHT; i++ {
		collision = append(collision, CollisionBlock{7, [2]int{HALF_WIDTH, i}})
		collision = append(collision, CollisionBlock{7, [2]int{-HALF_WIDTH - 1, i}})

	}
}

func main() {
	ebiten.SetWindowSize(1152, 864)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetFullscreen(true)

	// Load the image from a file
	f, err := os.Open("assets/texture_simple.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := ebitenutil.NewImageFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	texture = img
	game := &Game{}

	gameInit()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
