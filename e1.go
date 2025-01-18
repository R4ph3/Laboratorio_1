package main

import (
	"fmt"
	"strings"
)

func genLinear(a, seed, c, m int) string {
	x := seed
	var results []string

	for i := 0; i < 200; i++ {
		val := (a*x + c) % m
		results = append(results, fmt.Sprintf("%d", val))
		x = val
	}

	return strings.Join(results, " ")
}

func main() {
	a := 21
	X0 := 35
	c := 31
	m := 100

	res := genLinear(a, X0, c, m)

	fmt.Println(res)
}
