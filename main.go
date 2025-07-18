package main

import (
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	gs := initGameState()
	//fmt.Println(playerColor1, playerColor2)

	fmt.Println("What is your terminal text color?(b/w default -> w[white]): ")
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

	}

	//fmt.Println(playerColor2)

	////TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	//// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	//s := "gopher"
	//fmt.Printf("Hello and welcome, %s!\n", s)
	//
	//for i := 1; i <= 5; i++ {
	////TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
	//// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
	//fmt.Println("i =", 100/i)
	//}
}
