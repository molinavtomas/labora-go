package main

import "fmt"

func main() {

	var persona1 = Persona{nombre: "Juan", edad: 20, ciudad: "Madrid", telefono: 555 - 1234}
	var persona2 = Persona{nombre: "Ana", edad: 20, ciudad: "Barcelona", telefono: 555 - 5678}

	fmt.Println("Persona 1: ", persona1.obtenerDatos())
	fmt.Println("Persona 2: ", persona2.obtenerDatos())

	persona1.cambiarCiudad("Paris")
	fmt.Println("Persona 1: ", persona1.obtenerDatos())

	persona1.cambiarCiudad("Paris")

}

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono int
}

func (p *Persona) obtenerDatos() string {
	return fmt.Sprintf("%s, %d, %s, %d", p.nombre, p.edad, p.ciudad, p.telefono)
}

func (p *Persona) cambiarCiudad(s string) {
	if !(p.ciudad == s) {
		p.ciudad = s
	} else {
		fmt.Println("Las ciudades son iguales")
	}
}
