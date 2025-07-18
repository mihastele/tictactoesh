package main

import "fmt"

// Define a custom type

// Declare constants
const (
	Player1 int = 1 + iota
	Player2
	Tie
)

/*
victory -> 0 game is on, 1 player1 winds, 2 player2 wins, 3 Tie
board -> board pieces: 0 empty, 1 player1, 2 player2
moves -> moves played
MAX_MOVES -> bad constant to track when the game is over to avoid infinite loop
*/
type GameState struct {
	terminalTextColor rune
	playerColors      [2]rune
	victory           int
	board             [3][3]int
	moves             int
	MAX_MOVES         int
}

func initGameState() GameState {
	return GameState{
		victory:   0,
		board:     [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		moves:     0,
		MAX_MOVES: 9,
	}
}

func assignPlayerMove(gs *GameState, player int, move string) error {
	if len(move) < 2 {
		return fmt.Errorf("invalid move format")
	}
	moveRune := []rune(move)
	letterRune := moveRune[0]
	digitRune := moveRune[1]

	// Convert letter to coordX
	var coordX int
	switch letterRune {
	case 'A', 'a':
		coordX = 0
	case 'B', 'b':
		coordX = 1
	case 'C', 'c':
		coordX = 2
	default:
		return fmt.Errorf("invalid letter")
	}

	// Convert digit to coordY
	if digitRune < '1' || digitRune > '9' {
		return fmt.Errorf("invalid digit")
	}
	coordY := int(digitRune - '1') // '1' -> 0, '2' ->1, etc.

	if gs.board[coordX][coordY] != 0 {
		return fmt.Errorf("square already occupied")
	}
	gs.board[coordX][coordY] = player
	gs.moves++
	return nil
}
