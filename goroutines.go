package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombrelento("Jhonnatan") // Si no se indica go el proceso se realizara sincrono, en caso de definirla se ejecuta asincrono o concurrente
	fmt.Println("Al fin termino!!!")

	// Definir una rutina anónima
	go func() {
		var wait string
		fmt.Scanln(&wait)
	}()
}

func miNombrelento(name string) {
	letras := strings.Split(name, "") // Mandar cadena vacia "" separa todas las letras del string

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
