package main

import (
	"fmt"
	"math"
)

func getIfPrime(val int) bool {
	if val <= 1 {
		return false
	}
	if val <= 3 {
		return true
	}
	if val%2 == 0 || val%3 == 0 {
		return false //borra multiplos 2 y 3
	}

	max := int(math.Sqrt(float64(val)))
	for k := 1; ; k++ {
		testVal := 6*k - 1
		if testVal > max {
			break
		}
		if val%testVal == 0 {
			return false
		}

		testVal = 6*k + 1
		if testVal > max {
			break
		}
		if val%testVal == 0 {
			return false
		}
	}
	return true
}

func main() {
	val := 93
	if getIfPrime(val) {
		fmt.Printf("%d is prime\n", val)
	} else {
		fmt.Printf("%d is not prime\n", val)
	}
}
