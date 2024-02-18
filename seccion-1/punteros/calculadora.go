package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Print("Ingrese el primer numero: ")
	fmt.Scan(&x)
	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scan(&y)

	calcular(&x, &y)

}

func calcular(a *int, b *int) {
	suma := *a + *b
	resta := *a - *b
	multiplicacion := *a * *b
	division := *a / *b

	fmt.Printf("Suma: %d\n", suma)
	fmt.Printf("Resta: %d\n", resta)
	fmt.Printf("multiplicacion: %d\n", multiplicacion)
	fmt.Printf("Division: %d\n", division)

}
