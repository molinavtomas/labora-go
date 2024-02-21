package main

import "fmt"

var (
	sem_A  = make(chan struct{})
	sem_B  = make(chan struct{})
	sem_C  = make(chan struct{})
	sem_BC = make(chan struct{})
)

func escribirA() {
	for {
		<-sem_A //Se bloquea la go rutine esperando recibir un valor, que se activa cuando lo manda desde BoC
		fmt.Print("A")
		sem_BC <- struct{}{}
	}
}

func escribirB() {
	for {
		<-sem_B //Se bloquea la go rutine esperando recibir un valor, que se activa cuando lo manda desde BoC
		<-sem_BC
		fmt.Print("B")
		sem_A <- struct{}{}
		sem_C <- struct{}{}
	}
}

func escribirC() {
	for {
		<-sem_C
		<-sem_BC
		fmt.Print("C")
		sem_A <- struct{}{}
		sem_B <- struct{}{}
	}
}

func main() {

	go escribirB()
	go escribirA()
	go escribirC()

	sem_B <- struct{}{}
	sem_BC <- struct{}{}

	select {}

}
