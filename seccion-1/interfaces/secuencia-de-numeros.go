package main

import (
	"fmt"
	"math/rand"
)

func main() {

	if rand.Intn(2) == 0 {
		var numCrec = numerosCrecientes{}

		numCrec.Title()
		//Pruebo agregar valores
		numCrec.Next()
		numCrec.Next()
		numCrec.Next()
		imprimirArray(numCrec.valores)

		//imprimo los primeros 30, no importa que se hayan agregado valores antes
		numCrec.Primeros30()
		imprimirArray(numCrec.valores)

		//pruebo con imprimir el proximo valor, el 31
		numCrec.Next()
		imprimirArray(numCrec.valores)
	} else {
		//Ahora con numeros primos

		var numPrim = numerosPrimos{}

		numPrim.Title()
		numPrim.Next()
		numPrim.Next()
		numPrim.Next()
		imprimirArray(numPrim.valores)

		numPrim.Primeros30()
		imprimirArray(numPrim.valores)

		numPrim.Next()
		imprimirArray(numPrim.valores)
	}

}

type IntSequence interface {
	Next() int
	Primeros30() []int
	Title()
}

type numerosCrecientes struct {
	valores []int
}

type numerosPrimos struct {
	valores []int
}

func (obj *numerosCrecientes) Next() int {
	proximo := len(obj.valores) + 1

	obj.valores = append(obj.valores, proximo)

	return proximo
}

func (obj *numerosPrimos) Next() int {

	//agrega el primer elemento
	if obj.valores == nil {
		obj.valores = append(obj.valores, 1)
	}

	ultimoElemento := obj.valores[len(obj.valores)-1]
	nuevoNumero := ultimoElemento + 1
	for {

		if esPrimo(nuevoNumero) {
			obj.valores = append(obj.valores, nuevoNumero)
			break
		} else {
			nuevoNumero++
		}

	}

	return nuevoNumero
}

func (obj *numerosCrecientes) Primeros30() []int {

	for len(obj.valores) < 30 {
		obj.Next()
	}

	return obj.valores
}

func (obj *numerosPrimos) Primeros30() []int {

	for len(obj.valores) < 30 {
		obj.Next()
	}

	return obj.valores
}

func (obj *numerosCrecientes) Title() {
	fmt.Println("Secuencia de numeros crecientes")
}

func (obj *numerosPrimos) Title() {
	fmt.Println("Secuencia de numeros Primos")
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

func imprimirArray(arr []int) {
	fmt.Print("[")
	for i, valor := range arr {
		fmt.Print(valor)
		if i < len(arr)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}
