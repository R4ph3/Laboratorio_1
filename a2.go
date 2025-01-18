package main

import (
	"fmt"
	"math"
)

func minmax(x float64, y float64) {

	//variables
	var lista []int

	min := math.Min(x, y)

	minInt := int(min)
	xInt := int(x)
	yInt := int(y)

	for i := 1; i < minInt; i++ {

		if xInt%i == 0 && yInt%i == 0 {
			lista = append(lista, i)

		}
	}

	if len(lista) > 0 {
		list0 := lista[0]

		for _, v := range lista {
			if list0 < v {
				list0 = v
			}
		}
		fmt.Println("El menor es:", list0)
	} else {

		fmt.Println("No se han encontrado divisores comunes")
	}

}

func main() {

	var num1 float64
	var num2 float64
	fmt.Println("Introduce el primer numero:")
	fmt.Scanln(&num1)
	fmt.Println("Introduce el segundo numero:")
	fmt.Scanln(&num2)

	minmax(num1, num2)

}
