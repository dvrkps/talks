package main

import (
	"fmt"
	"math/bits"
)

// START OMIT
func ExampleRotateLeft() {
	fmt.Printf("%08b\n", 3)
	gotLeft := bits.RotateLeft8(3, 2) // HL
	fmt.Printf("%08b\n", gotLeft)
	gotRight := bits.RotateLeft8(3, -1) // HL
	fmt.Printf("%08b\n", gotRight)
	// Output:
	// 00000011
	// 00001100
	// 10000001
}

// END OMIT
