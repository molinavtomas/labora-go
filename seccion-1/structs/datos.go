package main

import (
	"fmt"
	"strings"
)

func main() {
	conjunto := Data{datos: make(map[int]int), id: 0}

	conjunto.Add(10)
	conjunto.Add(22)
	conjunto.Add(55)

	fmt.Println("Conjunto inicial:", conjunto.ToString())
	fmt.Println("Longitud:", fmt.Sprint(conjunto.Lenght()))

	conjunto.RemoveAt(2)
	fmt.Println("Conjunto después de remover el elemento 2, osea el 22:", conjunto.ToString())

	conjunto.Set(1, 77)
	fmt.Println("Conjunto después de modificar el elemento en la posición 1 con el numero 77:", conjunto.ToString())
}

type Data struct {
	datos map[int]int
	id    int
}

func (d *Data) Add(x int) {
	if d.datos == nil {
		//inicializo los datos
		d.datos = make(map[int]int)
		d.id = 0
	}
	d.id += 1
	d.datos[d.id] = x
}

func (d *Data) RemoveAt(i int) bool {
	_, existe := d.datos[i]

	if existe {
		delete(d.datos, i)
		return true
	} else {
		return false
	}
}

func (d *Data) Lenght() int {
	return d.id
}

func (d *Data) Set(i int, new_value int) bool {
	_, existe := d.datos[i]

	if existe {
		d.datos[i] = new_value
		return true
	} else {
		return false
	}
}

func (d *Data) ToString() string {
	var elementos []string

	for _, valor := range d.datos {
		elementos = append(elementos, fmt.Sprint(valor))
	}

	return strings.Join(elementos, ", ")
}
