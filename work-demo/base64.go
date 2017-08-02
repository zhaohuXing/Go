package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	input := []byte("Fairel")
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Println(encodeString)
}
