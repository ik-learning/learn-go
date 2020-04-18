package main

import (
	"sync/atomic"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	jobs := make(chan int32)
	done := make(chan bool)
	var ops int32

	go func() {
		for {
			j, ok := <-jobs
			if ok {
				val := ar[j]
				result := int32(0)
				array := ar[j+1:]
				for i := 0; i < len(array); i++ {
					if (array[i]+val)%k == 0 {
						result++
					}
				}
				atomic.AddInt32(&ops, result)
			} else {
				done <- true
				return
			}
		}
	}()

	for j := 0; j < len(ar); j++ {
		if len(ar[j+1:]) > 0 {
			jobs <- int32(j)
		}
	}
	close(jobs)
	<-done

	return atomic.LoadInt32(&ops)
}

func TestDivisibleSumPair(t *testing.T) {
	for _, test := range testsDivisibleSumPairs {
		result := divisibleSumPairs(0, test.k, test.arr)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}

func BenchmarkTestDivisibleSumPair(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testsDivisibleSumPairs {
			_ = divisibleSumPairs(0, test.k, test.arr)
		}
	}
}

var testsDivisibleSumPairs = []struct {
	k        int32
	arr      []int32
	expected int32
}{
	{3, []int32{1, 3, 2, 6, 1, 2}, 5},
	{95, []int32{27, 89, 4, 69, 18, 56, 93, 41, 51, 11, 39, 48, 99, 57, 67, 32, 23, 23, 39, 70, 26, 79, 93, 75, 76, 72, 36, 88, 60, 67, 95, 58, 29, 7, 70, 60, 6, 72, 24, 97, 19, 98, 64, 38, 14, 64, 88, 34, 5, 98, 8, 79, 57, 5, 43, 27, 57, 77, 89, 8, 45, 66, 60, 98, 20, 79, 99, 98, 6, 48, 42, 77, 43, 83, 48, 77, 83, 49, 40, 32, 13, 99, 23, 55, 2, 94, 80, 62, 20, 60, 97, 80, 9, 54, 67, 84, 60, 62, 97, 64}, 36},
}

// 28213 ns/op
// 16083 ns/op
