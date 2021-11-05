package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 31

	edad_str := strconv.Itoa(edad)

	fmt.Println("Mi edad es " + edad_str)

	edad_int, _ := strconv.Atoi(edad_str) // _ se utiliza para un retorno que no necesito

	fmt.Println(edad_int + 10)
}
