package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4}

	fmt.Println(slice)

	slice2 := []int{}

	if slice2 == nil {
		fmt.Println("Está Vacío")
	}

	/*
		Un Slice es un puntero al arreglo
		Longitud del arreglo al que apunta
		Capacidad
	*/

	arreglo := [3]int{1, 2, 3}
	fmt.Println("Arreglo para Slice")
	fmt.Println(arreglo)
	slice3 := arreglo[1:2] //Tomar valores de un arreglo para convertirlo en slice, solo toma valor de la posición 1 (valor = 2) y la posición 2 (valor = 3) no se considera
	fmt.Println(slice3)
	slice4 := arreglo[0:] // Toma todo el arreglo
	fmt.Println(slice4)
}
