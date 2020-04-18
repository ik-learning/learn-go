package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func birthdayCakeCandles(ar []int32) int32 {
	counter := make(map[int32]int32)
	var tallest int32 = 0

	for i := 0; i < len(ar); i++ {
		if tallest < ar[i] {
			tallest = ar[i]
		}
		if val, ok := counter[ar[i]]; ok {
			counter[ar[i]] += val
		} else {
			counter[ar[i]] = 1
		}
	}
	return counter[tallest]
}

func TestCakeHandles(t *testing.T) {
	var tests = []struct {
		input    []int32
		expected int32
	}{
    {[]int32{1, 2, 3, 4, 4}, 2},
    {[]int32{3, 2, 1, 3}, 2},
	}

	for _, test := range tests {
		val := birthdayCakeCandles(test.input)
		assert.Equal(t, int32(test.expected), val, "they should be equal")
	}
}

// func thhisISHHow() {
//     var n int
//     fmt.Scan(&n)

//     a := make([]int, 10000000)

//     scanner := bufio.NewScanner(os.Stdin)
//     scanner.Split(bufio.ScanWords)

//     for i:=0; i<n; i++ {
//         var h int
//         scanner.Scan()
//         fmt.Sscanf(scanner.Text(), "%d", &h)

//         a[h]++
//     }

//     result := 0
//     for j:=0; j<10000000; j++ {
//         if a[j]>result {
//             result = a[j]
//         }
//     }

//     fmt.Println(result)
// }