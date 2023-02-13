package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

func main() {
	var input string
	var choice int
	fmt.Println("Enter the encoded string:")
	fmt.Scanln(&input)
	fmt.Println("Choose the type of encoding:")
	fmt.Println("1. Base64")
	fmt.Println("2. URL")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		data, err := base64.StdEncoding.DecodeString(input)
		if err != nil {
			fmt.Println("Error decoding base64:", err)
			return
		}
		fmt.Println("Base64 Decoded:", string(data))
	case 2:
		data, err := url.QueryUnescape(input)
		if err != nil {
			fmt.Println("Error decoding URL:", err)
			return
		}
		fmt.Println("URL Decoded:", data)
	default:
		fmt.Println("Invalid choice.")
	}
}
