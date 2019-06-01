package main

import (
	"fmt"
	"math/bits"
)

// START OMIT
func ExampleReverse() {
	fmt.Printf("%08b\n", 12)
	got := bits.Reverse8(12) // HL
	fmt.Printf("%08b\n", got)
	// Output:
	// 00001100
	// 00110000
}

// END OMIT
