package main

import (
	"fmt"
)

func main() {

	var num uint
	fmt.Println("Dime el numero que quieres:")
	fmt.Scanln(&num)

	fmt.Println("Decimal form:", num)
	fmt.Println("Shift left (1):", num<<1)
	fmt.Println("Shift left (2):", num<<2)
	fmt.Println("Shift right(1):", num>>1)
	fmt.Println("Shift right(2):", num>>2)

}
