package main

import (
	"fmt"
	"hangman/game"
	"os"
)

func main() {
	g := game.New(8, "Golang")
	fmt.Println(g)

	l, err := game.ReadGuess()
	if err != nil {
		fmt.Println("Error reading guess: %v", err)
		os.Exit(1)
	}
	fmt.Println(l)
}
