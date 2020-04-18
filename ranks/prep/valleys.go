package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var m = map[string]int {
	"U": 1,
	"D": -1,
}

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	chars := strings.Split(s, "")
	sum := 0
	count := 0
	var prev string
	for i := 0; i < len(chars); i++ {
		val := m[chars[i]]
		sum += val
		if sum < 0 {
			prev = "Valley"
		} else if sum == 0 {
			if prev == "Valley" {
				count += 1
			}
			prev = "Sea Level"
		} else {
			prev = "Mountain"
		}
	}
	return int32(count)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	result := countingValleys(n, s)

	fmt.Println(result)

	// fmt.Fprintf(writer, "%d\n", result)

	// writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}


// func main() {
//     var n int
//     fmt.Scanln(&n)

//     var s string
//     fmt.Scanln(&s)

//     l, v := 0, 0
//     for i, _ := range s {
//         if s[i] == 'U' {
//             if l < 0 && l+1 == 0 {
//                 v++
//             }

//             l++
//         } else if s[i] == 'D' {
//             l--
//         }
//     }

//     fmt.Println(v)
// }