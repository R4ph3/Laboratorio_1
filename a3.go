package main

import (
	b64 "encoding/b64"
	"encoding/hex"
	"fmt"
)

func pasarb64(frase string) {

	codificacion := b64.URLEncoding.EncodeToString([]byte(frase))
	fmt.Printf("%v", codificacion)

}

func pasarhex(frase string) {

	src := []byte(frase)
	encodedStr := hex.EncodeToString(src)
	fmt.Printf("%v", encodedStr)
}

func main() {

	var frase string

	fmt.Println("Dime la frase:")
	fmt.Scanln(&frase)

	pasarb64(frase)
	pasarhex(frase)

}
