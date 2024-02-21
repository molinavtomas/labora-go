package main

import (
	"fmt"
	"sync"
)

type Carrera struct {
	wg      sync.WaitGroup
	pistola chan struct{}
}

func NuevaCarrera() *Carrera {
	return &Carrera{
		pistola: make(chan struct{}),
	}
}

func (p *Carrera) Esperar() {
	<-p.pistola
}

func (p *Carrera) Empezar() {
	close(p.pistola)
}

func jugador(color string, carrera *Carrera) {
	defer carrera.wg.Done()

	carrera.Esperar()
	fmt.Printf("Jugador %s listo para correr\n", color)

	// Simulamos la carrera realizando operaciones ficticias
	for i := 0; i < 1000; i++ {
	}

	fmt.Printf("Jugador %s llegó a la meta\n", color)
}

func main() {
	carrera := NuevaCarrera()

	// Agregar 3 al WaitGroup, uno por cada goroutine de jugador
	carrera.wg.Add(3)

	go jugador("Rojo", carrera)
	go jugador("Azul", carrera)
	go jugador("Verde", carrera)

	var inicio string
	fmt.Print("Presione cualquier tecla para iniciar la carrera: ")
	fmt.Scan(&inicio)

	// Imprimir mensaje de inicio de la carrera
	fmt.Println("¡Comienza la carrera!")

	// Permitir que todas las goroutines comiencen al mismo tiempo
	carrera.Empezar()

	// Esperar a que todos los jugadores terminen
	carrera.wg.Wait()

	fmt.Println("Carrera terminada.")
}
