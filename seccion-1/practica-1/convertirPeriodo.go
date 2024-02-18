package main

import "fmt"

func main() {
	// Ejemplos de uso
	periodo1 := 1030
	d1, h1, m1, s1 := convertirPeriodo(periodo1)
	fmt.Printf("%d segundos son %d días, %d horas, %d minutos con %d segundos\n", periodo1, d1, h1, m1, s1)

	periodo2 := 12045
	d2, h2, m2, s2 := convertirPeriodo(periodo2)
	fmt.Printf("%d segundos son %d días, %d horas, %d minutos con %d segundos\n", periodo2, d2, h2, m2, s2)

	periodo3 := 176520
	d3, h3, m3, s3 := convertirPeriodo(periodo3)
	fmt.Printf("%d segundos son %d días, %d horas, %d minutos con %d segundos\n", periodo3, d3, h3, m3, s3)
}

func convertirPeriodo(segundos int) (dias, horas, minutos, segundosRestantes int) {
	dias = segundos / 86400 // Un día tiene 86400 segundos
	horas = (segundos % 86400) / 3600
	minutos = (segundos % 3600) / 60
	segundosRestantes = segundos % 60
	return dias, horas, minutos, segundosRestantes
}
