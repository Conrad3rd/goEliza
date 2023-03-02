package main

import (
	"fmt"

	"github.com/conrad3rd/goMods/mods/doctor"
	"github.com/inancgumus/screen"
)

func main() {
	// Clear all the characters on the screen
	screen.Clear()

	// Moves the cursor to the top-left position of the screen
	screen.MoveTopLeft()

	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)
}
