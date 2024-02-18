package main

import "fmt"

func main() {

	var x int
	fmt.Println("Por favor ingrese un numero: ")
	fmt.Scan(&x)
	esMayorMenorIgualACero(x)

}

func esMayorMenorIgualACero(x int) {
	if x > 0 {
		fmt.Println(x, " Es mayor a 0")
	} else if x < 0 {
		fmt.Println(x, " Es menor a 0")
	} else {
		fmt.Println(x, " Es igual a 0")
	}
}
