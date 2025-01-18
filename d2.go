package main

import (
	"fmt"
	"os"
	"strconv"
)

func PrimosA1000(n int) []int {
	size := n / 2
	sieve := make([]bool, size)
	for i := range sieve {
		sieve[i] = true
	}

	limit := int(float64(n) / 2)
	for i := 1; i < limit; i++ {
		if sieve[i] {
			val := 2*i + 1
			for j := i + val; j < size; j += val {
				sieve[j] = false
			}
		}
	}

	primes := []int{2}
	for i, isPrime := range sieve {
		if isPrime && i > 0 {
			primes = append(primes, 2*i+1)
		}
	}

	return primes
}

func main() {
	test := 1000

	if len(os.Args) > 1 {
		if value, err := strconv.Atoi(os.Args[1]); err == nil {
			test = value
		}
	}

	primes := PrimosA1000(test)
	fmt.Println(primes)
}
