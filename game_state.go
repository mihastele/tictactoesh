package main

import "fmt"

// Define a custom type

// Declare constants
const (
	Player1 int = 1 + iota
	Player2
	Tie
)

type GameState struct {
	victory   int
	board     [3][3]int
	moves     int
	MAX_MOVES int
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
