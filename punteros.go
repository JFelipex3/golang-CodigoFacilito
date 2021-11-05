package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		Es una dirección de memoria
		En lugar del valor, tenermos la dirección en la que está el valor

		Ej: X, Y apuntan a as123d con valor 5
			se modifica X a valor 6
			Y tiene como valor 6

		*T, donde T corresponde al tipo de dato. int, string, Struct
	*/

	var x, y *int
	entero := 5

	x = &entero // agregar & accede a la dirección de memoria
	y = &entero

	fmt.Println("Valor x: " + strconv.Itoa(*x) + ", Valor y: " + strconv.Itoa(*y))

	*x = 6
	fmt.Println("Se modifica x a 6")
	fmt.Println("Valor x: " + strconv.Itoa(*x) + ", Valor y: " + strconv.Itoa(*y))
}
