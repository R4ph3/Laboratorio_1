package main

import (
	"fmt"
	"math/big"
)

func main() {
	var messageStr, exponentStr, primeStr string

	fmt.Print("Ingrese el mensaje: ")
	fmt.Scanln(&messageStr)
	fmt.Print("Ingrese el exponente: ")
	fmt.Scanln(&exponentStr)
	fmt.Print("Ingrese el primo: ")
	fmt.Scanln(&primeStr)

	message := new(big.Int)
	exponent := new(big.Int)
	prime := new(big.Int)

	message.SetString(messageStr, 10)
	exponent.SetString(exponentStr, 10)
	prime.SetString(primeStr, 10)

	cipher := new(big.Int).Exp(message, exponent, prime)

	fmt.Printf("El mensaje cifrado es: %s\n", cipher.String())
}
