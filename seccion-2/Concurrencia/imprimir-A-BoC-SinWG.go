package main

import (
	"fmt"
	"math/rand"
)

var (
	semA   = make(chan struct{})
	semBoC = make(chan struct{})
)

func escribirA() {
	for {
		<-semA //Se bloquea la go rutine esperando recibir un valor, que se activa cuando lo manda desde BoC
		fmt.Print("A")
		semBoC <- struct{}{}
	}

}

func escribirBoC() {
	for {
		<-semBoC
		if rand.Intn(2) == 0 {
			fmt.Print("B")
		} else {
			fmt.Print("C")
		}
		semA <- struct{}{}
	}
}

func main() {

	// Iniciar las gorutinas
	go escribirA()
	go escribirBoC()

	// Iniciar la secuencia con una A
	semA <- struct{}{}
	select {}

}
