package main

import "fmt"

func main() {

	fmt.Println("17 es primo?: ", esPrimo(17))
	fmt.Println("24 es primo?:", esPrimo(24))
}

func esPrimo(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
