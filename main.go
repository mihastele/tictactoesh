package main

import "fmt"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	var playerColor1, playerColor2 rune
	//fmt.Println(playerColor1, playerColor2)

	fmt.Println("Select player 1 color (w/r/g/b/y/c/m/k): ")

	_, err := fmt.Scanf("%c", &playerColor1)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	//fmt.Println(playerColor1)

	fmt.Println("Select player 2 color (w/r/g/b/y/c/m/k): ")
	_, err = fmt.Scanf("%c", &playerColor2)

	if err != nil {
		fmt.Println("Error: ", err)
	}

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
