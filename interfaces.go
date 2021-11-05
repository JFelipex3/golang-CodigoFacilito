package main

import "fmt"

type User interface {
	Permisos() int // 1 - 5
	Nombre() string
}

// El implementar los permisos de la interfaz indica que la estructura utiliza la interfaz
type Admin struct {
	nombre string
}

func (this Admin) Permisos() int {
	return 5
}

func (this Admin) Nombre() string {
	return this.nombre
}

// El implementar los permisos de la interfaz indica que la estructura utiliza la interfaz
type Editor struct {
	nombre string
}

func (this Editor) Permisos() int {
	return 3
}

func (this Editor) Nombre() string {
	return this.nombre
}

func auth(this User) string {
	msj := this.Nombre() + " no tiene permisos de administrador"
	if this.Permisos() > 4 {
		msj = this.Nombre() + " tiene permisos de administrador"
	}

	return msj
}

func main() {
	admin := Admin{"Gabriel"}
	fmt.Println(auth(admin))

	editor := Editor{"Marcos"}
	fmt.Println(auth(editor))

	// Permite usar la interfaz como tipo de dato
	usuarios := []User{Admin{"Marcela"}, Editor{"Carolina"}}

	for _, usuario := range usuarios {
		fmt.Println(auth(usuario))
	}
}
