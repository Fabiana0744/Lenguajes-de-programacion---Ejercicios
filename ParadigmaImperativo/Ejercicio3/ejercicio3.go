package main

import (
	"fmt"
)

// EJERCICIO 3
func rotateSequence(sequence *[]interface{}, numRotations int, direction int) []interface{} {
	length := len(*sequence)

	//Validacion
	if length <= 1 || numRotations == 0 {
		return nil
	}

	numRotations %= length
	var rotatedSequence []interface{}

	if direction == 0 { // Rotacion a la izquierda
		rotatedSequence = append((*sequence)[numRotations:], (*sequence)[:numRotations]...)
	} else if direction == 1 {
		rotatedSequence = append((*sequence)[length-numRotations:], (*sequence)[:length-numRotations]...)
	}

	fmt.Println("Secuencia Original =", *sequence)
	fmt.Printf("Cantidad de rotaciones = %d, Dirección = %s\n", numRotations, map[int]string{0: "izq", 1: "der"}[direction])
	fmt.Println("Secuencia final rotada =", rotatedSequence)
	fmt.Println()
	return rotatedSequence
}

func main() {
	originalSequence := []interface{}{"a", "b", "c", "d", "e", "f", "g", "h"}
	rotateSequence(&originalSequence, 3, 0) // Izquierda
	rotateSequence(&originalSequence, 2, 1) // Derecha
	rotateSequence(&originalSequence, 5, 0) // Izquierda
}
