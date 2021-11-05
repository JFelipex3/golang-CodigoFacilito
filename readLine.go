package main

import (
	"bufio"
	"fmt"
	"os"
)

// Se lee línea por línea
func main() {
	archivo, error := os.Open("./hola.txt")

	if error != nil {
		fmt.Println(error)
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea)
	}

	archivo.Close() // Se debe cerrar el archivo
}
