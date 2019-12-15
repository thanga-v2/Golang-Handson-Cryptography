package main

import (
	"encoding/base64"
	"fmt"
)


// To encode string to base64 string, we can use EncodeToString().
// Otherwise, we can decode it using DecodeString() function.

// string message to base64
// decoding base64 to original string message



func main(){

	message := "hello,thanga %$%54** (*w3hu%#"
	demoBase64(message)

}


func demoBase64(message string) {

	fmt.Println("--------Demo encoding base64--------")

	fmt.Println("plaintext:")
	fmt.Println(message)

	// stdEncoding to encode


	encoding := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println("base64 msg:")
	fmt.Println(encoding)

	decoding, _ := base64.StdEncoding.DecodeString(encoding)
	fmt.Println("decoding base64 msg:")
	fmt.Println(string(decoding))

}

