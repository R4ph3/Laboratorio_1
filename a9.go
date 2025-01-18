package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var str string
	fmt.Println("Dime la cadena a codificar")
	fmt.Scanln(&str)

	encoded := base64.StdEncoding.EncodeToString([]byte(str))

	fmt.Printf("Cadena original: %s\n", str)
	fmt.Printf("Codificado en Base64: %s\n", encoded)
}
