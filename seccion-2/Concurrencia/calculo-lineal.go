package main

import "fmt"

func main() {
	n := 16777215999
	sumaCompleta := 0

	for i := 1; i <= n; i++ {
		sumaCompleta = sumaCompleta + i
	}

	fmt.Println("La suma completa desde 1 hasta 16777215 es:", sumaCompleta)
}
