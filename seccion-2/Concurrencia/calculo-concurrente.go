package main

import (
	"fmt"
)

func main() {

	numGoRutines := 2

	doneChannel := make(chan int)

	final := 16777215999

	for i := 0; i < final; i = i + (final / numGoRutines) + 1 {
		paso := i + (final / numGoRutines)
		if paso > final {
			paso = final
		}

		fmt.Printf("Ejecutantdo %d %d\n", i, paso)
		go calcNums(i, paso, doneChannel)
	}

	doneNum := 0
	results := 0
	for doneNum < numGoRutines {
		value := <-doneChannel
		results = results + value
		fmt.Printf("Termino Uno, resultado por ahora: %d\n", results)
		doneNum++
	}
	fmt.Printf("Listo! Resultado final: %d\n", results)
}

func calcNums(start, end int, doneChannel chan int) {
	acumulador := 0
	for i := start; i <= end; i++ {
		acumulador = acumulador + i
	}
	doneChannel <- acumulador
}
