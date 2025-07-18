package main

import "fmt"

var playerColorMapToBashColor = map[rune]string{
	'w': "\u001B[97m",
	'r': "\u001B[31m",
	'g': "\u001B[32m",
	'b': "\u001B[34m",
	'y': "\u001B[33m",
	'c': "\u001B[36m",
	'm': "\u001B[35m",
	'k': "\u001B[30m",
	'p': "\u001B[37m",
}

var playerSprites [2][4]string = [2][4]string{
	{"X  X", " XX ", " XX ", "X  X"},
	{" OO ", "O  O", "O  O", " OO "},
}

func drawBoard(state GameState) {
	fmt.Println("")
}
