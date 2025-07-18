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

/*
index:
0 -> player1 sprite
1 -> player2 sprite
each subarray is a row forming a 5x5 image
*/
var playerSprites [2][5]string = [2][5]string{
	{"X   X", " X X ", "  X  ", " X X ", "X   X"},
	{" OOO ", "O   O", "O   O", "O   O", " OOO "},
}

func drawBoard(state GameState) {
	fmt.Println("    1  |  2  |  3 ")
	fmt.Println("  -----|-----|-----")
	drawBoardLine(state, 0)
	drawBoardLine(state, 1)
	drawBoardLine(state, 2)
	fmt.Printf("\n\n")
}

func drawBoardLine(gs GameState, index int) {
	for i := 0; i < 5; i++ {
		if i == 2 {
			var rightLabel rune
			switch index {
			case 0:
				rightLabel = 'A'
			case 1:
				rightLabel = 'B'
			case 2:
				rightLabel = 'C'
			}
			fmt.Printf("%c ", rightLabel)
		} else {
			fmt.Print("  ")
		}

		for j := 0; j < 3; j++ {
			if gs.board[index][j] == 0 {
				fmt.Print("     ")
			} else {
				player := gs.board[index][j] - 1
				fmt.Printf("%v%v%v",
					playerColorMapToBashColor[gs.playerColors[player]],
					playerSprites[player][i],
					playerColorMapToBashColor['w'], // reset color, assuming black background with white text
				)
			}

			if i != 2 {
				fmt.Printf("|")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("  -----|-----|-----")
}
