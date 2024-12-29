package main

import (
	"bytes"
	"fmt"
)

func main() {
	values := []string{
		"ZPitL0NGVQBZbTD6DwXJzD8RiStSazzYXQsdUowLURY=",
		"01234567890123456789012345678901",
		"passphrasewhichneedstobe32bytes!",
		"hjBnqfYIHYRTamzDhwnA6hpuyDNDIcAHqb18OrFWYx8=",
		"QJ1cdxSIDYRH9Sf8mGd7rb-YY5zD0oQ2D-HtARMEfEE=",
	}
	for _, value := range values {
		fmt.Print("value:", value, "\n")
		fmt.Print("\tlength:", len([]byte(value)), "\n")
		fmt.Print("\tuint8:", uint8(len(value)+28), "\n")
		fmt.Println("\tinBytes:", inBytes([]byte(value)))
	}
}

func inBytes(content []byte) []byte {
	var result bytes.Buffer
	for i := 0; i < len(content); i++ {
		result.WriteByte(content[i])
	}
	fmt.Println("\tresult:", len(result.Bytes()))
	return result.Bytes()
}
