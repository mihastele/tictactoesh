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
	player            int
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
		player:    0,
		MAX_MOVES: 9,
	}
}

func assignPlayerMove(gs *GameState, move string) error {
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
		return fmt.Errorf("invalid move (Letter overflow)")
	}

	// Convert digit to coordY
	if digitRune < '1' || digitRune > '9' {
		return fmt.Errorf("invalid move (Digit overflow)")
	}
	coordY := int(digitRune - '1') // '1' -> 0, '2' ->1, etc.

	if gs.board[coordX][coordY] != 0 {
		return fmt.Errorf("square already occupied")
	}
	gs.board[coordX][coordY] = gs.player
	gs.moves++
	return nil
}

func (g *GameState) checkVictory() {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if g.board[i][0] != 0 && g.board[i][0] == g.board[i][1] && g.board[i][1] == g.board[i][2] {
			g.victory = g.board[i][0] // Return the winner
		}
		if g.board[0][i] != 0 && g.board[0][i] == g.board[1][i] && g.board[1][i] == g.board[2][i] {
			g.victory = g.board[0][i]
		}
	}

	// Check diagonals
	if g.board[0][0] != 0 && g.board[0][0] == g.board[1][1] && g.board[1][1] == g.board[2][2] {
		g.victory = g.board[0][0]
	}
	if g.board[0][2] != 0 && g.board[0][2] == g.board[1][1] && g.board[1][1] == g.board[2][0] {
		g.victory = g.board[0][2]
	}

	// Check for tie (if all moves played and no winner)
	if g.moves >= 9 {
		g.victory = 3 // Tie
	}

	g.victory = 0 // Game is still ongoing
}

//func (g *GameState) checkVictory() int {
//	// Check rows and columns
//	for i := 0; i < 3; i++ {
//		if g.board[i][0] != 0 && g.board[i][0] == g.board[i][1] && g.board[i][1] == g.board[i][2] {
//			return g.board[i][0] // Return the winner
//		}
//		if g.board[0][i] != 0 && g.board[0][i] == g.board[1][i] && g.board[1][i] == g.board[2][i] {
//			return g.board[0][i]
//		}
//	}
//
//	// Check diagonals
//	if g.board[0][0] != 0 && g.board[0][0] == g.board[1][1] && g.board[1][1] == g.board[2][2] {
//		return g.board[0][0]
//	}
//	if g.board[0][2] != 0 && g.board[0][2] == g.board[1][1] && g.board[1][1] == g.board[2][0] {
//		return g.board[0][2]
//	}
//
//	// Check for tie (if all moves played and no winner)
//	if g.moves >= 9 {
//		return 3 // Tie
//	}
//
//	return 0 // Game is still ongoing
//}
