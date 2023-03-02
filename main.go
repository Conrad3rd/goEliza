package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			fmt.Println(doctor.Response(userInput))
			break
		}
		fmt.Println(doctor.Response(userInput))

	}
}
