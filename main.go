package main

import (
	"fmt"
)

func main() {

	gs := initGameState()
	//fmt.Println(playerColor1, playerColor2)

	fmt.Println("What is your terminal text color?(b/w/g default -> w[white]): ")
	_, err := fmt.Scanf("%c", &gs.terminalTextColor)

	if err != nil {
		gs.terminalTextColor = 'w'
	}

	fmt.Println("Select player 1 color (w/r/g/b/y/c/m/k): ")

	_, err = fmt.Scanf("%c", &gs.playerColors[0])

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	//fmt.Println(playerColor1)

	fmt.Println("Select player 2 color (w/r/g/b/y/c/m/k): ")
	_, err = fmt.Scanf("%c", &gs.playerColors[1])

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for gs.victory < 1 {
		if gs.moves >= gs.MAX_MOVES {
			gs.victory = Tie // tied state
		}

		fmt.Println("Enter a move (Examples A1, B3, C2, ...)")
		fmt.Printf("Player %d move: ", gs.player+1)
		var move string
		fmt.Scanf("%s", &move)

		err := assignPlayerMove(&gs, move)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		drawBoard(gs)
		gs.checkVictory()

	}
}
