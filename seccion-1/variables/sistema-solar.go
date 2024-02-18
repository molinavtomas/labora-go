package main

import "fmt"

func main() {
	// Información de los planetas y sus lunas
	planetas := map[string]int{
		"Mercurio": 0,
		"Venus":    0,
		"Tierra":   1,
		"Marte":    2,
		"Júpiter":  79,
		"Saturno":  83,
		"Urano":    27,
		"Neptuno":  14,
	}

	// Mostrar la información de los planetas y sus lunas
	for planeta, lunas := range planetas {
		fmt.Printf("%s tiene %d lunas.\n", planeta, lunas)
	}
}
