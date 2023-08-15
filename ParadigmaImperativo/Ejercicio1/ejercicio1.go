package main

import (
	"fmt"
	"strings"
)

// EJERCICIO 1

func countCharacters(text string) int {
	numCharacters := len(text)
	return numCharacters
}

func countWords(text string) int {
	words := strings.Fields(text) // Devuelve una cadena con el texto
	numWords := len(words)
	return numWords
}

func countLines(text string) int {
	numLines := strings.Count(text, "\n") + 1
	return numLines
}

func main() {
	text := "Este es un ejemplo de texto.\nTiene varias líneas.\nCada línea tiene varias palabras."
	fmt.Printf("El texto es: \n%s\n\n", text)
	numCharacters := countCharacters(text)
	fmt.Printf("Número de caracteres: %d\n", numCharacters)
	numWords := countWords(text)
	fmt.Printf("Número de palabras: %d\n", numWords)
	numLines := countLines(text)
	fmt.Printf("Número de líneas: %d\n", numLines)
}
