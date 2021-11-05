package main

import (
	"fmt"
	"io/ioutil"
)

// El contenido del txt se carga completamente en memoria
func main() {
	file_data, err := ioutil.ReadFile("./hola.txt") // La ruta es relativa desde el directorio que se ejecute

	if err != nil {
		fmt.Println("Hay un error")
	}

	fmt.Println(string(file_data))
}
