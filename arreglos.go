package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var arreglo [10] int Inicialización tamaño fijo, valor por defecto del tipo en caso de int es valor 0

	arreglo := [3]int{1, 2, 3}

	fmt.Println("Largo Arreglo: " + strconv.Itoa(len(arreglo)))
	fmt.Println(arreglo)

	fmt.Print("\n")
	fmt.Println("Matriz")
	var matriz [3][2]int
	matriz[1][1] = 3
	fmt.Println(matriz)
}
