package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Printf("Valores iniciales A= %d, B=%d\n", a, b)
	intercambiarValores(&a, &b)
	fmt.Printf("Valores intercambiados A= %d, B=%d\n", a, b)
}

func intercambiarValores(ptrA *int, ptrB *int) {
	*ptrA, *ptrB = *ptrB, *ptrA
}
