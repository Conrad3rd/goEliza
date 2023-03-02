package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/conrad3rd/goMods/mods/doctor"
	"github.com/inancgumus/screen"
)

func main() {
	// Clear all the characters on the screen
	screen.Clear()
	// Moves the cursor to the top-left position of the screen
	screen.MoveTopLeft()

	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	userInput, _ := reader.ReadString('\n')

	fmt.Println(userInput)
}
