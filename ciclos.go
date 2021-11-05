package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Ciclo FOR con formato i := 0; i < 10; i++")
	for i := 0; i < 3; i++ {
		fmt.Println("Iteración N°" + strconv.Itoa(i))
	}

	fmt.Print("\n")
	fmt.Println("Ciclo FOR simulando un While")
	x := 0
	for x < 3 {
		fmt.Println("Iteración N°" + strconv.Itoa(x))
		x++
	}

	fmt.Print("\n")
	fmt.Println("Ciclo For infinito + Break")
	y := 0
	for y < 3 {
		fmt.Println("Iteración N°" + strconv.Itoa(y))
		y++

		if y > 3 {
			break //Finaliza el ciclo
		}
	}

	fmt.Print("\n")
	fmt.Println("Ciclo For con Continue")
	for i := 0; i < 5; i++ {

		if i == 2 {
			// Salta la iteración
			fmt.Println("Salta Iteración")
			continue // vuelve al inicio del ciclo para la nueva iteración
		}
		fmt.Println("Iteración N°" + strconv.Itoa(i))
	}
}
