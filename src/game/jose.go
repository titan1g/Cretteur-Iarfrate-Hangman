package game

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GetHangmanAscii(filename string) (result []string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Convertit les bytes du fichier en string, puis
	// sépare les José dès qu'il y a deux sauts de lignes de suite.
	content := string(file)
	result = strings.Split(content, "\r\n\r\n")
	result = result[:len(result)-1]
	// Vérifie s'il y a bien 10 dessins de José.
	if len(result) != 10 {
		fmt.Printf("Invalid Hangman file: Should contain 10 elements, currently contains %v element(s).", len(result))
		os.Exit(1)
	}
	return result
}
