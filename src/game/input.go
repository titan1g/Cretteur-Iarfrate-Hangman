package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {
	valid := false
	for !valid {
		fmt.Print("Guess a letter or a word :")
		guess, err = reader.ReadString('\n')
		if err != nil {
			return guess, err
		}
		guess = strings.TrimSpace(guess)
		if len(guess) == 0 {
			fmt.Println("Invalid guess.")
			continue
		}

		valid = true
	}
	return guess, nil
}
