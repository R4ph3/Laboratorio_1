package main

import "fmt"

func main() {
	elec := true
	var numero int
	fmt.Println("Introduce un numero:")
	fmt.Scanln(&numero)

	if numero <= 1 {
		fmt.Println("Numero no valido")
		return
	}

	for i := 2; i < numero; i++ {
		if numero%i == 0 {
			fmt.Printf("El numero %v no es primo", numero)
			return
		} else {
			elec = false
		}
	}
	if elec == false {
		fmt.Printf("El numero %v es primo", numero)
	}
}
