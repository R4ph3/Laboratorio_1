package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

func decode_b64(data string) {

	decoded, err := base64.StdEncoding.DecodeString(data)
	//toca poner la salida de error porque la libreria esta devuelve dos valores
	//el valor decodificado, y si rula o no, asi que si lo intentas poner con una variable solo te da error
	if err != nil {
		log.Fatal("error:", err)
	}

	fmt.Printf("%q\n", decoded)

}

func decode_hex(data string) {

	decoded, err := hex.DecodeString(data)
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		//la funcion DecodeString devuelve el valor en bytes, asi que hay que poner %s para que lo traduzca a string, si no no funca
		fmt.Printf("%s", decoded)
	}

}

func decode_bites(bits string) (string, error) {
	//hay qye separar los bits en grupos de 8 para decodificalos
	if len(bits)%8 != 0 {
		return "", fmt.Errorf("la longitud de la cadena de bits no es multiplo de 8")
	}

	var result string

	//pillamos cada byte (8 bits)
	for i := 0; i < len(bits); i += 8 {
		//pillamos un bbyte
		byteStr := bits[i : i+8]

		//Lo pasamos a entero
		byteVal, err := strconv.ParseUint(byteStr, 2, 8)
		if err != nil {
			return "", err
		}

		//Pasamos el entero a ascii
		result += string(byteVal)
	}

	return result, nil
}
func main() {

	var data string
	var eleccion string

	fmt.Println(`Elige la decodificacion:
	
				[*]Base64
				[*]Hexa
				[*]Bites
				`)

	fmt.Scan(&eleccion)
	fmt.Println("Dime la data")
	fmt.Scanln(&data)

	switch eleccion {
	case "Base64":
		decode_b64(data)
	case "Hexa":
		decode_hex(data)
	case "Bites":
		ascii, err := decode_bites(data)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(ascii)

	}

}
