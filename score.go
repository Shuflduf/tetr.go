package main

import "log"

var basic = map[int]int{
	1: 100,
	2: 300,
	3: 500,
	4: 800,
}

var tSpin = map[int]int{
	0: 400,
	1: 800,
	2: 1200,
	3: 1600,
}

var miniTSpin = map[int]int{
	0: 100,
	1: 200,
	2: 400,
}

var perfectClears map[int]int

// 0 = None, 1 = Mini, 2 = Full
var lastTSpin = 0
var linesJustCleared = 0
var score = 0
var highScore = 0

func UpdateScore() {
	log.Println(lastTSpin)
	switch lastTSpin {
	case 0:
		if linesJustCleared > 0 {
			score += basic[linesJustCleared]
		}
	case 1:
		score += miniTSpin[linesJustCleared]
	case 2:
		score += tSpin[linesJustCleared]
	}
	linesJustCleared = 0
	if score > highScore {
		highScore = score
	}
}
