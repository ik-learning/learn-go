package main

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
)

var values = []map[string]string{
	{"QJ1cdxSIDYRH9Sf8mGd7rb-YY5zD0oQ2D-HtARMEfEE=": "64 157 92 119 20 136 13 132 71 245 39 252 152 103 123 173 191 152 99 156 195 210 132 54 15 225 237 1 19 4 124 65 "},
}

func main() {
	str := "QJ1cdxSIDYRH9Sf8mGd7rb-YY5zD0oQ2D-HtARMEfEE="

	fmt.Println("source B64:", str)
	result, err := b64.StdEncoding.DecodeString(str)
	b64.StdEncoding.DecodeString()

	if err != nil {
		fmt.Println("error b64:", err)
	}

	fmt.Println(len(result))

	b := inBytes1(result)
	fmt.Println("inBytes:", len(b), b)
}

func inBytes1(content []byte) []byte {
	var result bytes.Buffer
	for i := 0; i < len(content); i++ {
		result.WriteByte(content[i])
	}
	return result.Bytes()
}
