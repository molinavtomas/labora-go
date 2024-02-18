package main

import "fmt"

func main() {
	var temp1, temp2, temp3 int

	fmt.Print("Ingrese la temperatura 1: ")
	fmt.Scan(&temp1)

	fmt.Print("Ingrese la temperatura 2: ")
	fmt.Scan(&temp2)

	fmt.Print("Ingrese la temperatura 3: ")
	fmt.Scan(&temp3)

	suma := temp1 + temp2 + temp3
	promedio := suma / 3

	fmt.Printf("La suma de las temperaturas es: %d \n", suma)
	fmt.Printf("El promedio de las temperaturas es: %d\n", promedio)
}
