package main

import (
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

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
func PieceUpdate() {
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if currentPiece.CanMove([2]int{-1, 0}) {
			currentPiece.position[0] -= 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if currentPiece.CanMove([2]int{1, 0}) {
			currentPiece.position[0] += 1
		}
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
				currentPiece = GetNextPiece()
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
		currentPiece = GetNextPiece()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		newRotIndex := (currentPiece.rotationIndex + 3) % 4
		if currentPiece.CanRotate(newRotIndex) {
			currentPiece.rotationIndex = newRotIndex
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		newRotIndex := (currentPiece.rotationIndex + 1) % 4
		if currentPiece.CanRotate(newRotIndex) {
			currentPiece.rotationIndex = newRotIndex
		}
	}
}
