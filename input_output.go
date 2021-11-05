package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	edadF := 31
	fmt.Printf("Mi edad es %d\n", edadF) // Usar Printf para interpolar variables

	nombreF := "Jhonnatan"
	fmt.Printf("(Usando v) Mi nombre es %v\n", nombreF)
	fmt.Printf("(Usando s) Mi nombre es %s\n", nombreF)

	bandera := true
	fmt.Printf("El valor de Bandera es %t\n", bandera)

	precio := 14.3
	fmt.Printf("(Usando f) El Precio es %f\n", precio)
	fmt.Printf("(Usando e) El Precio es %e\n", precio)

	var edad int
	fmt.Print("Coloca tu edad: ")
	fmt.Scanf("%d\n", &edad)
	fmt.Printf("La edad es %d\n", edad)

	var nombre string
	fmt.Print("Coloca tu nombre: ")
	fmt.Scanf("%v\n", &nombre)
	fmt.Printf("El nombre es %v\n", nombre)

	//Utilizando bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingresa un Nuevo Nombre: ")
	text, err := reader.ReadString('\n')
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + text)
	}
}
