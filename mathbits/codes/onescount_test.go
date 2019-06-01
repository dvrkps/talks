package main

import (
	"fmt"
	"math/bits"
	"testing"
)

// START OMIT
func ExampleOnesCount() {
	fmt.Printf("%b\n", 42)
	got := bits.OnesCount(42) // HL
	fmt.Println(got)
	// Output:
	// 101010
	// 3
}

// END OMIT

func NaiveOnesCount64(x uint64) int {
	var c, i uint64
	for i = 0; i < 64; i++ {
		c += (x >> i) & 1
	}
	return int(c)
}

var ocNums = []uint64{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	100, 200, 300, 400, 500, 600, 700, 800, 900,
	11111, 22222, 33333, 44444, 55555, 66666, 77777, 88888, 99999,
	190239012390, 123901312, 6549654, 45904059,
}

// ocOutput is a global sync variable so that test outputs
// do not get elided by the compiler.
var ocOutput int

func TestNaiveOnesCount64(t *testing.T) {
	m := len(ocNums)
	for n := 0; n < m; n++ {
		x := ocNums[n%m]
		got := NaiveOnesCount64(x)
		want := bits.OnesCount64(x)
		if got != want {
			t.Errorf("NaiveOnesCount64(%d) = %d, want %d", x, got, want)
		}
	}
}

func BenchmarkNaiveOnesCount64(b *testing.B) {
	// run NaiveOnesCount b.N times
	var s int
	for n := 0; n < b.N; n++ {
		// use &31 to avoid modulo operation
		s += NaiveOnesCount64(ocNums[n&31])
	}
	ocOutput = s
}

func BenchmarkBitsOnesCount64(b *testing.B) {
	// run NaiveOnesCount b.N times
	var s int
	for n := 0; n < b.N; n++ {
		// use &31 to avoid modulo operation
		s += bits.OnesCount64(ocNums[n&31])
	}
	ocOutput = s
}
