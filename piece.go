package main

import (
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var movingDir = [2]bool{false, false}
var dirTimers = [2]int{0, 0}
var ghostPieceHeight = -1

// How many frames between each auto shift
var arr int = 2

// The delay before the piece starts auto shifting
var das int = 10

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

func (p *Piece) CanRotate(rotIndex int) bool {
	var positions [][2]int
	for _, block := range collision {
		positions = append(positions, block.position)
	}
	for _, block_offset := range PIECES[p.colourIndex][rotIndex] {
		block_position := [2]int{block_offset[0] + p.position[0], block_offset[1] + p.position[1]}
		if slices.Contains(positions, block_position) {
			return false
		}
	}
	return true
}

func UpdateInputs() {
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		MovePiece([2]int{-1, 0})
		movingDir[0] = true
		movingDir[1] = false
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		MovePiece([2]int{1, 0})
		movingDir[0] = false
		movingDir[1] = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyA) {
		movingDir[0] = false
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			MovePiece([2]int{-1, 0})
			movingDir[1] = true
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyD) {
		movingDir[1] = false
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			MovePiece([2]int{1, 0})
			movingDir[0] = true
		}
	}
}

func PieceUpdate() {
	UpdateInputs()
	if movingDir[0] {
		dirTimers[0]++
		if dirTimers[0] > das {
			if dirTimers[0]%arr == 0 {
				MovePiece([2]int{-1, 0})
			}
		}
	} else {
		dirTimers[0] = 0
	}
	if movingDir[1] {
		dirTimers[1]++
		if dirTimers[1] > das {
			if dirTimers[1]%arr == 0 {
				MovePiece([2]int{1, 0})
			}
		}
	} else {
		dirTimers[1] = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if currentPiece.CanMove([2]int{0, 1}) {
			currentPiece.position[1] += 1
			if !currentPiece.CanMove([2]int{0, 1}) {
				for _, pos := range PIECES[currentPiece.colourIndex][currentPiece.rotationIndex] {
					blockPos := AddVec2(pos, currentPiece.position)
					collision = append(collision, CollisionBlock{currentPiece.colourIndex, blockPos, false})
				}
				CheckBoard()
				currentPiece = nextPiece
				nextPiece = GetNextPiece()
        UpdateGhost()
			}
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		for {
			if !currentPiece.CanMove([2]int{0, 1}) {
				break
			}
			currentPiece.position[1] += 1
		}
		for _, pos := range PIECES[currentPiece.colourIndex][currentPiece.rotationIndex] {
			blockPos := AddVec2(pos, currentPiece.position)
			collision = append(collision, CollisionBlock{currentPiece.colourIndex, blockPos, false})
		}
		CheckBoard()
		currentPiece = nextPiece
		nextPiece = GetNextPiece()
    UpdateGhost()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		newRotIndex := (currentPiece.rotationIndex + 3) % 4
		if currentPiece.CanRotate(newRotIndex) {
			currentPiece.rotationIndex = newRotIndex
      UpdateGhost()
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		newRotIndex := (currentPiece.rotationIndex + 1) % 4
		if currentPiece.CanRotate(newRotIndex) {
			currentPiece.rotationIndex = newRotIndex
      UpdateGhost()
		}
	}
}

func MovePiece(dir [2]int) {
	if currentPiece.CanMove(dir) {
		currentPiece.position[0] += dir[0]
		currentPiece.position[1] += dir[1]
    UpdateGhost()
  }
}

func UpdateGhost() {
  ghostPieceHeight = -1
  for {
    if currentPiece.CanMove([2]int{0, ghostPieceHeight + 1}) {
      ghostPieceHeight++
    } else {
      break
    }
  }
}
