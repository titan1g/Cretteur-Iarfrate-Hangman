package game

import "strings"

type Game struct {
	State        string   // "playing", "won", "lost"
	Letters      []string // letters in the word
	FoundLetters []string // letters found
	UsedLetters  []string // letters used
	TurnsLeft    int      // remaining turns
}

func New(turns int, word string) *Game {
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
	}
	return g
}
