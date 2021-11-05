package main

import "fmt"

func main() {
	// var nombre string; declaración explicita
	nombre := "Coco" // Declaración implicita
	numero := 1

	var x, z, y int
	x, z, y = 1, 2, 3

	fmt.Println(nombre)
	fmt.Println(numero)
	fmt.Println(x, z, y)
}
