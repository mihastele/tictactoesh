package main

type GameState struct {
	victory int
	board   [3][3]int
}

func initGameState() GameState {
	return GameState{
		victory: 0,
		board:   [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
	}
}
