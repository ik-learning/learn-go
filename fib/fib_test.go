package fib

import "testing"

// use values that you know are right
var tests = []uint64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

func TestFibFunc(t *testing.T) {
	fn := FibFunc()
	for i, v := range tests {
		if val := fn(); val != v {
			t.Fatalf("at index id %d, expected %d, god %d", i, v, val)
		}
	}
}

func BenchmarkFibFunc(b *testing.B) {
    fn := FibFunc()
    for i := 0; i < b.N; i++ {
        _ = fn()
    }
}

// var mul_tests = []struct {
//     a, b     int
//     expected int
// }{
//     {1, 1, 1},
//     {2, 2, 4},
//     {3, 3, 9},
//     {4, 4, 16},
//     {5, 5, 25},
// }

// func TestMul(t *testing.T) {
//     for _, mt := range mul_tests {
//         if v := Mul(mt.a, mt.b); v != mt.expected {
//             t.Errorf("Mul(%d, %d) returned %d, expected %d", mt.a, mt.b, v, mt.expected)
//         }
//     }
// }