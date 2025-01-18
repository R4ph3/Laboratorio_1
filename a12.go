package main

import (
	"fmt"
)

func main() {

	var num1 int
	var list []int
	fmt.Println("Dime el numero para factorizar:")
	fmt.Scanln(&num1)

	for i := 2; i*i <= num1; i++ {

		for num1%i == 0 {
			list = append(list, i)
			num1 /= i
		}

	}
	if num1 > 1 {
		list = append(list, num1)

	}
	fmt.Println("Numeros: ", list)

}
