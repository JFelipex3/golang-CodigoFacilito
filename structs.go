package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func (p User) nombreCompleto() string {
	return p.nombre + " " + p.apellido
}

//En este caso se pasa como una copia
func (p User) setName(n string) {
	p.nombre = n
}

//En este caso se para un puntero
func (p *User) setNamePuntero(n string) {
	p.nombre = n
}

func main() {
	var persona User

	persona.edad = 23
	persona.nombre = "Carlos"
	persona.apellido = "Carrera"

	fmt.Println(persona)
	persona.setName("Mario")
	fmt.Println(persona.nombreCompleto())
	persona.setNamePuntero("Mario")
	fmt.Println(persona.nombreCompleto())

	persona2 := User{nombre: "Eduardo", apellido: "Carmona"}
	fmt.Println(persona2)

	persona3 := User{26, "María", "Garrido"}
	fmt.Println(persona3)

	persona4 := new(User) // retona un puntero
	persona4.nombre = "Marcelo"
	fmt.Println(persona4)
	fmt.Println(persona4.nombre)
}
