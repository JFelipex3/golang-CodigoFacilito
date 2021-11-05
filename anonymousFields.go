package main

import "fmt"

type Human struct {
	name string
}

func (this Human) hablarHumano() string {
	return "Soy un Humano"
}

func (this Human) hablar() string {
	return "Soy un Humano"
}

type Tutor struct {
	Human //Campos de la estructura Human, simula la herencia
	rango string
}

func (this Tutor) hablar() string {
	return "Soy un Tutor que sobre escribe"
}

func (this Tutor) hablarConHumano() string {
	return this.Human.hablar() + " Soy un Tutor que sobre escribe"
}

func main() {
	tutor := Tutor{Human{"Gonzalo"}, "Rango 1"}
	fmt.Println(tutor) // tambien se puede acceder como tutor.Human.name, esto es útil cuando hay dos estructuras con el mismo nombre
	fmt.Println(tutor.hablarHumano())
	fmt.Println(tutor.hablar())
	fmt.Println(tutor.hablarConHumano())
}
