package main

import (
	"image"
	"log"
	"math/rand/v2"
	"os"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	permanent   bool
}

// Atlas texture
var texture *ebiten.Image
var collision []CollisionBlock
var currentPiece Piece = GenerateRandomPiece()

func AddVec2(first, second [2]int) [2]int {
	return [2]int{first[0] + second[0], first[1] + second[1]}
}

// Taken from stack overflow
func remove(s []CollisionBlock, i int) []CollisionBlock {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func GenerateRandomPiece() Piece {
	return Piece{
		rand.IntN(7),
		0,
		[2]int{-2, -10},
	}
}

func CheckBoard() {
	var positions [][2]int
	for _, block := range collision {
		if !block.permanent {
			positions = append(positions, block.position)
		}
	}
	var foundRows []int
	for y := HALF_HEIGHT; y > -HALF_HEIGHT; y-- {
		found := true
		for x := -HALF_WIDTH; x < HALF_WIDTH; x++ {
			testPos := [2]int{x, y}
			if !slices.Contains(positions, testPos) {
				found = false
				break
			}
		}
		if found {
			foundRows = append(foundRows, y)
		}
	}
	ClearLines(foundRows)
}

func ClearLines(lines []int) {
	slices.Sort(lines)

	for _, line := range lines {
		for i := 0; i < len(collision); {
			if collision[i].position[1] == line && !collision[i].permanent {
				collision = remove(collision, i)
			} else {
				i++
			}
		}
		for i := range collision {
			if collision[i].position[1] < line && !collision[i].permanent {
				collision[i].position[1]++
			}
		}
	}
}

func (g *Game) Update() error {
	PieceUpdate()
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
		collision = append(collision, CollisionBlock{7, [2]int{i, HALF_HEIGHT}, true})
	}
	for i := -HALF_HEIGHT; i < HALF_HEIGHT; i++ {
		collision = append(collision, CollisionBlock{7, [2]int{HALF_WIDTH, i}, true})
		collision = append(collision, CollisionBlock{7, [2]int{-HALF_WIDTH - 1, i}, true})

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
