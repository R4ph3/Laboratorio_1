package main

import (
	"fmt"
	"strconv"
)

func main() {
	// valor inicial
	var val1 int
	fmt.Println("Dime el valor a modificar")
	fmt.Scanln(&val1)

	// decimal
	fmt.Printf("Dec:\t%d\n", val1)

	// binario
	binary := strconv.FormatInt(int64(val1), 2)
	fmt.Printf("Bin:\t%s\n", binary)

	// hexa
	hexadecimal := strconv.FormatInt(int64(val1), 16)
	fmt.Printf("Hex:\t0x%s\n", hexadecimal)

	// octal
	octal := strconv.FormatInt(int64(val1), 8)
	fmt.Printf("Oct:\t0o%s\n", octal)

	// Caracter (si está en el rango ASCII imprimible)
	if val1 >= 32 && val1 <= 126 {
		fmt.Printf("Char:\t%c\n", val1)
	} else {
		fmt.Println("Char:\t(Not printable ASCII)")
	}
}
