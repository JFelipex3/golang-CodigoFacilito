package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Se lee línea por línea
func main() {
	fmt.Println("**************** INICIO ARCHIVO ****************")
	leido := readFile()
	fmt.Println("***************** FIN ARCHIVO *****************")
	fmt.Println("Archivo Leído correctamente: " + strconv.FormatBool(leido))
}

func readFile() bool {
	archivo, error := os.Open("./hola1.txt")

	// En caso de problemas o termino correcto se ejecuta el defer
	defer func() {
		archivo.Close()
	}()

	if error != nil {
		fmt.Println("Error:")
		panic(error)   // Imprime una descripción detallada del error
		r := recover() // Detiene el panic y puede continuar ejecución de la función que invocó readFile, pero no continua la ejecución de este último
		fmt.Println(r)
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea)
	}

	return true
}
