package main

import (
	"fmt"
	"math/bits"
	"testing"
)

// START OMIT
func ExampleTrailingZeros() {
	fmt.Printf("%08b\n", 4)
	got := bits.TrailingZeros8(4) // HL
	fmt.Println(got)
	// Output:
	// 00000100
	// 2
}

// END OMIT

// NaiveTrailingZeros64 returns the number of trailing zero bits in x;
// the result is 64 for x == 0.
func NaiveTrailingZeros64(x uint64) int {
	var c, i uint64
	for i = 0; i < 64; i++ {
		if (x>>i)&1 == 1 {
			break
		}
		c++
	}
	return int(c)
}

var tzNums = []uint64{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	100, 200, 300, 400, 500, 600, 700, 800, 900,
	11111, 22222, 33333, 44444, 55555, 66666, 77777, 88888, 99999,
	190239012390, 123901312, 6549654, 45904059,
}

// tzOutput is a global sync variable so that test outputs
// do not get elided by the compiler.
var tzOutput int

func TestNaiveTrailingZeros64(t *testing.T) {
	m := len(tzNums)
	for n := 0; n < m; n++ {
		x := tzNums[n%m]
		got := NaiveTrailingZeros64(x)
		want := bits.TrailingZeros64(x)
		if got != want {
			t.Errorf("NaiveTrailingZeros64(%d) = %d, want %d", x, got, want)
		}
	}
}

func BenchmarkNaiveTrailingZeros64(b *testing.B) {
	s := 0
	for n := 0; n < b.N; n++ {
		s += NaiveTrailingZeros64(tzNums[n&31])
	}
	tzOutput = s
}

func BenchmarkBitsTrailingZeros64(b *testing.B) {
	s := 0
	for n := 0; n < b.N; n++ {
		s += bits.TrailingZeros64(tzNums[n&31])
	}
	tzOutput = s
}
