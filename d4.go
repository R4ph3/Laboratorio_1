package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

const mrptNumTrials = 5

func PosiblePrimo(n int64) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	s := 0
	d := n - 1
	for d%2 == 0 {
		s++
		d /= 2
	}

	tryComposite := func(a int64) bool {
		if new(big.Int).Exp(big.NewInt(a), big.NewInt(d), big.NewInt(n)).Int64() == 1 {
			return false
		}
		for i := 0; i < s; i++ {
			if new(big.Int).Exp(big.NewInt(a), big.NewInt((1<<i)*d), big.NewInt(n)).Int64() == n-1 {
				return false
			}
		}
		return true
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < mrptNumTrials; i++ {
		a := rand.Int63n(n-2) + 2
		if tryComposite(a) {
			return false
		}
	}

	return true
}

func main() {
	testval := int64(97)

	if PosiblePrimo(testval) {
		fmt.Printf("%d is a prime\n", testval)
	} else {
		fmt.Printf("%d is not a prime\n", testval)
	}
}
