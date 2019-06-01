package main

import (
	"fmt"
	"math/bits"
)

// START OMIT
func ExampleReverseBytes() {
	fmt.Printf("%016b\n", 810)
	gotBytes := bits.ReverseBytes16(810) // HL
	fmt.Printf("%016b\n", gotBytes)
	gotReverse := bits.Reverse16(810)
	fmt.Printf("%016b\n", gotReverse)
	// Output:
	// 0000001100101010
	// 0010101000000011
	// 0101010011000000
}

// END OMIT
