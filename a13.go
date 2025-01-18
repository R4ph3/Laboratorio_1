package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"log"
)

func main() {
	// Base64 encoded strings
	base64Strings := []string{
		"eJzzyc9Lyc8DAAgpAms=",
		"eJxzSi3KycwDAAfXAl0=",
		"eJzzSy1XiMwvygYADKUC8A==",
	}

	// Loop through each string, decode and decompress
	for _, b64 := range base64Strings {
		decoded, err := base64.StdEncoding.DecodeString(b64)
		if err != nil {
			log.Fatalf("Error decoding Base64: %v", err)
		}

		decompressed, err := decompress(decoded)
		if err != nil {
			log.Fatalf("Error decompressing data: %v", err)
		}

		fmt.Printf("Uncompressed Output: %s\n", decompressed)
	}
}

// Decompress the zlib-compressed data
func decompress(data []byte) (string, error) {
	reader, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	defer reader.Close()

	var out bytes.Buffer
	_, err = io.Copy(&out, reader)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
