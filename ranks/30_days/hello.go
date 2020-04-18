package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

const (
	ReaderBufferSize = 1024 * 1024
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	// fmt.Println("Welcome to 30 Days of Code!")
	reader := bufio.NewReaderSize(os.Stdin, ReaderBufferSize)
	input := readLine(reader)

	println("Hello, World.")
	println(input)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

// func checkError(err error) {
//     if err != nil {
//         panic(err)
//     }
// }
