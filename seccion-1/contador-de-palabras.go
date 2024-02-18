package main

import (
	"fmt"
	"strings"
)

func main() {
	//wc.Test(WordCount("I ate a donut. Then I ate another donut."))
	imprimirMapa(WordCount("I ate a donut. Then I ate another donut."))
}

func WordCount(s string) map[string]int {
	// Dividir la cadena en palabras
	palabras := strings.Fields(s)

	contador := make(map[string]int)

	for _, palabra := range palabras {
		contador[palabra]++
	}

	return contador
}

func imprimirMapa(m map[string]int) {
	fmt.Println("Contenido del mapa:")
	for clave, valor := range m {
		fmt.Printf("%s: %d\n", clave, valor)
	}
}
