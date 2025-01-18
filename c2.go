package main

import (
	"fmt"
	"math"
)

// Function to check if a number is prime
func isPrime(p int) bool {
	if p <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(p))); i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

// Function to compute Me mod p
func modP(M, e, p int) int {
	return int(math.Pow(float64(M), float64(e))) % p
}

func main() {
	var M, e, p int

	// Input values for M, e, and p
	fmt.Print("Enter the value for M: ")
	fmt.Scanln(&M)
	fmt.Print("Enter the value for e: ")
	fmt.Scanln(&e)
	fmt.Print("Enter the value for p: ")
	fmt.Scanln(&p)

	// Check if p is a prime number
	if !isPrime(p) {
		fmt.Printf("The number %d is not a prime number.\n", p)
		return
	}

	// Compute M^e mod p
	result := modP(M, e, p)
	fmt.Printf("The result of M^e mod p is: %d\n", result)
}
