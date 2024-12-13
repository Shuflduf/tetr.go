package main

import (
	"math/rand/v2"
	"slices"
)

const BAG_SIZE = 7

var nextPieces []Piece
var bag []Piece
var startingPos = [2]int{-2, -10}

func InitBag() {
	bag = []Piece{}
	for i := 0; i < BAG_SIZE; i++ {
		bag = append(bag, Piece{
			i,
			0,
			startingPos,
		})
	}
	rand.Shuffle(len(bag), func(i, j int) {
		bag[i], bag[j] = bag[j], bag[i]
	})
}

func GetNextPiece() Piece {
	if len(bag) <= 0 {
		InitBag()
	}
	lastValue := bag[len(bag)-1]
	bag = slices.Delete(bag, len(bag)-1, len(bag))
	return lastValue
}
