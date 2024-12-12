package main

import (
	"log"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var movingDir = [2]bool{false, false}
var dirTimers = [2]int{0, 0}
var ghostPieceHeight = -1
var lockDelay = 30
var maxLockDelay = 120
var lockDelayTimer = 0
var maxLockDelayTimer = 0
var onGround = false
var gravityDelay = 60
var gravityDelayTimer = 0
var softDropping = false

// How many frames between each auto shift
var arr = 2

// The delay before the piece starts auto shifting
var das = 10

// How much gravity gets multiplied by when soft dropping
var sdf = 10

func (p *Piece) Moved(dir [2]int) Piece {
	return Piece{
		p.colourIndex,
		p.rotationIndex,
		AddVec2(p.position, dir),
	}
}

func (p *Piece) Rotated(newRotIndex int) Piece {
	return Piece{
		p.colourIndex,
		newRotIndex,
		p.position,
	}
}

// Change current piece to next piece, and move current piece to grid
func (p *Piece) SetPiece() {
	lockDelayTimer = 0
	maxLockDelayTimer = 0
	onGround = false
	for _, pos := range PIECES[p.colourIndex][p.rotationIndex] {
		blockPos := AddVec2(pos, p.position)
		collision = append(collision, CollisionBlock{p.colourIndex, blockPos, false})
	}
	CheckBoard()
	currentPiece = nextPiece
	nextPiece = GetNextPiece()
	justHeld = false
	UpdateGhost()
	if !IsFree(currentPiece) {
		ResetGame()
	}
}

func (p *Piece) TouchingGround() bool {
	return !IsFree(p.Moved([2]int{0, 1}))
}

func IsFree(p Piece) bool {
	var positions [][2]int
	for _, block := range collision {
		positions = append(positions, block.position)
	}
	for _, block_offset := range PIECES[p.colourIndex][p.rotationIndex] {
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
	softDropping = ebiten.IsKeyPressed(ebiten.KeyW)
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		gravityDelayTimer = gravityDelay
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
	gravityDelayTimer++
	actualGravityDelay := gravityDelay
	if softDropping {
		actualGravityDelay /= sdf
	}
	if gravityDelayTimer > actualGravityDelay {
		if IsFree(currentPiece.Moved([2]int{0, 1})) {
			currentPiece.position[1] += 1
			UpdateGhost()
		}
		gravityDelayTimer = 0
	}
	if currentPiece.TouchingGround() {
		lockDelayTimer++
		onGround = true
		if lockDelayTimer > lockDelay {
			currentPiece.SetPiece()
		}
	} else {
		lockDelayTimer = 0
	}
	if onGround {
		maxLockDelayTimer++
		if maxLockDelayTimer > maxLockDelay {
			HardDrop()
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		HardDrop()
	}
	var newRotIndex = currentPiece.rotationIndex
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		newRotIndex = (currentPiece.rotationIndex + 3) % 4
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		newRotIndex = (currentPiece.rotationIndex + 1) % 4
	}
	if newRotIndex != currentPiece.rotationIndex {
		if IsFree(currentPiece.Rotated(newRotIndex)) {
			currentPiece.rotationIndex = newRotIndex
			UpdateGhost()
			return
		}
		kickIndex := GetKickIndex(currentPiece.rotationIndex, newRotIndex)
		newPieceUnrotated := currentPiece.Rotated(newRotIndex)
		var table [8][4][2]int
		if currentPiece.colourIndex == 4 {
			table = I_KICKS
		} else {
			table = KICKS
		}
		log.Println("Before: ", currentPiece.rotationIndex, " After: ", newRotIndex, " Which is: ", kickIndex)
		for _, kick := range table[kickIndex] {
			flippedKick := [2]int{kick[0], -kick[1]}
			newPiece := newPieceUnrotated.Moved(flippedKick)
			if IsFree(newPiece) {
				currentPiece.rotationIndex = newRotIndex
				currentPiece.position = AddVec2(currentPiece.position, flippedKick)
				UpdateGhost()
				lockDelayTimer = 0
				return
			}
		}
	}
}

func GetKickIndex(before, after int) int {
	var kickIndex int
	if after == (before+1)%4 {
		kickIndex = before * 2
	} else if after == (before+3)%4 {
		kickIndex = (before*2 + 7) % 8
	}
	return kickIndex
}

func HardDrop() {
	for {
		if !IsFree(currentPiece.Moved([2]int{0, 1})) {
			break
		}
		currentPiece.position[1] += 1
	}
	currentPiece.SetPiece()
}

func MovePiece(dir [2]int) {
	if IsFree(currentPiece.Moved(dir)) {
		currentPiece.position[0] += dir[0]
		currentPiece.position[1] += dir[1]
		UpdateGhost()
	}
	lockDelayTimer = 0
}

func UpdateGhost() {
	ghostPieceHeight = -1
	for {
		if IsFree(currentPiece.Moved([2]int{0, ghostPieceHeight + 1})) {
			ghostPieceHeight++
		} else {
			break
		}
	}
}
