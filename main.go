package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//go:embed assets/texture_simple.png
var textureData []byte

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
var currentPiece Piece
var nextPiece Piece
var nullPiece = Piece{-1, 0, [2]int{0, 0}}

// Initiliaze as null piece (index is -1)
var heldPiece = nullPiece
var justHeld = false

func AddVec2(first, second [2]int) [2]int {
	return [2]int{first[0] + second[0], first[1] + second[1]}
}

// Taken from stack overflow
func remove(s []CollisionBlock, i int) []CollisionBlock {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
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

func ResetGame() {
	collision = []CollisionBlock{}
	heldPiece = nullPiece
	gameInit()
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
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyShift) && !justHeld {
		if heldPiece == nullPiece {
			heldPiece = currentPiece
			heldPiece.position = [2]int{0, 0}
			heldPiece.rotationIndex = 0
			currentPiece = nextPiece
			nextPiece = GetNextPiece()
		} else {
			t := currentPiece
			currentPiece = heldPiece
			heldPiece = t
			heldPiece.position = [2]int{0, 0}
			currentPiece.position = startingPos
		}
		justHeld = true
		UpdateGhost()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(UI(), &ebiten.DrawImageOptions{})
	offsetGrid := [2]int{14, 10}
	// Draw grid
	for _, block := range collision {
		drawOptions := ebiten.DrawImageOptions{}
		drawOptions.GeoM.Translate(float64(block.position[0]+offsetGrid[0])*32, float64(block.position[1]+offsetGrid[1])*32)
		pieceIndex := block.colourIndex
		cropRect := image.Rect(32*pieceIndex, 0, 32*(pieceIndex+1), 32)
		screen.DrawImage(texture.SubImage(cropRect).(*ebiten.Image), &drawOptions)
	}
	// Draw current piece
	for _, pos := range PIECES[currentPiece.colourIndex][currentPiece.rotationIndex] {
		drawOptions := ebiten.DrawImageOptions{}
		drawOptions.GeoM.Translate(
			float64(pos[0]+offsetGrid[0]+currentPiece.position[0])*32,
			float64(pos[1]+offsetGrid[1]+currentPiece.position[1])*32,
		)
		value := (((((float64(lockDelayTimer) / float64(lockDelay)) - 0.5) * -1) + 0.5) * 0.7) + 0.3
		// yeah this is deprecated
		drawOptions.ColorM.ChangeHSV(0.0, 1.0, value)
		pieceIndex := currentPiece.colourIndex
		cropRect := image.Rect(32*pieceIndex, 0, 32*(pieceIndex+1), 32)
		screen.DrawImage(texture.SubImage(cropRect).(*ebiten.Image), &drawOptions)

		// Reuse code to draw ghost
		drawOptions.ColorScale.ScaleAlpha(0.2)
		drawOptions.GeoM.Translate(
			0.0,
			float64(ghostPieceHeight)*32,
		)
		screen.DrawImage(texture.SubImage(cropRect).(*ebiten.Image), &drawOptions)
	}
	// Draw next piece
	for _, pos := range PIECES[nextPiece.colourIndex][0] {
		drawOptions := ebiten.DrawImageOptions{}
		drawOptions.GeoM.Translate(
			float64(7+pos[0]+offsetGrid[0])*32,
			float64(-8+pos[1]+offsetGrid[1])*32,
		)
		pieceIndex := nextPiece.colourIndex
		cropRect := image.Rect(32*pieceIndex, 0, 32*(pieceIndex+1), 32)
		screen.DrawImage(texture.SubImage(cropRect).(*ebiten.Image), &drawOptions)
	}
	// Draw hold piece
	if heldPiece != nullPiece {
		for _, pos := range PIECES[heldPiece.colourIndex][0] {
			drawOptions := ebiten.DrawImageOptions{}
			drawOptions.GeoM.Translate(
				float64(-11+pos[0]+offsetGrid[0])*32,
				float64(-8+pos[1]+offsetGrid[1])*32,
			)
			if justHeld {
				drawOptions.ColorScale.ScaleAlpha(0.2)
			}
			pieceIndex := heldPiece.colourIndex
			cropRect := image.Rect(32*pieceIndex, 0, 32*(pieceIndex+1), 32)
			screen.DrawImage(texture.SubImage(cropRect).(*ebiten.Image), &drawOptions)
		}
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
	InitBag()
	currentPiece = GetNextPiece()
	nextPiece = GetNextPiece()
	UpdateGhost()
}

func main() {
	ebiten.SetWindowSize(1152, 864)
	ebiten.SetWindowTitle("Hello, World!")
	// ebiten.SetFullscreen(true)

	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(textureData))
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
