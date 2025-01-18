package main

import (
	"fmt"
)

func minmax(x int, y int) {
	var lista_x []int
	var lista_y []int
	var coprimo bool
	coprimo = false

	for i := 1; i <= x; i++ {
		if x%i == 0 {
			lista_x = append(lista_x, i)
		}
	}

	if len(lista_x) > 0 {
		lista_x = lista_x[1:]
	}

	for i := 1; i <= y; i++ {
		if y%i == 0 {
			lista_y = append(lista_y, i)
		}
	}

	if len(lista_y) > 0 {
		lista_y = lista_y[1:]
	}

	if len(lista_x) != len(lista_y) {
		fmt.Println("No son co-primos, las listas de divisores son de diferentes longitudes).")
		return
	}

	for i := range lista_x {
		if lista_x[i] == lista_y[i] {
			coprimo = true
			break
		}
	}

	if coprimo {
		fmt.Println("Son co-primos")
	} else {
		fmt.Println("No son co-primos")
	}
}

func main() {
	minmax(5432, 634)
}
