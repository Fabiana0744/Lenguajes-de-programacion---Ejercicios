package main

import (
	"fmt"
	"math"
	"strings"
)

// EJERCICIO 2
func printFigure(centerLineLen int) {
	// Validacion centerLineLen valido
	if centerLineLen <= 0 || centerLineLen%2 == 0 {
		fmt.Println("La cantidad de elementos de la línea del centro debe ser un número impar positivo.")
		return
	}

	for i := 0; i < centerLineLen; i++ {
		spaces := int(math.Abs(float64(centerLineLen/2 - i))) // Cantidad de espacios en blanco para poner rombo
		stars := centerLineLen - 2*spaces                     // Cantidad de asteriscos
		fmt.Printf("%s%s\n", strings.Repeat(" ", spaces), strings.Repeat("*", stars))
	}
}

func main() {
	lineSize := 5 // Cambia este valor para ajustar el tamaño del rombo
	printFigure(lineSize)
}
