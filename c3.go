package main

import (
	"fmt"
	"math/big"
)

func main() {
	//a
	messageA := int64(5)
	eA := int64(5)
	pA := int64(53)
	resultA := new(big.Int).Exp(big.NewInt(messageA), big.NewInt(eA), big.NewInt(pA))
	fmt.Printf("(a) 5^5 mod 53 = %d\n", resultA)

	//b
	messageB := int64(4)
	eB := int64(11)
	pB := int64(79)
	resultB := new(big.Int).Exp(big.NewInt(messageB), big.NewInt(eB), big.NewInt(pB))
	fmt.Printf("(b) 4^11 mod 79 = %d\n", resultB)

	//c
	messageC := int64(101)
	eC := int64(7)
	pC := int64(293)
	resultC := new(big.Int).Exp(big.NewInt(messageC), big.NewInt(eC), big.NewInt(pC))
	fmt.Printf("(c) 101^7 mod 293 = %d\n", resultC)
}
