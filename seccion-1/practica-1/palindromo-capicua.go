package main

import (
	"fmt"
)

func main() {

	fmt.Println("el numero 12321 es capicua?: ", esCapicua(12321))
	fmt.Println("el numero 45677 es capicua?: ", esCapicua(45677))
	fmt.Println("neuquen es palindromo?: ", esPalindromo("neuquen"))
	fmt.Println("hola es palindromo?: ", esPalindromo("hola"))

}

// Función para verificar si una cadena de texto es palíndromo
func esPalindromo(texto string) bool {
	longitud := len(texto) - 1

	for i := 0; i < longitud/2; i++ {
		if texto[i] != texto[longitud-i] {
			return false
		}
	}

	return true
}

func esCapicua(numero int) bool {
	numeroStr := fmt.Sprint(numero)
	return esPalindromo(numeroStr)
}
