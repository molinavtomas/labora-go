package main

import "fmt"

func main() {
	fmt.Println("El par 220 y 284 son numeros amigos?: ", sonNumerosAmigos(220, 284))
}

func sonNumerosAmigos(x int, y int) bool {
	if obtenerSumaDivisores(x) > y || obtenerSumaDivisores(x) < y {
		return false
	} else if obtenerSumaDivisores(y) > x || obtenerSumaDivisores(y) < x {
		return false
	} else if obtenerSumaDivisores(x) == y && obtenerSumaDivisores(y) == x {
		return true
	} else {
		return true
	}
}

func obtenerSumaDivisores(numero int) int {
	var divisores int

	for i := 1; i < numero; i++ {
		if numero%i == 0 {
			divisores += i
		}
	}

	//fmt.Println("suma divisores de ", numero, " es: ", divisores)

	return divisores
}
