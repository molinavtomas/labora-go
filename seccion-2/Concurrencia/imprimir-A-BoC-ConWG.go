package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	semA   = make(chan struct{})
	semBoC = make(chan struct{})
)

func escribirA(wg *sync.WaitGroup) {

	for {
		<-semA //Se bloquea la go rutine esperando recibir un valor, que se activa cuando lo manda desde BoC
		fmt.Print("A")
		semBoC <- struct{}{}
	}

}

func escribirBoC(wg *sync.WaitGroup) {

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

	var wg sync.WaitGroup

	// Incrementar el contador del WaitGroup
	wg.Add(2)

	// Iniciar las gorutinas
	go escribirA(&wg)
	go escribirBoC(&wg)

	// Iniciar la secuencia con una A
	semA <- struct{}{}

	// Esperar a que gorutinas terminen
	wg.Wait()
}
