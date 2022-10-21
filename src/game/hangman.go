package game

import (
	"fmt"
	"strings"
)

type Game struct {
	State        string   // "playing", "won", "lost"
	Letters      []string // letters in the word
	FoundLetters []string // letters found
	UsedLetters  []string // letters used
	TurnsLeft    int      // remaining turns
	Hangman      []string // Jose
}

func New(turns int, word string) (*Game, error) {
	if len(word) < 3 {
		return nil, fmt.Errorf("word '%s' must be at least 3 characters. got=%v", word, len(word))
	}

	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
		Hangman:      GetHangmanAscii("hangman.txt"),
	}

	return g, nil
}

func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)

	switch g.State {
	case "won", "lost":
		return
	}

	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)

		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	} else {
		g.State = "badGuess"
		g.LoseTurn(guess)

		if g.TurnsLeft <= 0 {
			g.State = "lost"
		}
	}
}

func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}

func (g *Game) LoseTurn(guess string) {
	g.TurnsLeft--
	g.UsedLetters = append(g.UsedLetters, guess)
}

func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}

func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}
