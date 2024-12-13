package main

var basic = map[int]int{
	1: 100,
	2: 300,
	3: 500,
	4: 800,
}

var tSpins map[int]int
var miniTSpins map[int]int
var perfectClears map[int]int

var linesJustCleared = 0
var score = 0
var highScore = 0

func UpdateScore() {
	if linesJustCleared != 0 {
		score += basic[linesJustCleared]
	}
	linesJustCleared = 0
  if score > highScore {
    highScore = score
  }
}
