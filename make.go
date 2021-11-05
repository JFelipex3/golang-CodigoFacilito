package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		Estructura de un Slice
		Puntero --> Apunta a un arreglo
		Longitud es la cantidad de elementos
		Capacidad
	*/
	slice := make([]int, 3, 5) // make(tipo, Longitud, Capacity); Capacity es opcional
	fmt.Println("Largo del Slice: " + strconv.Itoa(len(slice)))
	fmt.Println("Capacity del Slice: " + strconv.Itoa(cap(slice)))
	fmt.Println(slice)

	fmt.Print("\n")

	// Agregar un nuevo elemento
	slice = append(slice, 2)
	fmt.Println("Largo del Slice: " + strconv.Itoa(len(slice)))
	fmt.Println("Capacity del Slice: " + strconv.Itoa(cap(slice)))
	fmt.Println(slice)

	fmt.Print("\n")

	// Copiar arreglos. copy(destino, fuente)
	/*
		La función copy copia el mínimo de elementos en ambos arreglos
	*/
	slice2 := []int{1, 2, 3, 4}
	sliceCopia := make([]int, 3) // forma correcta: make([]int, len(slice2))

	fmt.Println("Antes de Realizar Copia")
	fmt.Println(slice2)
	fmt.Println(sliceCopia)

	fmt.Print("\n")

	fmt.Println("Después de Realizar Copia")
	copy(sliceCopia, slice2)
	fmt.Println(slice2)
	fmt.Println(sliceCopia)

}
