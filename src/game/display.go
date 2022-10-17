package game

import "fmt"

func DrawWelcome() {
	fmt.Println("Welcome to Hangman!")

}

func Draw(g *Game, guess string) {
	drawTurn(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `   
		+---+  
		|   |  
		O   |  
	   /|\  |  
	   / \  |  
			|  
	  =========`
	case 1:
		`  
	+---+  
	|   |  
	O   |  
   /|\  |  
   /    |  
		|  
  =========
  `
	case 2:
		`  
	+---+  
	|   |  
	O   |  
   /|\  |  
		|  
		|  
  =========
  `
	case 3:
		`
	+---+  
	|   |  
	O   |  
   /|   |  
		|  
		|  
  =========
  `
	case 4:
		`  
	+---+  
	|   |  
	O   |  
	|   |  
		|  
		|  
  =========
  `
	case 5:
		`
	+---+  
	|   |  
	O   |  
		|  
		|  
		|  
  =========
  `
	case 6:
		`
	+---+  
	|   |  
		|  
		|  
		|  
		|  
  =========
  `
	case 7:
		`
	+---+  
	|  
	|  
	|  
	|  
	|  
=========
`
	case 8:
		`
	|  
	|  
	|  
	|  
	|  
=========
`
	case 9:
		`         
=========
`
	}
	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Println("Guessed: ")

	fmt.Println("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Println("Good guess!")
	case "alreadyGuessed":
		fmt.Println("You already guessed that letter!")
	case "badGuess":
		fmt.Println("Bad guess! %s is not in the word.", guess)
	case "lost":
		fmt.Println("You lost! The word was :")
		drawLetters(g.Letters)
	case "won":
		fmt.Println("You won!")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%s ", c)
	}
	fmt.Println()
}
