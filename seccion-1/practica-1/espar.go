package main

import "fmt"

func main() {

	fmt.Println("17 es par?: ", esPar(17))
	fmt.Println("8 es par?: ", esPar(8))
}

func esPar(x int) bool {
	return x%2 == 0
}
